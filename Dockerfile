FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /builder/src
COPY . .
ENV GO111MODULE=on
RUN go build -ldflags="-w -s" -o  /builder/src/app

FROM alpine:3.20

WORKDIR /usr/local/bin
COPY --from=builder /builder/src/app ./app
COPY config.toml ./config.toml
RUN chmod +x ./app
ENV GIN_MODE=release
EXPOSE 8080
CMD ["./app"]