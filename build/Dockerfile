# Build stage
FROM golang:1.21-alpine as build

# if we are in just '/'' then there are issues.
WORKDIR /app

# Install certificates
RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
    update-ca-certificates

COPY . .

# Build binary
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main cmd/main.go

# App stage (finaly image)
FROM scratch

# Copy all certificates to final image
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

# Copy binary from build
COPY --from=build /app/main .

# Expose port 8080 and run binary
EXPOSE 8080
CMD ["./main"]