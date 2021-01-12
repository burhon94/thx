package main

import (
	"log"

	"github.com/burhon94/thx/cmd/app"
	"github.com/burhon94/thx/pkg/configs"

	lumberJack "gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	log.SetOutput(&lumberJack.Logger{
		Filename:   "app.log",
		MaxSize:    100, //megabytes
		MaxAge:     28,  //days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	})
}

func main() {

	cfg, err := configs.Init()
	if err != nil {
		log.Printf("error, can't read configs file: %v", err)
		return
	}

	app.Run(cfg)

}
