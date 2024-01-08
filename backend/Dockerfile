# Use the official Golang image
FROM golang:latest

# Set the working directory
WORKDIR /home/backend

# Copy the Go modules files
COPY go.mod ./
COPY go.* ./ 

# Download dependencies
RUN go mod download

# Copy the entire application
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Use go run for live-reloading
CMD ["go", "run", "main.go"]
