# Dockerfile References: https://docs.docker.com/engine/reference/builder/

## Build backend
# Start from the latest golang base image
FROM golang:1.14.0-buster as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/url-shortener

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install requirements
RUN go mod download

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -tags netgo -a -o urlserver

## Build website
FROM node:stretch-slim as web-builder

WORKDIR /app

# install app dependencies
COPY react_src/ ./
RUN yarn
RUN yarn build

# Create final image
FROM scratch

# Copy binary and static files
WORKDIR /app
COPY --from=builder /go/src/url-shortener /app
COPY --from=web-builder /app/build /app
# Command to run the executable
ENTRYPOINT ["./urlserver"]
