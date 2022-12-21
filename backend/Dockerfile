FROM golang AS builder
WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0
RUN ["make", "backend"]

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/backend ./

ENV GIN_MODE=release
CMD ["./backend"]

EXPOSE 8080
