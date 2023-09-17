package util

import (
	customerror "auth-service/internal/customError"
	"net/http"

	httputils "github.com/krissolui/go-utils/http-utils"
)

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func WriteResponse(w http.ResponseWriter, message string, data ...any) {
	response := JsonResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	httputils.ResponseJSON(w, response)
}

func ErrorResponse(w http.ResponseWriter, err *customerror.CustomError) {
	response := JsonResponse{
		Success: false,
		Message: err.Message,
		Error:   err.Error,
	}

	httputils.ResponseError(w, response, err.Code)
}
