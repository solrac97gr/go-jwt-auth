package repositories

import "os"

type JSONRepository struct {
	path string
}

func NewJSONRepository(path string) *JSONRepository {
	return &JSONRepository{
		path: path,
	}
}

func (j *JSONRepository) GetConfigFile() (*os.File, error) {
	return os.Open(j.path)
}
