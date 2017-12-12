package main

import (
	"mkgo/mkconfig"
	"mkgo/mklog"
	"mkgo/server"
	"mkgo/mkdb"
)

func main() {
	mkconfig.Init()
	mklog.Init()
	mkdb.InitDB()
	server.Init()
}
