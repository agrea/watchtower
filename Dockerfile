# vim: ft=dockerfile :
FROM golang:1.9

RUN go get -u github.com/tockins/realize

RUN mkdir -p /go/src/github.com/agrea/watchtower
WORKDIR /go/src/github.com/agrea/watchtower

EXPOSE 3000

CMD ["true"]
