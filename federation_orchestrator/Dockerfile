FROM golang:1.23.5 AS builder

ADD ./federation_orchestrator /app/src
WORKDIR /app/src

RUN go build -o federation_orchestrator main.go

FROM scratch

COPY --from=builder /app/src /app
RUN chmod +x /app/federation_orchestrator

ENTRYPOINT ["./federation_orchestrator"]