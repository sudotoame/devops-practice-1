FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY src/ .

RUN go mod tidy
RUN go build -o service .

FROM alpine:3.19 AS final
WORKDIR /app

COPY --from=builder /app/service .
EXPOSE 9091
CMD ["./service"]
