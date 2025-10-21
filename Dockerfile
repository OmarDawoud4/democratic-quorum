FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8001

okCMD ["./app"]
