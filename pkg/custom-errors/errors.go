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
	// ErrDatabaseNameRequired is returned when a database name is not provided
	ErrDatabaseNameRequired = errors.New("database name is required")
	// ErrDatabaseURLRequired is returned when a database url is not provided
	ErrDatabaseURLRequired = errors.New("database url is required")
	// ErrServerPortRequired is returned when a port is not provided
	ErrServerPortRequired = errors.New("port is required")
	// ErrAppNameRequired is returned when a app name is not provided
	ErrAppNameRequired = errors.New("app name is required")
	// ErrAppDescriptionRequired is returned when an app description is not provided
	ErrAppDescriptionRequired = errors.New("app description is required")
	// ErrAppBaseURLRequired is returned when an app base url is not provided
	ErrAppBaseURLRequired = errors.New("app base url is required")
	// ErrAppVersionRequired is returned when an app version is not provided
	ErrAppVersionRequired = errors.New("app version is required")
	// ErrLogLevelRequired is returned when a log level is not provided
	ErrLogLevelRequired = errors.New("log level is required")
	// ErrLogMessageRequired is returned when a log message is not provided
	ErrLogMessageRequired = errors.New("log message is required")
	// ErrAppHostRequired is returned when an app host is not provided
	ErrAppHostRequired = errors.New("app host is required")
	// ErrSecretRequired is returned when a secret is not provided
	ErrSecretRequired = errors.New("secret is required")
	// ErrExpiresInRequired is returned when an expires_in is not provided
	ErrExpiresInRequired = errors.New("days of expiration is required")
	// ErrPasswordConfirmation is returned when a password confirmation is not same as password
	ErrPasswordConfirmation = errors.New("password confirmation is not same as password")
	// ErrInvalidEmail is returned when an email is not valid
	ErrInvalidEmail = errors.New("email is not valid")
	// ErrPasswordLength is returned when a password is not valid
	ErrPasswordLength = errors.New("password must be at least 8 characters")
	// ErrPasswordFormat is returned when a password is not valid
	ErrPasswordFormat = errors.New("password must contain at least one uppercase letter, one lowercase letter, one number and one special character")
)
