.PHONY: run build test clean lint

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Main package path
MAIN_PATH=./cmd/api/main.go

# Binary name
BINARY_NAME=app

run:
	@echo "Running $(BINARY_NAME)..."
	@$(GORUN) $(MAIN_PATH)

build:
	@echo "Building $(BINARY_NAME)..."
	@$(GOBUILD) -o bin/$(BINARY_NAME) $(MAIN_PATH)

docker-build:
	@echo "Building docker image..."
	@docker build -t $(BINARY_NAME):latest .

test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf bin/

tidy:
	@echo "Tidying module..."
	@$(GOMOD) tidy

lint:
	@echo "Linting..."
	@golangci-lint run ./...
