package app

import (
	"context"
	"log"
	"time"

	"github.com/burhon94/thx/pkg/configs"

	"github.com/jackc/pgx/v4/pgxpool"
)


func Run(configuration configs.Configuration) {
	log.Printf("configs: %v", configuration)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute*5)
	pool, err := pgxpool.Connect(ctx, configuration.DSN)
	if err != nil {
		log.Fatalf("can't connect to DB: %v", err)
	}
	log.Printf("pgxPool: %v", pool)

	err = InitServer(configuration)
	if err != nil {
		log.Panicln("server: " + err.Error())
	}
}
