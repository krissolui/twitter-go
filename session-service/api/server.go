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

const defualtPort = "80"

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
	var port = defualtPort
	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	log.Printf("Starting session service on port %s...\t", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.route(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Broker listening on port %s...\t", port)
}
