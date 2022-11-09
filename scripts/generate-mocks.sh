mockgen -destination=pkg/mocks/mock_user_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports UserApplication &&
mockgen -destination=pkg/mocks/mock_user_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports UserRepository &&
mockgen -destination=pkg/mocks/mock_user_handlers.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/internal/user/domain/ports UserHandlers &&
mockgen -destination=pkg/mocks/mock_mdl_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports MiddlewareApplication &&
mockgen -destination=pkg/mocks/mock_mdl_handlers.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports MiddlewareHandlers &&
mockgen -destination=pkg/mocks/mock_mdl_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/middleware/domain/ports MiddlewareRepository &&
mockgen -destination=pkg/mocks/mock_validator_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports ValidatorApplication &&
mockgen -destination=pkg/mocks/mock_validator_evaluable.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/validator/domain/ports Evaluable &&
mockgen -destination=pkg/mocks/mock_config_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports ConfigApplication &&
mockgen -destination=pkg/mocks/mock_config_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/config/domain/ports ConfigRepository &&
mockgen -destination=pkg/mocks/mock_logger_application.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports LoggerApplication &&
mockgen -destination=pkg/mocks/mock_logger_repository.go -package=mocks --build_flags=--mod=mod github.com/solrac97gr/go-jwt-auth/pkg/logger/domain/ports LoggerRepository &&
echo "Mocks generated successfully 🚀[pkg/mocks]"