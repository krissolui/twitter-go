package storage

import (
	"database/sql"
	"time"
)

type Storage struct {
	DB *sql.DB
}

type ConnOptions struct {
	Attempts      int
	DelayInSecond int
}

type User struct {
	UserID    string    `json:"userID"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Icon      string    `json:"icon,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type UserPassword struct {
	UserID    string    `json:"userID"`
	Password  string    `json:"password"`
	Algorithm string    `json:"algorithm,omitempty"`
	Enabled   bool      `json:"enabled,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CreateUserEntry struct {
	User
	Password string `json:"password"`
}

type UpdatePasswordEntry struct {
	UserPassword
	OldPassword string `json:"oldPassword"`
}

type LoginEntry struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
