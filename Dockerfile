FROM golang:1.26rc3-alpine

RUN apk add --no-cache git bash make

WORKDIR /go/src/app/
COPY . /go/src/app/
RUN go get
