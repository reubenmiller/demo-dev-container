package main

import (
	"log"
	"os"
	"time"
)

func main() {
	name := os.Getenv("NAME")
	if name == "" {
		name = "world"
	}

	for {

		log.Printf("Hello %s\n", name)
		time.Sleep(1 * time.Second)
	}
}
