package orm

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codebdy/entify/db"
	"github.com/codebdy/entify/db/dialect"
	"github.com/codebdy/entify/model/data"
	"github.com/codebdy/entify/model/graph"
	"github.com/codebdy/entify/shared"
)

type QueryResponse struct {
	Nodes []map[string]interface{} `json:"nodes"`
	Total int                      `json:"total"`
}

func (s *Session) buildQueryEntitySQL(
	entity *graph.Entity,
	args map[string]interface{},
	whereArgs interface{},
	argEntity *graph.ArgEntity,
	queryBody string,
) (string, []interface{}) {
	var paramsList []interface{}
	//whereArgs := con.v.WeaveAuthInArgs(entity.Uuid(), args[shared.ARG_WHERE])
	// argEntity := graph.BuildArgEntity(
	// 	entity,
	// 	whereArgs,
	// 	con,
	// )
	builder := dialect.GetSQLBuilder()
	queryStr := queryBody
	if where, ok := whereArgs.(graph.QueryArg); ok {
		whereSQL, params := builder.BuildWhereSQL(argEntity, entity.AllAttributes(), where)
		if whereSQL != "" {
			queryStr = queryStr + " WHERE " + whereSQL
		}
		paramsList = append(paramsList, params...)
	}

	queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[shared.ARG_ORDERBY])
	return queryStr, paramsList
}

func (s *Session) buildQueryEntityRecordsSQL(entity *graph.Entity, args map[string]interface{}, attributes []*graph.Attribute) (string, []interface{}) {
	whereArgs := args[shared.ARG_WHERE]
	argEntity := graph.BuildArgEntity(
		entity,
		whereArgs,
		s,
	)
	builder := dialect.GetSQLBuilder()
	queryStr := builder.BuildQuerySQLBody(argEntity, attributes)
	sqlStr, params := s.buildQueryEntitySQL(entity, args, whereArgs, argEntity, queryStr)

	if args[shared.ARG_LIMIT] != nil {
		sqlStr = sqlStr + fmt.Sprintf(" LIMIT %d ", args[shared.ARG_LIMIT])
	}
	if args[shared.ARG_OFFSET] != nil {
		sqlStr = sqlStr + fmt.Sprintf(" OFFSET %d ", args[shared.ARG_OFFSET])
	}

	return sqlStr, params
}

func (s *Session) buildQueryEntityCountSQL(entity *graph.Entity, args map[string]interface{}) (string, []interface{}) {
	whereArgs := args[shared.ARG_WHERE]
	argEntity := graph.BuildArgEntity(
		entity,
		whereArgs,
		s,
	)
	builder := dialect.GetSQLBuilder()
	queryStr := builder.BuildQueryCountSQLBody(argEntity)
	return s.buildQueryEntitySQL(
		entity,
		args,
		whereArgs,
		argEntity,
		queryStr,
	)
}

func (s *Session) Query(entityName string, args map[string]interface{}, fieldNames []string) QueryResponse {
	var instances []InsanceData
	entity := s.model.Graph.GetEntityByName(entityName)
	fields := []*graph.Attribute{}
	allAttributes := entity.AllAttributes()

	for i := range allAttributes {
		for _, name := range fieldNames {
			if allAttributes[i].Name == name {
				fields = append(fields, allAttributes[i])
			}
		}
	}

	if len(fields) > 0 {
		sqlStr, params := s.buildQueryEntityRecordsSQL(entity, args, fields)
		log.Println("doQueryEntity SQL:", sqlStr, params)
		rows, err := s.Dbx.Query(sqlStr, params...)
		defer rows.Close()
		if err != nil {
			log.Panic(err.Error(), sqlStr)
		}

		for rows.Next() {
			values := makeEntityQueryValues(fields)
			err = rows.Scan(values...)
			if err != nil {
				panic(err.Error())
			}
			instances = append(instances, convertValuesToEntity(values, fields))
		}
	}

	sqlStr, params := s.buildQueryEntityCountSQL(entity, args)
	log.Println("doQueryEntity count SQL:", sqlStr, params)
	count := 0
	err := s.Dbx.QueryRow(sqlStr, params...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		count = 0
	case err != nil:
		log.Panic(err.Error())
	}

	return QueryResponse{
		Nodes: instances,
		Total: count,
	}
}

func (s *Session) QueryOneById(entityName string, id interface{}) interface{} {
	return s.QueryOne(entityName, graph.QueryArg{
		shared.ARG_WHERE: graph.QueryArg{
			shared.ID_NAME: graph.QueryArg{
				shared.ARG_EQ: id,
			},
		},
	})
}

