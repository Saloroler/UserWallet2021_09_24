FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app/cmd

RUN go build

EXPOSE 8080
ENTRYPOINT ./cmd
