FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum main.go .
COPY backend ./backend
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
