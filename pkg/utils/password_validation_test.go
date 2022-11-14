package utils_test

import (
	"github.com/solrac97gr/go-jwt-auth/pkg/utils"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := map[string]struct {
		password string
		want     bool
	}{
		"valid": {
			password: "Aa1!Password",
			want:     true,
		},
		"invalid": {
			password: "Aa1Password",
			want:     false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := utils.IsValid(tt.password); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
