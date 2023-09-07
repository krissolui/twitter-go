package main

import (
	"auth-service/storage"
)

type Config struct {
	Storage storage.Storage
}

type ErrorCode int

type JsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
