# Build stage
FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o go-http-server main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go-http-server .
COPY static/ ./static/

RUN adduser -D appuser
USER appuser

EXPOSE 3000
CMD ["./go-http-server"]