# Use the official Go image as the base image
FROM golang:3.19-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/app .

# Expose port 4000
EXPOSE 4000

# Command to run the executable
CMD ["./app"]