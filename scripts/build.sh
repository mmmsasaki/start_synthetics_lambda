#!/bin/sh
go get

CGO_ENABLED=0 GOOS=linux go build -a -v -installsuffix cgo -o bin/app main.go
