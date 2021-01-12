package main

import (
	"fmt"
	"log"
	"github.com/burhon94/thx/cmd/app"
)

func main() {
	fmt.Printf("print here some text: ")

	var text string
	_, err := fmt.Scanf("%s", &text)
	if err != nil {
		log.Printf("error when scan: %v", err)
		return
	}
	app.Run(text)
}
