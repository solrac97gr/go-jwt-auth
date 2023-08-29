package application_test

import (
	"github.com/solrac97gr/go-jwt-auth/internal/user/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/validator/application"
	"github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestValidator(t *testing.T) {
	test := map[string]struct {
		input ports.Evaluable
		want  error
	}{
		"Valid": {
			input: &models.User{
				ID:        primitive.ObjectID{},
				Email:     "test@gmail.com",
				Password:  "test12345@P",
				CreatedAt: time.Now(),
			},
		},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			validator := application.NewValidator()
			err := validator.Struct(tc.input)
			if err != nil {
				t.Errorf("got %v, want %v", err, tc.want)
			}
		})
	}
}
