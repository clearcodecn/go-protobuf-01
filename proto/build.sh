#!bin/bash

protoc -I/usr/local/include -I. --go_out=. --go-grpc_out=. ./*.proto
mv grpc-starter/proto/*.go . && rm -rf grpc-starter