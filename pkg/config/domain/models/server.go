package models

import customErr "github.com/solrac97gr/go-jwt-auth/pkg/custom-errors"

type Server struct {
	Port string `json:"port"`
}

func (s *Server) Validate() error {
	if s.Port == "" {
		return customErr.ErrServerPortRequired
	}
	return nil
}
