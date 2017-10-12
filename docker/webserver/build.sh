#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t kylews/testserver .
docker push kylews/testserver
go clean