func (s *Session) QueryOne(entityName string, args map[string]interface{}) interface{} {
	entity := s.model.Graph.GetEntityByName(entityName)
	queryStr, params := s.buildQueryEntityRecordsSQL(entity, args, entity.AllAttributes())

	values := makeEntityQueryValues(entity.AllAttributes())
	//log.Println("doQueryOneEntity SQL:", queryStr, params)
	err := s.Dbx.QueryRow(queryStr, params...).Scan(values...)
	switch {
	case err == sql.ErrNoRows:
		log.Println(fmt.Sprintf("Can not find instance %s, %s", entity.Name(), args))
		return nil
	case err != nil:
		log.Panic(err.Error())
	}

	instance := convertValuesToEntity(values, entity.AllAttributes())
	return instance
}

func (s *Session) QueryAssociatedInstances(r *data.AssociationRef, ownerId uint64) []InsanceData {
	var instances []InsanceData
	builder := dialect.GetSQLBuilder()
	entity := r.TypeEntity()
	queryStr := builder.BuildQueryAssociatedInstancesSQL(entity, ownerId, r.Table().Name, r.OwnerColumn().Name, r.TypeColumn().Name)
	rows, err := s.Dbx.Query(queryStr)
	defer rows.Close()
	if err != nil {
		log.Panic(err.Error())
	}

	for rows.Next() {
		values := makeEntityQueryValues(entity.AllAttributes())
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToEntity(values, entity.AllAttributes()))
	}

	return instances
}

func (s *Session) QueryByIds(entityName string, ids []interface{}) []InsanceData {
	entity := s.model.Graph.GetEntityByName(entityName)
	var instances []map[string]interface{}
	builder := dialect.GetSQLBuilder()
	sql := builder.BuildQueryByIdsSQL(entity, len(ids))
	rows, err := s.Dbx.Query(sql, ids...)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		values := makeEntityQueryValues(entity.AllAttributes())
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instances = append(instances, convertValuesToEntity(values, entity.AllAttributes()))
	}

	return instances
}

func (s *Session) BatchRealAssociations(
	association *graph.Association,
	ids []uint64,
	args graph.QueryArg,
) []InsanceData {
	var instances []map[string]interface{}
	var paramsList []interface{}

	builder := dialect.GetSQLBuilder()
	typeEntity := association.TypeEntity()
	whereArgs := args[shared.ARG_WHERE]
	argEntity := graph.BuildArgEntity(
		typeEntity,
		whereArgs,
		s,
	)

	queryStr := builder.BuildBatchAssociationBodySQL(argEntity,
		typeEntity.AllAttributes(),
		association.Relation.Table.Name,
		association.Owner().TableName(),
		association.TypeEntity().TableName(),
		ids,
	)

	if where, ok := whereArgs.(graph.QueryArg); ok {
		whereSQL, params := builder.BuildWhereSQL(argEntity, typeEntity.AllAttributes(), where)
		if whereSQL != "" {
			queryStr = queryStr + " AND " + whereSQL
		}
		paramsList = append(paramsList, params...)
	}

	queryStr = queryStr + builder.BuildOrderBySQL(argEntity, args[shared.ARG_ORDERBY])
	log.Println("doBatchRealAssociations SQL:	", queryStr)
	rows, err := s.Dbx.Query(queryStr, paramsList...)
	defer rows.Close()
	if err != nil {
		log.Println("出错SQL:", queryStr)
		log.Panic(err.Error())
	}

	for rows.Next() {
		values := makeEntityQueryValues(typeEntity.AllAttributes())
		var idValue db.NullUint64
		values = append(values, &idValue)
		err = rows.Scan(values...)
		if err != nil {
			panic(err.Error())
		}
		instance := convertValuesToEntity(values, typeEntity.AllAttributes())
		instance[shared.ASSOCIATION_OWNER_ID] = values[len(values)-1].(*db.NullUint64).Uint64
		instances = append(instances, instance)
	}

	return instances
}

func merageInstances(source []InsanceData, target []InsanceData) {
	for i := range source {
		souceObj := source[i]
		for j := range target {
			targetObj := target[j]
			if souceObj[shared.ID_NAME] == targetObj[shared.ID_NAME] {
				targetObj[shared.ASSOCIATION_OWNER_ID] = souceObj[shared.ASSOCIATION_OWNER_ID]
				source[i] = targetObj
			}
		}
	}
}
