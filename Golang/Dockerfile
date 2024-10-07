# Use the official Golang image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o app main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
