FROM golang:1.19.3-alpine3.16

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

EXPOSE 8080