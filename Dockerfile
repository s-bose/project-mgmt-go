FROM golang:1.20-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


ENTRYPOINT [ "air" ]




