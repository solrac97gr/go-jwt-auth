package server_test

import (
	"errors"
	"testing"

	"github.com/solrac97gr/go-jwt-auth/pkg/config/domain/models"
	"github.com/solrac97gr/go-jwt-auth/pkg/mocks"
	"github.com/solrac97gr/go-jwt-auth/pkg/server"
	"go.uber.org/mock/gomock"
)

func TestServer_Run(t *testing.T) {

	tests := map[string]struct {
		port   string
		setup  func(handlers *mocks.MockMiddlewareHandlers, user *mocks.MockUserHandlers, config *mocks.MockConfigApplication)
		assert func(t *testing.T, err error)
	}{
		"should return error if config is not found": {
			port: "8080",
			setup: func(handlers *mocks.MockMiddlewareHandlers, user *mocks.MockUserHandlers, config *mocks.MockConfigApplication) {
				config.EXPECT().GetConfig().Return(nil, errors.New("error"))
			},
			assert: func(t *testing.T, err error) {
				if err == nil {
					t.Error("expected error, got nil")
				}
			},
		},
		"should return success if server not fails": {
			port: "8080",
			setup: func(handlers *mocks.MockMiddlewareHandlers, user *mocks.MockUserHandlers, config *mocks.MockConfigApplication) {
				config.EXPECT().GetConfig().Return(&models.Config{
					App:      &models.App{Name: "test"},
					Database: &models.Database{},
					Server:   &models.Server{Port: "8080"},
					JWT:      &models.JWT{},
				}, nil)

				handlers.EXPECT().Authenticate().Return(nil)
			},
			assert: func(t *testing.T, err error) {
				if err != nil {
					t.Error("expected error, got nil")
				}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			handlers := mocks.NewMockMiddlewareHandlers(ctrl)
			user := mocks.NewMockUserHandlers(ctrl)
			config := mocks.NewMockConfigApplication(ctrl)
			if tt.setup != nil {
				tt.setup(handlers, user, config)
			}
			s := server.NewServer(handlers, user, config)
			err := s.Run(tt.port)
			if tt.assert != nil {
				tt.assert(t, err)
			}
			err = s.Stop()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
