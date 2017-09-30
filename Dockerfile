# vim: ft=dockerfile :
FROM golang:1.9

RUN go get -u github.com/codegangsta/gin

RUN mkdir -p /go/src/github.com/agrea/watchtower
WORKDIR /go/src/github.com/agrea/watchtower

CMD ["true"]
