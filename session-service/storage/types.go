package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type Session struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Token     string    `json:"token" bson:"token"`
	TTL       string    `json:"ttl" bson:"ttl"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ExpireAt  time.Time `json:"expire_at" bson:"expire_at"`
}
