# Use the latest version of the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the Docker container to /app
WORKDIR /app

# Copy the rest of the code from your local directory to the current directory in the Docker container
COPY . .

# Build the Go application. The resulting binary will be named "main"
RUN go build -o main .

# Inform Docker that the container is listening on port 8080 at runtime
EXPOSE 8080

# Specify the command to run when the Docker container starts up ("/app/main")
ENTRYPOINT [ "/app/main" ]