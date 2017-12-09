package mkdb

import (
	"github.com/jmoiron/sqlx"
	"math/rand"
	"mkgo/mklog"
)

func GetWriteDB() (db *sqlx.DB) {
	db = dbService().WriteDB
	if db == nil {
		mklog.Logger.Error("[database] Write database engine is not configured! ")
	}
	return db
}

func GetReadDB() (db *sqlx.DB) {
	engineSize := len(dbService().ReadDBs)
	if engineSize <= 0 {
		return nil
	}
	if engineSize == 1 {
		return dbService().ReadDBs[0]
	}
	return dbService().ReadDBs[rand.Intn(engineSize)]
}

type Table interface {
	TableName() string
}
