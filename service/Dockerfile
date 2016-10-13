FROM golang:1.7-alpine

ADD . /go/src/github.com/balashVI/go-micro-test/service

RUN apk add git --no-cache && \
    go get -u github.com/kardianos/govendor && \
    cd /go/src/github.com/balashVI/go-micro-test/service/app && \
    govendor install +vendor,^program && \
    go build -o /go/bin/service && \
    apk del git

CMD /go/bin/service

EXPOSE 8080
