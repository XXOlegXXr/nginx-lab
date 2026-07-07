FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app_bin main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/app_bin .
CMD ["./app_bin"]
