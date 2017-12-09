package mkdb

import (
	"mkgo/mkconfig"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"mkgo/mklog"
)

type DBService struct {
	WriteDB *sqlx.DB
	ReadDBs []*sqlx.DB
}

var dbService = func() (serv *DBService) {
	var readConfigArr []*mkconfig.DBConfig
	for i := 0; i < mkconfig.Config.MLGO.Datasource.ReadSize; i++ {
		switch i {
		case 1:
			readConfigArr = append(readConfigArr, mkconfig.Config.MLGO.Datasource.Read1)
		case 2:
			readConfigArr = append(readConfigArr, mkconfig.Config.MLGO.Datasource.Read2)
		case 3:
			readConfigArr = append(readConfigArr, mkconfig.Config.MLGO.Datasource.Read3)
		case 4:
			readConfigArr = append(readConfigArr, mkconfig.Config.MLGO.Datasource.Read4)
		case 5:
			readConfigArr = append(readConfigArr, mkconfig.Config.MLGO.Datasource.Read5)
		}
	}

	serv = &DBService{}

	var db *sqlx.DB
	var err error
	db, err = sqlx.Connect(mkconfig.Config.MLGO.Datasource.Write.Driver, mkconfig.Config.MLGO.Datasource.Write.Host)
	if err != nil {
		mklog.Logger.Error("[database]", zap.Error(err))
	} else {
		setupDB(mkconfig.Config.MLGO.Datasource.Write, db)
		serv.WriteDB = db
	}

	var conf *mkconfig.DBConfig
	for _, conf = range readConfigArr {
		db, err = sqlx.Connect(conf.Driver, conf.Host)
		if err != nil {
			mklog.Logger.Error("[database]", zap.Error(err))
			continue
		}
		setupDB(conf, db)
		serv.ReadDBs = append(serv.ReadDBs, db)
	}
	return
}

func setupDB(conf *mkconfig.DBConfig, db *sqlx.DB) {
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxOpenConns)
}

