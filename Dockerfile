# Variable Go_Version
ARG GO_VERSION=1.17

# Pull golang-alpine image 
FROM golang:${GO_VERSION}-alpine AS builder

# Environment variable
RUN go env -w GOPROXY=direct