package models

import customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"

type Database struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (d *Database) Validate() error {
	if d.Name == "" {
		return customErr.ErrDatabaseNameRequired
	}
	if d.URL == "" {
		return customErr.ErrDatabaseURLRequired
	}
	return nil
}
