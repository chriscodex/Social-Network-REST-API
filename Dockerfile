# Variable Go_Version
ARG GO_VERSION=1.17

# Pull golang-alpine image 
FROM golang:${GO_VERSION}-alpine AS builder

# Environment variable
RUN go env -w GOPROXY=direct

# Install git
RUN apk add --no-cache git

# Security certificates
RUN apk --no-cache add ca-certificates && update-ca-certificates