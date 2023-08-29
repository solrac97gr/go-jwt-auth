# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOINST=$(GOCMD) install

MODULE_NAME=github.com/solrac97gr/go-jwt-auth

# Binary name
BINARY_NAME=main

# Set environment file
set-env:
	@cp ./config/env.example.json ./config/env.json
	@echo "📡 config file generated: ./config/env.json"

# Build
build:
	@$(GOBUILD) -o $(BINARY_NAME) ./cmd/http
	@echo "📦 Build Done"

# Clean
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -f test.out
	@echo "🧹 Program removed"

# Test
test:mocks
	@$(GOTEST) -v ./... >> test.out
	@echo "🧪 Test Completed"

# Lint
lint:
	@golangci-lint run
	@echo "🔦 Code Linted"

# Download dependencies
deps:
	@$(GOINST) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	@$(GOINST) go.uber.org/mock/gomock@latest
	@$(GOINST) github.com/swaggo/swag/cmd/swag@latest
	@$(GOMOD) download
	@echo "⬇️ Dependencies downloaded"

# Build and install
install: deps build
	@echo "🛡️ Successfully Installed"

# Generate the doc
doc:
	@$(GOINST) github.com/swaggo/swag/cmd/swag@latest 
	@swag init --parseDependency=true -g cmd/http/main.go >> output.out
	@rm output.out
	@echo "📓 Docs Generated"


# Generate Coverage Report
cover: mocks
	@$(GOTEST) -coverprofile=coverage.out ./...
	@$(GOTOOL) cover -html=coverage.out
	@rm coverage.out
	@echo "🎯 Cover profile generated"

# Build and run
run: doc build
	@echo "🚀 Running App"
	@./$(BINARY_NAME)

mocks:
	@echo "🔧 Generating mocks"
	@./scripts/generate-mocks.sh