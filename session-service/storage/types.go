package storage

import "time"

type Session struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	TTL       string    `json:"ttl"`
	Expired   bool      `json:"expired"`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:updated_at`
}
