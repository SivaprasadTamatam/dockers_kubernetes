# Use the official Go image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application and print output
RUN go build -o main . 2>&1

# Expose port 8081 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]

