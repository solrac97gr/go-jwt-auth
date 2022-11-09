package repositories

import (
	"encoding/csv"
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports"
	"os"
)

type CSVFile struct {
	path string
}

var _ ports.LoggerRepository = (*CSVFile)(nil)

func NewCSVFile(path string) *CSVFile {
	return &CSVFile{path: path}
}

func (c *CSVFile) Save(log *models.Log) error {
	file, err := os.Create(c.path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Write data to csv file
	err = csvWriter.Write([]string{log.ID, log.Level, log.Message, log.CreatedAt.String()})
	if err != nil {
		return err
	}
	return nil
}
