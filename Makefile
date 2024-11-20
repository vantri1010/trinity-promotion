# Variables
CMD_DIR := cmd
APP_NAME := trinity
BUILD_OUTPUT := ./$(APP_NAME)
SCRIPT_DIR := scripts

# Default target
all: build

# Run the application
run:
	go run $(CMD_DIR)/main.go

# Build the application
build:
	go build -o $(BUILD_OUTPUT) $(CMD_DIR)/main.go
	@echo "Built $(APP_NAME) at $(BUILD_OUTPUT)"

# Clean the generated binaries
clean:
	@if [ -f $(BUILD_OUTPUT) ]; then \
		rm -f $(BUILD_OUTPUT); \
		echo "Removed $(BUILD_OUTPUT)"; \
	else \
		echo "No binary to clean"; \
	fi

# Generate Swagger documentation
swag:
	@if ! command -v swag &> /dev/null; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	swag init -g $(CMD_DIR)/main.go -o ./docs
	@echo "Swagger documentation generated in ./docs"

# Run all steps using the script
run-all:
	chmod +x $(SCRIPT_DIR)/run-all.sh
	bash $(SCRIPT_DIR)/run-all.sh

# Show this help message
help:
	@echo "Makefile for $(APP_NAME)"
	@echo
	@echo "Usage:"
	@echo "  make all          Build the application (default target)"
	@echo "  make run          Run the application"
	@echo "  make build        Build the application"
	@echo "  make clean        Clean the generated binaries"
	@echo "  make swag         Generate Swagger documentation"
	@echo "  make run-all      Run all steps using the script"
	@echo "  make help         Show this help message"
