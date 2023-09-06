package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *Config) readJSON(req *http.Request) (RequestPayload, error) {
	var payload RequestPayload

	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return RequestPayload{}, err
	}

	return payload, nil
}

func (app *Config) responseJSON(w http.ResponseWriter, status int, response JsonResponse, headers ...http.Header) {

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(payload)
}

func (app *Config) writeResponse(w http.ResponseWriter, message string) {
	response := JsonResponse{
		Success: true,
		Message: message,
	}

	app.responseJSON(w, http.StatusAccepted, response)
}

func (app *Config) errorResponse(w http.ResponseWriter, err error, errCode int, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	response := JsonResponse{
		Success: false,
		Message: err.Error(),
		Error:   errorCode(errCode),
	}

	app.responseJSON(w, statusCode, response)
}
