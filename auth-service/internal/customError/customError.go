package customerror

import "net/http"

type CustomError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func InvalidCredentials(message string) *CustomError {
	return &CustomError{
		Error:   "INVALID_CREDENTIALS",
		Message: message,
		Code:    http.StatusNotAcceptable,
	}
}

func PathNotFound(message string) *CustomError {
	return &CustomError{
		Error:   "PATH_NOT_FOUND",
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func InvalidParams(message string) *CustomError {
	return &CustomError{
		Error:   "INVALID_PARAMS",
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func InternalError(message string) *CustomError {
	return &CustomError{
		Error:   "INTERNAL_ERROR",
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
