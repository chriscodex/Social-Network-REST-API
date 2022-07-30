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

# Copy the rest of the files
COPY [".","./"]

# Build the aplication (We indicate to not use C++ compiler)
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /REST-API-WebSocket

FROM scratch as runner 

# Copy the certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy .env file
COPY .env ./

# Copy the executable
COPY --from=builder /REST-API-WebSocket /REST-API-WebSocket

# Expose the port 5050
EXPOSE 5050

# Run the API executable
ENTRYPOINT [ "/REST-API-WebSocket" ]