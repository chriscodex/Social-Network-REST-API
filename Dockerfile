# API Service
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

# Working directory
WORKDIR /src

# Copy go.mod and go.sum into /src directory
COPY ["./go.mod","./go.sum","./"]

# Install dependencies
RUN go mod download
