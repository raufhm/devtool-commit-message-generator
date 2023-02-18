# Use an official Go runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o myapp

# Start
CMD ["/app"]
