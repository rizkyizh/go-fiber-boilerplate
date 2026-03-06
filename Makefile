.PHONY: run build lint swagger test tidy seed

## run: Start the server with hot reload (requires air)
run:
	air

## build: Compile the application binary
build:
	go build -o ./tmp/main .

## lint: Run golangci-lint
lint:
	golangci-lint run

## swagger: Regenerate Swagger documentation
swagger:
	swag init

## test: Run all unit tests
test:
	go test ./... -v

## test/coverage: Run tests with coverage report
test/coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

## tidy: Tidy go modules
tidy:
	go mod tidy

## help: Show this help message
help:
	@echo "Available targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'
