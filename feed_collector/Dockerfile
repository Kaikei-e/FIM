FROM golang:1.23.5 AS builder

ADD ./feed_collector /app/src
WORKDIR /app/src

RUN go build -o feed_collector main.go

FROM alpine:latest

COPY --from=builder /app/src/feed_collector /app/feed_collector
RUN chmod +x /app/feed_collector

ENTRYPOINT ["/app/feed_collector"]