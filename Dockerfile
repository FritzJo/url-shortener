# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.14.0-buster

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/url-shortener

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install requirements
RUN go mod download

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o urlserver . 

# Command to run the executable
ENTRYPOINT ["./urlserver"]
