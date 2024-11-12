# Use the official Golang image as the base
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code to the container
COPY . .

# Build the Go application
RUN go mod init go_docker_app && go build -o app main.go

# Expose the port on which the app will run
EXPOSE 3000

# Run the app
CMD ["./app"]