package models

import (
	customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"
	"time"
)

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return customErr.ErrEmailRequired
	}
	if u.Password == "" {
		return customErr.ErrPasswordRequired
	}
	if u.CreatedAt.IsZero() {
		return customErr.ErrCreateAtRequired
	}
	return nil
}
