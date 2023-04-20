package orm

import (
	"fmt"
	"log"

	"github.com/codebdy/entify/db"
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

func Open(dbConfig db.DbConfig) (*Session, error) {
	dbx, err := db.Open(dbConfig.Driver, DbString(dbConfig))
	if err != nil {
		return nil, err
	}
	session := Session{
		idSeed: 1,
		Dbx:    dbx,
		//model:  model,
	}
	return &session, nil
}

func IsEntityExists(name string, dbConfig db.DbConfig) bool {
	session, err := Open(dbConfig)
	if err != nil {
		log.Panic(err.Error())
	}
	return session.doCheckEntity(name)
}
