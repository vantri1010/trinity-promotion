# Base image for Go
FROM golang:1.22-bullseye AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the project files to the container
COPY . .

# Build the application
RUN make build

# Use a lightweight image for production
FROM debian:bullseye-slim AS dev

# Set the working directory in the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/trinity .

# Copy .env file into container
#COPY --from=builder /app/.env .
COPY .env /app/

# Make sure it executable
RUN chmod +x /app/trinity

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./trinity"]
