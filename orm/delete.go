package orm

import (
	"log"

	"github.com/codebdy/entify/db/dialect"
	"github.com/codebdy/entify/model/data"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/shared"
)

type InsanceData = map[string]interface{}

func (s *Session) clearSyncedAssociation(a *graph.Association, ownerId uint64, synced []*data.Instance) {

	//查出所有关联实例
	associatedInstances := s.QueryAssociatedInstances(a, ownerId)

	for _, associatedIns := range associatedInstances {
		willBeDelete := true

		//找出需要被删除的
		for _, syncedIns := range synced {
			if syncedIns.Id != 0 && syncedIns.Id == associatedIns[shared.ID_NAME] {
				willBeDelete = false
				continue
			}
		}

		//删除需要被删除的
		if willBeDelete {
			//如果是组合，被关联实例
			if a.IsCombination() {
				ins := data.NewInstance(associatedIns, a.TypeEntity())
				s.DeleteInstance(ins.Entity.Name(), ins.Id)
			}
			s.deleteAssociationPovit(a, associatedIns[shared.ID_NAME].(uint64))
		}
	}
}
func (con *Session) clearAssociation(r *graph.Association, ownerId uint64) {
	if r.IsCombination() {
		con.deleteAssociatedInstances(r, ownerId)
	}
	con.deleteAssociationPovit(r, ownerId)
}

func (s *Session) deleteAssociationPovit(a *graph.Association, ownerId uint64) {
	sqlBuilder := dialect.GetSQLBuilder()
	//先检查是否有数据，如果有再删除，避免死锁
	sql := sqlBuilder.BuildCheckAssociationSQL(ownerId, a.Table().Name, a.TypeColumn().Name)
	count := s.queryCount(sql)
	if count > 0 {
		sql = sqlBuilder.BuildClearAssociationSQL(ownerId, a.Table().Name, a.TypeColumn().Name)
		_, err := s.Dbx.Exec(sql)
		log.Println("deleteAssociationPovit SQL:" + sql)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (s *Session) queryCount(countSql string) int64 {
	rows, err := s.Dbx.Query(countSql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return 0
	} else {
		var count int64
		for rows.Next() {
			rows.Scan(&count)
		}
		return count
	}
}

func (s *Session) deleteAssociatedInstances(a *graph.Association, ownerId uint64) {
	typeEntity := a.TypeEntity()
	associatedInstances := s.QueryAssociatedInstances(a, ownerId)
	for i := range associatedInstances {
		s.DeleteInstance(typeEntity.Name(), associatedInstances[i]["id"].(shared.ID))
	}
}

func (s *Session) DeleteAssociationPovit(povit *data.AssociationPovit) {
	sqlBuilder := dialect.GetSQLBuilder()
	sql := sqlBuilder.BuildDeletePovitSQL(povit)
	_, err := s.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func (s *Session) DeleteInstance(entityName string, id shared.ID) {
	var sql string
	entity := s.model.Graph.GetEntityByName(entityName)
	sqlBuilder := dialect.GetSQLBuilder()
	tableName := entity.TableName()
	if entity.IsSoftDelete() {
		sql = sqlBuilder.BuildSoftDeleteSQL(id, tableName)
	} else {
		sql = sqlBuilder.BuildDeleteSQL(id, tableName)
	}

	log.Println("DeleteInstance:", sql)
	_, err := s.Dbx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	associstions := entity.Associations()
	for i := range associstions {
		asso := associstions[i]
		if asso.IsCombination() {
			if !asso.TypeEntity().IsSoftDelete() {
				s.deleteAssociationPovit(asso, id)
			}
			s.deleteAssociatedInstances(asso, id)
		}
	}
}
