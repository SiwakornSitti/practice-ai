.PHONY: run build test clean lint swagger

# ... (rest of file)

lint:
	@echo "Linting..."
	@golangci-lint run ./...

swagger:
	@echo "Generating swagger docs..."
	@$(GOCMD) install github.com/swaggo/swag/cmd/swag@latest
	@swag init -d ./cmd/api,./internal/user/delivery -o ./cmd/api/docs
