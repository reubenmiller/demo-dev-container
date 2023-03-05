package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Get name to greet with
func getName() string {
	name := os.Getenv("NAME")
	if name == "" {
		name = "world"
	}
	return name
}

func helloHandler(name string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s hello %s\n", time.Now().Format(time.RFC3339), name)
	}
}

func main() {
	name := getName()

	// Run background http server
	http.HandleFunc("/", helloHandler(name))
	http.HandleFunc("/hello", helloHandler(name))
	go http.ListenAndServe(":9000", nil)

	// Just print hello forever on the console
	for {
		log.Printf("Hello %s\n", name)
		time.Sleep(1 * time.Second)
	}
}
