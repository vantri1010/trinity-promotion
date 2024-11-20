#!/bin/bash

# Define the output binary path (current directory or project root)
OUTPUT_BINARY="./trinity"

# Step 1: Check if `swag` is installed and install if necessary
if ! command -v swag &> /dev/null; then
    echo "swag not found. Installing..."
    go install github.com/swaggo/swag/cmd/swag@latest || { 
        echo "Failed to install swag"; exit 1; 
    }
else
    echo "swag is already installed."
fi

# Step 2: Install dependencies
echo "Running go mod tidy..."
go mod tidy || { 
    echo "Failed to tidy Go modules"; exit 1; 
}

# Step 3: Build the application
echo "Building the application..."
if ! go build -o "$OUTPUT_BINARY" cmd/main.go; then
    echo "Failed to build the application."
    exit 1
fi

# Step 4: Generate Swagger documentation
echo "Generating Swagger documentation..."
if ! swag init -g cmd/main.go -o docs; then
    echo "Failed to generate Swagger documentation."
    exit 1
fi

# Step 5: Run the built application
echo "Running the built application..."
if ! "$OUTPUT_BINARY"; then
    echo "Failed to run the application."
    exit 1
fi

