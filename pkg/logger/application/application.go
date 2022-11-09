package application

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	validator "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
)

type Logger struct {
	repo ports.LoggerRepository
	val  validator.ValidatorApplication
}

var _ ports.LoggerApplication = (*Logger)(nil)

func NewLogger(repo ports.LoggerRepository, val validator.ValidatorApplication) *Logger {
	return &Logger{
		repo: repo,
		val:  val,
	}
}

func (l Logger) Close() error {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Debug(msg string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Info(msg string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Warn(msg string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) Error(msg string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}
