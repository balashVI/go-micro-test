FROM golang:1.7-alpine

ADD . /go/src/github.com/balashVI/go-micro-test/service-client

RUN apk add git --no-cache && \
    go get -u github.com/kardianos/govendor && \
    cd /go/src/github.com/balashVI/go-micro-test/service-client/app && \
    govendor install +vendor,^program && \
    go build -o /go/bin/service-client && \
    apk del git

CMD /go/bin/service-client

EXPOSE 8080
