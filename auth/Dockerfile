FROM golang AS builder
WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0
RUN ["make", "auth"]

FROM scratch

WORKDIR /app
COPY --from=builder /app/auth ./

CMD ["./auth"]

EXPOSE 6969
