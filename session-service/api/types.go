package main

import "go.mongodb.org/mongo-driver/mongo"

type ErrorCode string

type Config struct {
	mongo *mongo.Collection
}

type RequestPayload struct {
	Action string `json:"action"`
}

type JsonResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Error   ErrorCode `json:"error,omitempty"`
}
