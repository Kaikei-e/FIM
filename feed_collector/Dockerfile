FROM golang:1.23.3 AS builder

ADD . /app
WORKDIR /app

RUN go build -o feed_collector main.go

FROM scratch

COPY --from=builder /app/feed_collector /app/feed_collector
RUN chmod +x /app/feed_collector

ENTRYPOINT ["./feed_collector"]