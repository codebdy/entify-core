package orm

import (
	"fmt"
	"log"

	"github.com/codebdy/entify/db"
	"github.com/codebdy/entify/model"
)

func DbString(cfg db.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func Open(dbConfig db.DbConfig, m *model.Model) (*Session, error) {
	dbx, err := db.Open(dbConfig.Driver, DbString(dbConfig))
	if err != nil {
		return nil, err
	}
	session := Session{
		idSeed: 1,
		Dbx:    dbx,
		model:  m,
	}
	return &session, nil
}

func OpenWithoutRepository(dbConfig db.DbConfig) (*Session, error) {
	return Open(dbConfig, nil)
}

func IsEntityExists(name string, dbConfig db.DbConfig) bool {
	session, err := Open(dbConfig, nil)
	if err != nil {
		log.Panic(err.Error())
	}
	return session.doCheckEntity(name, dbConfig.Database)
}
