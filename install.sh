#!/bin/sh
GOPATH=$PWD
export GOPATH
mkdir ./bin
go build -i -o $PWD/bin/storage ./src/storage.go