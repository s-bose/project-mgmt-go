FROM golang:1.23-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
COPY ./wait-for-it.sh ./

RUN go mod download

COPY . .


ENTRYPOINT [ "air" ]




