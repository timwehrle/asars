# Variables
APP_NAME=asars
BUILD_DIR=

# Commands
GOCMD := go 
GOTEST := $(GOCMD) test
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOCLEAN := $(GOCMD) clean
STATICCHECK := staticcheck

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Run static analysis
.PHONY: lint
lint:
	@echo "Running static analysis..."
	$(STATICCHECK) ./...