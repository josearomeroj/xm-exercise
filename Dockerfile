FROM golang:1.20.3-alpine3.17 AS build

WORKDIR /app

COPY . .
RUN go build -o rx-server ./cmd/xm-server

FROM alpine:3.17.3

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app .

ENTRYPOINT ["./rx-server"]