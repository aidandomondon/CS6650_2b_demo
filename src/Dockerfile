# syntax=docker/dockerfile:1

# Use the Golang base image as a base.
FROM golang:1.27.1

# Execute all following commands while inside of the /app directory
WORKDIR /app

# Copy go.mod and go.sum from the directory this `Dockerfile` is in
# into the /app directory this `Dockerfile` is building.
COPY go.mod go.sum ./

# Download the go program's dependencies
RUN go mod download

# Copy all .go files from the directory this `Dockerfile` is in into /app
COPY *.go ./

# Cross compile the go program
# Compiling for AMD architecture because that is the EC2's
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /online-store

EXPOSE 8080

# When this image is used to start a new container,
# start the online store
CMD ["/online-store"]