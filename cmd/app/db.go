package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceDB struct {
	pool *pgxpool.Pool
}

func NewServiceDB(pool *pgxpool.Pool) *serviceDB {
	return &serviceDB{pool: pool}
}

func (s *serviceDB) add() (err error) {
	sql := `insert into users values (2);`
	_, err = s.pool.Exec(context.Background(), sql)
	if err != nil {
		return
	}

	return nil
}

func (s *serviceDB) sayThx() (err error) {
	sql := `update users
set count_star = count_star + 1
where id = 1;`

	_, err = s.pool.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}
