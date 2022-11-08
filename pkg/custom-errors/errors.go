package customErr

import "errors"

var (
	// ErrEmailRequired is returned when an email address is not provided
	ErrEmailRequired = errors.New("email address is required")
	// ErrPasswordRequired is returned when a password is not provided
	ErrPasswordRequired = errors.New("password is required")
	// ErrInvalidCredentials is returned when the provided email address/password combination is invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrCreateAtRequired is returned when a create_at is not provided
	ErrCreateAtRequired = errors.New("create_at is required")
	// ErrTokenRequired is returned when a token is not provided
	ErrTokenRequired = errors.New("token is required")
)
