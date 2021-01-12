package app

import "github.com/jackc/pgx/v4/pgxpool"

type serviceDB struct {
	pool *pgxpool.Pool
}

func NewServiceDB(pool *pgxpool.Pool) *serviceDB {
	return &serviceDB{pool: pool}
}