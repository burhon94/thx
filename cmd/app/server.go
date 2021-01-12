package app

import (
	"fmt"
	"github.com/burhon94/thx/pkg/configs"
	"log"
	"net/http"
	"time"
)

func InitServer(config configs.Configuration) (err error) {

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", config.Port),
		Handler:           nil,
		ReadHeaderTimeout: 120 * time.Second,
		WriteTimeout:      120 * time.Second,
	}

	log.Printf("server: starting on 127.0.0.1:%s\n", config.Port)
	return server.ListenAndServe()
}
