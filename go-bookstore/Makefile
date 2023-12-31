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

# Current delete commands for Windows only
.PHONY: clean
clean:
	@echo "Cleaning Service: $(service_name)"
	go clean
	if exist main.exe del /F /Q main.exe
	if exist coverage.out del /F /Q coverage.out

# Deploy DockerFile with environment variables
.PHONY: deploy
deploy:
	@echo "Deploying Service: $(service_name)"
	docker build -t $(service_name):tag_release_deploy .
	docker run --network sql-db-stack_default -d -p 9001:9001 $(service_name):tag_release_deploy

# Destroy Docker service
.PHONY: destroy
destroy:
	@echo "Destroying Service: $(service_name)"
	$(eval CONTAINER_ID := $(shell docker ps -q --filter ancestor=$(service_name):tag_release_deploy))
	docker stop $(CONTAINER_ID)
	docker rm $(CONTAINER_ID)
	docker rmi $(service_name):tag_release_deploy