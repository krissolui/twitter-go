package main

import (
	"net/http"

	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) writeResponse(w http.ResponseWriter, message string, data ...any) {
	response := JsonResponse{
		Success: true,
		Message: message,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	httputils.ResponseJSON(w, response, http.StatusAccepted)
}

func (app *Config) errorResponse(w http.ResponseWriter, err error, errCode int, status ...int) {
	response := JsonResponse{
		Success: false,
		Message: err.Error(),
		Error:   errorCode(errCode),
	}

	httputils.ResponseError(w, response, status...)
}
