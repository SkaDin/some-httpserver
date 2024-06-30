FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/some-server/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["./main"]