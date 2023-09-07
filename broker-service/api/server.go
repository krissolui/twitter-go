package main

import (
	"fmt"
	"log"
	"net/http"

	osutils "github.com/krissolui/go-utils/os-utils"
)

const (
	defualtWebPort           = "80"
	defaultSessionServiceURL = "http://session-service"
)

func main() {
	webPort := osutils.GetEnv("WEB_PORT", defualtWebPort)
	sessionServiceURL := osutils.GetEnv("SESSION_SERVICE_URL", defaultSessionServiceURL)

	log.Printf("Starting broker on port %s...\t", webPort)
	app := Config{
		SessionServiceURL: sessionServiceURL,
	}

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
