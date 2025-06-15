FROM golang:1.20-alpine

# Set platform for ARM64
ARG TARGETPLATFORM=linux/arm64
ARG BUILDPLATFORM=linux/arm64

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN GOOS=linux GOARCH=arm64 go build -o c2-server .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./c2-server"] 