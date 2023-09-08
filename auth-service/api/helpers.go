package main

import (
	"net/http"

	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) writeResponse(w http.ResponseWriter, message string, data ...any) {
	response := JsonResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	httputils.ResponseJSON(w, response, http.StatusAccepted)
}

func (app *Config) errorResponse(w http.ResponseWriter, err error, errorCode ErrorCode, statusCode ...int) {
	response := JsonResponse{
		Success: false,
		Message: err.Error(),
		Error:   errorCode.toString(),
	}

	httputils.ResponseError(w, response, statusCode...)
}
