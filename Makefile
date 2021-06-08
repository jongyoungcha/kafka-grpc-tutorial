.PHONY: all dep

CC=${GOROOT}/bin/go
MOD=${GOPATH}/pkg/mod
PWD=$(shell pwd)

all: dep build

build:
	protoc \
	-I$(PWD) \
	--go_out=. \
	--go_opt paths=source_relative \
	protocols/message.proto;

idl:

dep:
	$(CC) mod tidy
