package main

import (
	"fmt"
	"net/http"
	"strconv"

	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) getActiveSession(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")
	if userID == "" {
		app.errorResponse(w, fmt.Errorf("user_id is required"), EInvalidParams)
		return
	}

	session, err := app.mongo.GetActiveSession(userID)
	if err != nil || session == nil {
		app.writeResponse(w, "")
		return
	}

	response, err := app.sessionToString(*session)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, response)
}

func (app *Config) verifySession(w http.ResponseWriter, req *http.Request) {
	requestPayload, err := httputils.ReadJSON[RequestPayload](req)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
	}

	valid := app.mongo.VerifyToken(requestPayload.UserID, requestPayload.Token)
	app.writeResponse(w, strconv.FormatBool(valid))
}

func (app *Config) createSession(w http.ResponseWriter, req *http.Request) {
	requestPayload, err := httputils.ReadJSON[RequestPayload](req)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
	}

	session := app.requestPayloadToSession(*requestPayload)
	// TODO::validate session (expire_at > created_at)

	err = app.mongo.WriteSession(session)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	response, err := app.sessionToString(session)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, response)
}
