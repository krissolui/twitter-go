package main

import "time"

type ErrorCode string

type Config struct {
	SessionServiceURL string
}

type RequestPayload struct {
	Action string `json:"action"`
}

type JsonResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    any       `json:"data,omitempty"`
	Error   ErrorCode `json:"error,omitempty"`
}

type Session struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Token     string    `json:"token" bson:"token"`
	TTL       string    `json:"ttl" bson:"ttl"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ExpireAt  time.Time `json:"expire_at" bson:"expire_at"`
}
