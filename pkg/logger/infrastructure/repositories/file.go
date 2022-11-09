package repositories

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
)

type CSVFile struct {
	path string
}

var _ ports.LoggerRepository = (*CSVFile)(nil)

func NewCSVFile(path string) *CSVFile {
	return &CSVFile{path: path}
}

func (C CSVFile) Save(log models.Log) error {
	//TODO implement me
	panic("implement me")
}
