package orm

import (
	"database/sql"
	"log"

	"github.com/codebdy/entify-core/db"
	"github.com/codebdy/entify-core/db/dialect"
	"github.com/codebdy/entify-core/model"
)

type Session struct {
	idSeed int //use for sql join table
	model  *model.Model
	Dbx    *db.Dbx
}

func (s *Session) BeginTx() error {
	return s.Dbx.BeginTx()
}

func (s *Session) Commit() error {
	return s.Dbx.Commit()
}

func (s *Session) ClearTx() {
	s.Dbx.ClearTx()
}

// use for sql join table
func (s *Session) CreateId() int {
	s.idSeed++
	return s.idSeed
}

func (s *Session) doCheckEntity(name, database string) bool {
	sqlBuilder := dialect.GetSQLBuilder()
	var count int
	err := s.Dbx.QueryRow(sqlBuilder.BuildTableCheckSQL(name, database)).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Panic(err.Error())
	}
	return count > 0
}
