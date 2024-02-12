FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o online-store

EXPOSE 8080

CMD ["./online-store"]
