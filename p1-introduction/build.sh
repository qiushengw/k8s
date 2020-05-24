#!/bin/bash
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .
eval $(minikube docker-env)
docker build . -t weather:1.0

