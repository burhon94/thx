package app

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

var router mux.Router

type routes struct {
	pool *pgxpool.Pool
}

func NewRoutes(pool *pgxpool.Pool) *routes {
	return &routes{pool: pool}
}

func (r *routes) InitRoutes() mux.Router {
	router.HandleFunc("/health", handleHealth).Methods("GET")
	router.HandleFunc("/test", r.handleTest).Methods("POST")

	return router
}
