package main

import (
	"session-service/storage"
)

type ErrorCode string

type Config struct {
	mongo storage.Storage
}

type RequestPayload struct {
	storage.Session
}

type JsonResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Error   ErrorCode `json:"error,omitempty"`
}
