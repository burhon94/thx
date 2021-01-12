package main

import (
	"fmt"
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

	fmt.Printf("print here some text: ")

	var text string
	_, err = fmt.Scanf("%s", &text)
	if err != nil {
		log.Printf("error when scan: %v", err)
		return
	}
	app.Run(text, cfg)
}
