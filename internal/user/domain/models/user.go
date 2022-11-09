package models

import (
	"encoding/hex"
	customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/sha3"
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
	if u.CreatedAt.IsZero() {
		return customErr.ErrCreateAtRequired
	}
	return nil
}

func (u *User) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}
