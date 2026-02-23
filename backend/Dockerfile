FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./
EXPOSE 8080

CMD ["./main"]
