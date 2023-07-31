# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy all the files from the current directory (where the Dockerfile is located) into the container
COPY . .

# Build your Golang application
RUN go build cmd/main.go 

# Run the application when the container starts
CMD ["./main"]

