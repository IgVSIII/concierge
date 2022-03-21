FROM golang:1.17-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

# Build the binary.
RUN go build -o bot ./cmd/bot/main.go

FROM alpine:last

COPY --from=builder /app/ /app/

CMD ["/app/bot"]