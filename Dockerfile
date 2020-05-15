# Dockerfile References: https://docs.docker.com/engine/reference/builder/

## Build backend
# Start from the latest golang base image
FROM golang:1.14.0-buster as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/url-shortener

# Copy the source from the current directory to the Working Directory inside the container
COPY  ["*.go", "go.sum", "go.mod", "./"]

# Install requirements
RUN go mod download

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -tags netgo -a -o urlserver

# Create final image
FROM scratch

# Copy binary and static files
WORKDIR /app
COPY --from=builder /go/src/url-shortener /app
COPY ./static /app/static

# Command to run the executable
ENTRYPOINT ["./urlserver"]
