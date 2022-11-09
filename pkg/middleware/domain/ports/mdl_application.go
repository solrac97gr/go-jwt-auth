package ports

type MiddlewareApplication interface {
	Authenticate(token string) (code int, err error)
}
