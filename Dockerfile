FROM golang:1.19.4-alpine3.17 AS builder
MAINTAINER Son Nguyen<son.nguyen.2@neyu.co>

RUN apk add build-base

# Build
RUN mkdir /build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main

# Run App

FROM alpine:3.8
MAINTAINER Son Nguyen<son.nguyen.2@neyu.co>
RUN mkdir /app

RUN apk add tzdata

COPY --from=builder /build/main /app/main
COPY --from=builder /build/config /app/config
COPY --from=builder /build/public /app/public
COPY --from=builder /build/view /app/view
COPY --from=builder /build/files /app/files

RUN chmod 711 /app/main \
    && rm -rf /var/cache/apk/* \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

EXPOSE 8080

WORKDIR /app

ENTRYPOINT /app/main