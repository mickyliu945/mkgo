package mkdb

import (
	_ "github.com/go-sql-driver/mysql"
	"mkgo/mkconfig"
	"github.com/linxGnu/mssqlx"
)

var DB *mssqlx.DBs


func InitDB() {
	if len(mkconfig.Config.DataSource.Write) <= 0 {
		panic("DataSource config error!")
		return
	}

	DB, _ = mssqlx.ConnectMasterSlaves("mysql", mkconfig.Config.DataSource.Write, mkconfig.Config.DataSource.Read)
	if DB == nil {
		panic("Database init failed!")
		return
	}
	DB.SetMaxIdleConns(mkconfig.Config.DataSource.MaxIdleConns)
	DB.SetMaxOpenConns(mkconfig.Config.DataSource.MaxOpenConns)
	DB.SetHealthCheckPeriod(1000)
}
