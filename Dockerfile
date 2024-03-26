# Use golang 1.22 base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application when the container starts
CMD ["go", "run", "main.go"]
