#!/bin/bash
GOPATH=`pwd`
GOARCH=amd64 GOOS=linux go build -o src/github.com/pkar/transmogrify/bin/transmogrify_linux src/github.com/pkar/transmogrify/cmd/main.go
GOARCH=amd64 GOOS=darwin go build -o src/github.com/pkar/transmogrify/bin/transmogrify_mac src/github.com/pkar/transmogrify/cmd/main.go
