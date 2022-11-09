package ports

type MiddlewareRepository interface {
	Authenticate(token string) error
}
