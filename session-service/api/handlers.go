package main

import (
	"errors"
	"net/http"
)

func (app *Config) root(w http.ResponseWriter, req *http.Request) {
	app.writeResponse(w, "hello!")
}

func (app *Config) fail(w http.ResponseWriter, req *http.Request) {
	app.errorResponse(w, errors.New("should fail"), EInvalidAction)
}
