# Start from a minimal Go image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/app/main"]
