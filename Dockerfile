# --- Build Stage ---
# Use a specific Go version as the build environment
FROM golang:1.24.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's layer caching
# This means dependencies are only re-downloaded if go.mod or go.sum change
COPY go.mod .
COPY go.sum .

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a statically linked binary, suitable for scratch/alpine images
# GOOS=linux ensures the binary is built for a Linux environment
# -o /app/backend specifies the output path and name of the executable
# ./main.go specifies the entry point for the build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend ./main.go

# --- Run Stage ---
# Use a minimal base image for the final application
# Alpine is a very small Linux distribution, ideal for production images
FROM alpine:latest

# Set the working directory for the running application
WORKDIR /app

# Copy the compiled binary from the 'builder' stage
COPY --from=builder /app/backend .

# Expose the port your application listens on
# Assuming your Go API service listens on port 8080
EXPOSE 8080

# Define the command to run your executable when the container starts
ENTRYPOINT ["/app/backend"]
