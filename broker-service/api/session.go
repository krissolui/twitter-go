package main

import (
	"encoding/json"
	"log"
	"net/http"

	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) createSession(session Session) error {
	var empty map[string]string
	res, err := httputils.SendRequestWithBody(app.SessionServiceURL, "/create", http.MethodPost, empty, session)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Printf("create session response: %v\n", res.Body)

	return nil
}

func (app *Config) getSession(userID string) (*Session, error) {
	var query = map[string]string{}
	query["user_id"] = userID

	res, err := httputils.SendPlainRequest(app.SessionServiceURL, "session", http.MethodGet, query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Printf("get session response: %v\n", res.Body)
	var jsonResponse JsonResponse
	_ = json.NewDecoder(res.Body).Decode(&jsonResponse)

	log.Printf("jsonResponse: %v\n", jsonResponse)

	if jsonResponse.Message == "" {
		return nil, nil
	}

	var session Session
	_ = json.Unmarshal([]byte(jsonResponse.Message), &session)
	log.Printf("session: %v\n", session)

	return &session, nil
}
