package app

import (
	"fmt"
	"github.com/burhon94/thx/pkg/configs"
	"log"
	"os"
	"time"
)

func Run(text string, configuration configs.Configuration)  {
	log.Printf("You printed: %s", text)
	log.Printf("Your configs: %v", configuration)
	for i := 5; i >= 1; i-- {
		fmt.Printf("Bye: %v\n", i)
		time.Sleep(time.Second)
		continue
	}

	os.Exit(200)
}
