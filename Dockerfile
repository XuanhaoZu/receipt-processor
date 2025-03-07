# Use the official Go image
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o receipt-processor

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./receipt-processor"]
