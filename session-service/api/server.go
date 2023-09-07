package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"session-service/storage"
	"strconv"
)

const defualtWebPort = "80"

func main() {
	mongo, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Client.Disconnect(context.TODO())

	app := Config{
		mongo: mongo,
	}

	// get port from environment or use default port
	var webPort = defualtWebPort
	if p := os.Getenv("WEB_PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			webPort = p
		}
	}

	log.Printf("Starting session service on port %s...\t", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.route(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Broker listening on port %s...\t", webPort)
}
