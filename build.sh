#!/usr/bin/env bash

# Copy demo-multicast go binary
cp $GOBIN/multicast ./multicast_linux_amd64

gox -osarch="linux/amd64"

docker build -t drlee001/demo-multicast .
docker push drlee001/demo-multicast
