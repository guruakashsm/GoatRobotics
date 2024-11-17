#!/bin/bash

#FED Build script.

buildTime=$( date --utc +%FT%T.%3NZ )
cId=$(git rev-parse --short=7 HEAD)
version=$(git describe --tags --abbrev=0)-${cId}
Version="fednow/rpc.Version=${version}"
BuildTime="fednow/rpc.BuildTime=${buildTime}"
CommitID="fednow/rpc.CommitID=${cId}"
Ldflags="-X '${Version}' -X '${BuildTime}' -X '${CommitID}'"
Command='env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -ldflags='
# Command='env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags='
# Need to remove [CGO_ENABLED=0] in above line once jenkin setup finished.

echo 'Starting binary build'
echo ${Command}"${Ldflags}"
 ${Command}"${Ldflags}"
echo 'Binary build completed.'