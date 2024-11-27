FROM golang:1.23.3 AS builder

ADD . /app
WORKDIR /app

RUN go build -o federation_orchestrator main.go

FROM scratch

COPY --from=builder /app/federation_orchestrator /app/federation_orchestrator
RUN chmod +x /app/federation_orchestrator

ENTRYPOINT ["./federation_orchestrator"]