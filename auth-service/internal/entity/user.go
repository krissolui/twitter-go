package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type User struct {
	UserID    string    `json:"userID" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Icon      string    `json:"icon,omitempty" validate:"url"`
	CreatedAt time.Time `json:"createdAt,omitempty" validate:""`
	UpdatedAt time.Time `json:"updatedAt,omitempty" validate:""`
}

type UserPassword struct {
	UserID    string    `json:"userID" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Algorithm string    `json:"algorithm,omitempty" validate:""`
	Enabled   bool      `json:"enabled,omitempty" validate:""`
	CreatedAt time.Time `json:"createdAt,omitempty" validate:""`
	UpdatedAt time.Time `json:"updatedAt,omitempty" validate:""`
}

type CreateUserEntry struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Icon     string `json:"icon,omitempty" validate:"url"`
}

type UpdatePasswordEntry struct {
	UserID      string `json:"userID" validate:"required"`
	Password    string `json:"password" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}

type LoginEntry struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Session struct {
	UserID    string    `json:"user_id" bson:"user_id" validate:"required"`
	Token     string    `json:"token" bson:"token" validate:"required"`
	TTL       string    `json:"ttl" bson:"ttl" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" validate:"required"`
	ExpireAt  time.Time `json:"expire_at" bson:"expire_at" validate:"required"`
}

func (u *User) Validate() bool {
	return validate.Struct(*u) == nil
}

func (u *UserPassword) Validate() bool {
	return validate.Struct(*u) == nil
}

func (u *CreateUserEntry) Validate() bool {
	return validate.Struct(*u) == nil
}

func (u *UpdatePasswordEntry) Validate() bool {
	return validate.Struct(*u) == nil
}

func (l *LoginEntry) Validate() bool {
	return validate.Struct(*l) == nil
}
