package models

import (
	"encoding/hex"
	customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"
	"github.com/solrac97gr/go-jwt-auth/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/sha3"
	"net/mail"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return customErr.ErrEmailRequired
	}
	if u.Password == "" {
		return customErr.ErrPasswordRequired
	}
	if !utils.IsValid(u.Password) {
		return customErr.ErrPasswordFormat
	}
	if u.CreatedAt.IsZero() {
		return customErr.ErrCreateAtRequired
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return customErr.ErrInvalidEmail
	}
	return nil
}

func (u *User) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}
