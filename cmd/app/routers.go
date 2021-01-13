package app

import (
	"github.com/gorilla/mux"
)

var router mux.Router

type routes struct {
	DBService *serviceDB
}

func NewRoutes(dbService *serviceDB) *routes {
	return &routes{DBService: dbService}
}

func (routers *routes) InitRoutes() mux.Router {
	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/test", routers.handleTest).Methods("POST")

	router.HandleFunc("/sayThx", routers.SayThx).Methods("POST")

	return router
}
