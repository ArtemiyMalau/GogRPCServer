# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY src/server/go.mod ./
COPY src/server/go.sum ./
RUN go mod download

COPY src/server/* ./

RUN go build -o /grpc_server

EXPOSE 8000

CMD ["/grpc_server"]
