package server

import (
	"mkgo/mkconfig"
	"net/http"
	"time"
)

func Init() {
	router := GetRouter()
	port := mkconfig.Config.MKGo.ServerPort
	if port == "" {
		port = "8080"
	}
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    mkconfig.Config.MKGo.ReadTimeout * time.Second,
		WriteTimeout:   mkconfig.Config.MKGo.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
