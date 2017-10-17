#!/usr/bin/env bash

# Build & Copy demo-multicast go binary
go install $GOPATH/multicast.go
cp $GOBIN/multicast ./multicast_linux_amd64

gox -osarch="linux/amd64"

docker build -t drlee001/demo-multicast .
docker push drlee001/demo-multicast
