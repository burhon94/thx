package main

import (
	"fmt"
	"github.com/burhon94/thx/pkg/configs"
	"log"
	"github.com/burhon94/thx/cmd/app"
)

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
