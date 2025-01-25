FROM golang:1.23.5 AS builder

ADD ./feed_collector /app/src
WORKDIR /app/src

RUN go build -o feed_collector main.go

FROM scratch

COPY --from=builder /app/src /app
RUN chmod +x /app/feed_collector

ENTRYPOINT ["./feed_collector"]