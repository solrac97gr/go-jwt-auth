package application

import "github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Struct(s ports.Evaluable) error {
	return s.Validate()
}