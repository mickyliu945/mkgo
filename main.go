package main

import (
	"mkgo/mkconfig"
	"mkgo/mklog"
	"mkgo/server"
)

func main() {
	mkconfig.Init()
	mklog.Init()
	server.Init()
}
