package application

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	validator "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
	"go.uber.org/zap"
	"time"
)

type Logger struct {
	repo ports.LoggerRepository
	val  validator.ValidatorApplication
	zap  *zap.SugaredLogger
}

var _ ports.LoggerApplication = (*Logger)(nil)

func NewLogger(repo ports.LoggerRepository, val validator.ValidatorApplication) *Logger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return &Logger{
		repo: repo,
		val:  val,
		zap:  sugar,
	}
}

func (l Logger) Close() error {
	return l.zap.Sync()
}

func (l Logger) Debug(msg string, args ...interface{}) {
	l.zap.Debug(msg, args)
	log := new(models.Log)
	log.Level = "debug"
	log.Message = msg
	log.CreatedAt = time.Now()

	err := l.repo.Save(log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Info(msg string, args ...interface{}) {
	l.zap.Info(msg, args)
	log := new(models.Log)
	log.Level = "info"
	log.Message = msg
	log.CreatedAt = time.Now()

	err := l.repo.Save(log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Warn(msg string, args ...interface{}) {
	l.zap.Warn(msg, args)
	log := new(models.Log)
	log.Level = "warn"
	log.Message = msg
	log.CreatedAt = time.Now()

	err := l.repo.Save(log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}

func (l Logger) Error(msg string, args ...interface{}) {
	l.zap.Error(msg, args)
	log := new(models.Log)
	log.Level = "error"
	log.Message = msg
	log.CreatedAt = time.Now()

	err := l.repo.Save(log)
	if err != nil {
		l.zap.Error(err.Error())
	}
}
