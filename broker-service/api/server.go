package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defualtWebPort = "80"

func main() {
	// get port from environment or use default port
	var webPort = defualtWebPort
	if p := os.Getenv("WEB_PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			webPort = p
		}
	}

	log.Printf("Starting broker on port %s...\t", webPort)
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.route(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Broker listening on port %s...\t", webPort)
}
