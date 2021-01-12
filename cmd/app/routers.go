package app

import "github.com/gorilla/mux"

var router mux.Router


func InitRoutes() mux.Router {
	router.HandleFunc("/health", handleHealth).Methods("GET")

	return router
}
