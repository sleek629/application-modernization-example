#!/bin/bash -e

PATH=$PATH:$GOPATH/bin
protodir=../../pb

protoc --go_out=plugins=grpc:proto -I $protodir $protodir/word.proto
