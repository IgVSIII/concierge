FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the binary.
RUN go build -o bot ./cmd/bot/main.go

FROM alpine

COPY --from=builder /app/ /app/

WORKDIR /app

CMD ["./bot"]