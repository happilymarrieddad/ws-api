FROM golang:1.21 AS builder
WORKDIR /app

COPY . .
RUN go build -o /app/api cmd/main.go

FROM alpine:edge
WORKDIR /app

RUN apk add --no-cache libc6-compat

COPY --from=builder /app/api /app/api

CMD ["/app/api"]