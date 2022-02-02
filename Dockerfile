# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

RUN apk add build-base

COPY go.mod ./
COPY go.sum ./

RUN go mod download
# RUN go get github.com/mattn/go-sqlite3

COPY *.go ./
COPY view ./view
# COPY art.db ./

RUN go build -o /videowall

EXPOSE 1323

CMD [ "/videowall" ]