# Use the official Golang image
FROM golang:latest as builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o main main.go

# Expose the service port
EXPOSE 8001

# Run the application
CMD ["./main"]