package app

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Run(text string)  {
	log.Printf("You printed: %s", text)
	for i := 5; i >= 1; i-- {
		fmt.Printf("Bye: %v\n", i)
		time.Sleep(time.Second)
		continue
	}

	os.Exit(200)
}
