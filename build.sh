#!/bin/bash

go mod tidy

Command='env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -ldflags='

echo 'Starting binary build'
echo ${Command}"${Ldflags}"
 ${Command}"${Ldflags}"
echo 'Binary build completed.'