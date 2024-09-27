#!/usr/bin/env bash
set -eux
pushd $(dirname "${BASH_SOURCE[0]}")

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chests/chests.proto