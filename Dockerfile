FROM golang:1.25rc2-alpine

RUN apk add --no-cache git bash make

WORKDIR /go/src/app/
COPY . /go/src/app/
RUN go get
