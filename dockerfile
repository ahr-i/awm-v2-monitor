# Use an official Golang image as a parent image
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Run the binary program produced by `go build`
CMD ["./main"]
