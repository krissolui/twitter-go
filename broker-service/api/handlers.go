package main

import (
	"net/http"
	"time"
)

func (app *Config) root(w http.ResponseWriter, req *http.Request) {
	// TODO::remove hard coded session
	session := Session{
		UserID:    "123",
		Token:     "haha",
		TTL:       "60",
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(time.Minute),
	}
	err := app.createSession(session)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "created session")
}

func (app *Config) fail(w http.ResponseWriter, req *http.Request) {
	session, err := app.getSession("123")
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "", session)
}
