FROM golang:1.21-alpine as build
WORKDIR /app

COPY   ./ /app

RUN  GO111MODULE="on" go mod vendor
RUN  CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main cmd/main.go

EXPOSE 8080
CMD ["./main"]