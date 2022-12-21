FROM golang AS builder
WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0
RUN ["make", "backend"]

FROM scratch

WORKDIR /app

COPY --from=builder /app/assets ./
COPY --from=builder /app/backend ./

ENV GIN_MODE=release
CMD ["./backend"]

EXPOSE 8080
