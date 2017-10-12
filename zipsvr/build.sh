#!/usr/bin/env bash
set -e
export GOPATH=/home/kbox/Documents/kstuff/Go
CGO_ENABLED=0 go build -a
docker build -t kylews/zipsvr .
docker push kylews/zipsvr
go clean
