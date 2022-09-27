# Build stage
FROM golang:1.18.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8081
CMD ["/app/main"]