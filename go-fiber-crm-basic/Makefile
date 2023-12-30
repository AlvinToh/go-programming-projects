# Get current directory name based on splitting the directory separator and selecting the last element
service_name := $(lastword $(subst /, ,$(CURDIR)))

# Go tidy, generate swagger and build go service
.PHONY: build
build:
	@echo "Building Service: $(service_name)"
	go mod tidy
	swagger generate spec -o ./swagger/swagger.json
	go build ./cmd/$(service_name)/main.go

.PHONY: run
run: run
	@echo "Running Service: $(service_name)"
	go run ./cmd/$(service_name)/main.go

.PHONY: test
test:
	@echo "Testing Service: $(service_name)"
	go test -v ./...

.PHONY: coverage
coverage:
	@echo "Generating coverage report for Service: $(service_name)"
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Current delete commands are for Windows only
.PHONY: clean
clean:
	@echo "Cleaning Service: $(service_name)"
	go clean
	if exist main.exe del /F /Q main.exe
	if exist coverage.out del /F /Q coverage.out