# Start from the official Go image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory into the working directory
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the port on which the gRPC server will run
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
