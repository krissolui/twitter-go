package main

import (
	"fmt"
	"log"
	"net/http"

	osutils "github.com/krissolui/go-utils/os-utils"
)

const (
	defaultWebPort = "80"
)

func main() {
	webPort := osutils.GetEnv("WEB_PORT", defaultWebPort)

	app := Config{}

	log.Printf("Starting auth service on port %s...", webPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
