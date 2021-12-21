# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY view ./view

RUN go build -o /videowall

EXPOSE 1323

CMD [ "/videowall" ]