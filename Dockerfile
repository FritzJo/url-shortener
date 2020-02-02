# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/url-shortener

# Install requirements
RUN go get github.com/gorilla/mux
RUN go get github.com/boltdb/bolt

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /go/bin/main


FROM scratch

# Add binary
COPY --from=builder /go/bin/main /go/bin/main

# Command to run the executable
ENTRYPOINT ["/go/bin/main"]
