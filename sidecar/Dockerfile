FROM golang:1.7-alpine

RUN apk add git --no-cache && \
    go get github.com/micro/micro && \
    apk del git

CMD micro --enable_stats sidecar

EXPOSE 8080
