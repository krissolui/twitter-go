package main

import "net/http"

func (app *Config) verifyToken(w http.ResponseWriter, req *http.Request) {
	app.writeResponse(w, "verify token")
}

func (app *Config) signup(w http.ResponseWriter, req *http.Request) {
	app.writeResponse(w, "sign up")
}

func (app *Config) login(w http.ResponseWriter, req *http.Request) {
	app.writeResponse(w, "log in")
}

func (app *Config) logout(w http.ResponseWriter, req *http.Request) {
	app.writeResponse(w, "log out")
}
