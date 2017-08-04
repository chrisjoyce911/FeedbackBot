#!/bin/sh

printf "\033[34;1m🐘\033[0m  "
echo "Building mybot"
go build -v .

#printf "\033[34;1m🐘\033[0m  "
#echo "Building mybot for docker"
#SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mydockerbot .

#docker build -t mydockerbot -f Dockerfile.scratch .
