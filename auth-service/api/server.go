package main

import (
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

	app := Config{}

	log.Printf("Starting auth service on port %s...", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
