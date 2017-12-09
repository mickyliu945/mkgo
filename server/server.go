package server

import (
	"mkgo/mkconfig"
	"net/http"
	"time"
)

func Init() {
	router := GetRouter()
	port := mkconfig.Config.MLGO.ServerPort
	if port == "" {
		port = "8080"
	}
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
