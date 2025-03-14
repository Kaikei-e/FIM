FROM golang:1.23.5 AS builder

ADD ./federation_orchestrator/federation_orchestrator /app/src
ADD ./buf_schema /app/buf_schema
WORKDIR /app/src

RUN go build -o federation_orchestrator main.go

FROM alpine:latest

COPY --from=builder /app/src/federation_orchestrator /app/federation_orchestrator
RUN chmod +x /app/federation_orchestrator

ENTRYPOINT ["/app/federation_orchestrator"]