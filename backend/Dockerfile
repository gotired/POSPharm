FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/main ./cmd/main.go

FROM alpine:latest

WORKDIR /build

COPY --from=builder /build/main .

EXPOSE 8080

CMD ["./main"]

# CMD ["tail", "-f", "/dev/null"]