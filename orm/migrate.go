package orm

import (
	"log"

	"github.com/codebdy/entify-core/db"
	"github.com/codebdy/entify-core/db/dialect"
	"github.com/codebdy/entify-core/model"
	"github.com/codebdy/entify-core/model/table"
)

func Migrage(d *model.Diff, dbConfig db.DbConfig) {
	var undoList []string
	session, err := OpenWithoutRepository(dbConfig)
	dbx := session.Dbx
	if err != nil {
		panic("Open db error:" + err.Error())
	}

	for _, table := range d.DeletedTables {
		err = DeleteTable(table, &undoList, dbx)
		if err != nil {
			rollback(undoList, dbx)
			panic("Delete table error:" + err.Error())
		}
	}

	for _, table := range d.AddedTables {
		err = CreateTable(table, &undoList, dbx)
		if err != nil {
			rollback(undoList, dbx)
			panic("Create table error:" + err.Error())
		}
	}

	for _, tableDiff := range d.ModifiedTables {
		err = ModifyTable(tableDiff, &undoList, dbx)
		if err != nil {
			rollback(undoList, dbx)
			panic("Modify table error:" + err.Error())
		}
	}
}

func DeleteTable(table *table.Table, undoList *[]string, dbx *db.Dbx) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildDeleteTableSQL(table)
	undoSQL := sqlBuilder.BuildCreateTableSQL(table)
	_, err := dbx.Exec(excuteSQL)
	if err != nil {
		return err
	}
	*undoList = append(*undoList, undoSQL)
	log.Println("Delete Table SQL:", excuteSQL)
	return nil
}

func CreateTable(table *table.Table, undoList *[]string, dbx *db.Dbx) error {
	sqlBuilder := dialect.GetSQLBuilder()
	excuteSQL := sqlBuilder.BuildCreateTableSQL(table)
	undoSQL := sqlBuilder.BuildDeleteTableSQL(table)
	_, err := dbx.Exec(excuteSQL)
	if err != nil {
		return err
	}
	*undoList = append(*undoList, undoSQL)
	log.Println("Add Table SQL:", excuteSQL)

	return nil
}

func ModifyTable(tableDiff *model.TableDiff, undoList *[]string, dbx *db.Dbx) error {
	sqlBuilder := dialect.GetSQLBuilder()
	atoms := sqlBuilder.BuildModifyTableAtoms(tableDiff)
	for _, atom := range atoms {
		_, err := dbx.Exec(atom.ExcuteSQL)
		if err != nil {
			log.Println("Error atom", atom.ExcuteSQL, err.Error())
			return err
		}
		*undoList = append(*undoList, atom.UndoSQL)
	}
	return nil
}

func rollback(undoList []string, con *db.Dbx) {
	for _, sql := range undoList {
		_, err := con.Exec(sql)
		if err != nil {
			log.Println("Rollback failed:", sql)
		}
	}
}
