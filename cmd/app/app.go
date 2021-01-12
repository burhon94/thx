package app

import (
	"github.com/burhon94/thx/pkg/configs"
	"log"
	"os"
)

func Run(configuration configs.Configuration)  {
	log.Printf("server start on: %v", configuration.Port)
	os.Exit(200)
}
