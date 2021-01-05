# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.14.1-alpine3.11 as builder

ENV GO111MODULE=on

# The latest alpine image don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git

WORKDIR /go/src/goproject
COPY go.sum go.mod /go/src/goproject/

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . /go/src/goproject/

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/goproject ./cmd/app/

FROM alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash

COPY --from=builder /go/src/goproject/bin/goproject /

EXPOSE 8080

ENTRYPOINT ["/goproject"]