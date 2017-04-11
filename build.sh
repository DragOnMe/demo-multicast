#!/usr/bin/env bash

gox -osarch="linux/amd64"

docker build -t docker.io/mjelen/demo-multicast .
docker push docker.io/mjelen/demo-multicast
