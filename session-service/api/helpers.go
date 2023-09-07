package main

import (
	"encoding/json"
	"net/http"
	"session-service/storage"

	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) writeResponse(w http.ResponseWriter, message string) {
	response := JsonResponse{
		Success: true,
		Message: message,
	}

	httputils.ResponseJSON(w, response, http.StatusAccepted)
}

func (app *Config) errorResponse(
	w http.ResponseWriter,
	err error,
	errCode int,
	status ...int,
) {
	response := JsonResponse{
		Success: false,
		Message: err.Error(),
		Error:   errorCode(errCode),
	}
	httputils.ResponseError(w, response, status...)
}

func (app *Config) requestPayloadToSession(payload RequestPayload) storage.Session {
	return storage.Session{
		UserID:    payload.UserID,
		Token:     payload.Token,
		TTL:       payload.TTL,
		CreatedAt: payload.CreatedAt,
		ExpireAt:  payload.ExpireAt,
	}
}

func (app *Config) sessionToString(session storage.Session) (string, error) {
	jsonData, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
