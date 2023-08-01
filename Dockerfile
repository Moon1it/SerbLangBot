# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy all the files from the current directory (where the Dockerfile is located) into the container
COPY . .

# Build your Golang application
RUN go build cmd/main.go

# Set environment variables for the Telegram API key and bot token
ENV MONGO_STRING=MONGO_STRING
ENV TELEGRAM_TOKEN=TELEGRAM_TOKEN

# Run the application when the container starts
CMD ["./main"]
