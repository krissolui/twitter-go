package main

import (
	"auth-service/storage"
	"fmt"
	"log"
	"net/http"

	osutils "github.com/krissolui/go-utils/os-utils"
)

const (
	defaultPort = "80"
)

func main() {
	port := osutils.GetEnv("PORT", defaultPort)
	dsn := osutils.GetEnv("DSN")
	if dsn == "" {
		log.Fatal("DSN is required but not found!")
	}

	log.Printf("Connecting to postgres...\n")
	storage, err := storage.NewStorage(dsn, &storage.ConnOptions{
		Attempts:      10,
		DelayInSecond: 2,
	})
	if err != nil {
		log.Fatalf("Failed to connect postgres! %v", err)
	}
	log.Printf("Connected to postgres.\n")

	app := Config{
		Storage: storage,
	}
	defer app.Storage.DB.Close()

	log.Printf("Starting auth service on port %s...", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
