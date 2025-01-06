# Description: Makefile for apiblueprint gateway project

# Initialize the project
init:
	@go mod tidy
	@go install github.com/swaggo/swag/cmd/swag@latest

# Run the application
run:
	@go run cmd/gateway/main.go

# Seed data
seed:
	@go run cmd/seed/main.go

# Run stack.yaml with docker compose
stackup:
	@docker compose -f stack.yaml up -d

# Stop stack.yaml with docker compose
stackdown:
	@docker compose down -f stack.yaml

.PHONY: swagger
swagger:
	@swag fmt -g internal/gateway/v1/v1.go
	@swag init --parseDependency -g internal/gateway/v1/v1.go -o swagger/v1

.PHONY: init run proto stackup stackdown swagger