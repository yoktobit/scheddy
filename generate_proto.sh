#!/bin/sh
PREFIX=github.com/yoktobit/scheddy_backend
protoc --go_out=. --go-grpc_out=. --go_opt=module=$PREFIX --go-grpc_opt=module=$PREFIX proto/event.proto