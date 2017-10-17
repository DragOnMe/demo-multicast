#!/usr/bin/env bash

gox -osarch="linux/amd64"

docker build -t drlee001/demo-multicast .
docker push drlee001/demo-multicast
