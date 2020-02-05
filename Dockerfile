# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.13.7

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/url-shortener

# Install requirements
RUN go get github.com/gorilla/mux
RUN go get github.com/boltdb/bolt

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o urlserver . 

# Command to run the executable
ENTRYPOINT ["./urlserver"]
