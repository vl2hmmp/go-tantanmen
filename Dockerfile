FROM golang:1.15-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache git \
 && go build -o main

CMD ["./main"]