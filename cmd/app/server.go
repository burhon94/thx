package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	db     *serviceDB
	url    string
}

func NewServer(url string, router *mux.Router, db *serviceDB) *server {
	return &server{
		router: router,
		db:     db,
		url:    url,
	}
}

func InitServer(server *server) error {
	newServer := http.Server{
		Addr:         fmt.Sprintf(":%s", server.url),
		Handler:      server.router,
		ReadTimeout:  time.Second * 120,
		WriteTimeout: time.Second * 120,
	}

	log.Printf("server starting on: %s", server.url)
	return newServer.ListenAndServe()
}
