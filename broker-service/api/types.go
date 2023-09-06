package main

type ErrorCode string

type Config struct{}

type RequestPayload struct {
	Action string `json:"action"`
}

type JsonResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Error   ErrorCode `json:"error,omitempty"`
}
