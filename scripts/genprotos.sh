#!/usr/bin/env bash

set -euo pipefail

API_DIR=${1:-$PWD/api}
cd api

protos=$(find . -type f -name '*.proto')

for proto in $protos; do
	echo " - $proto";
	protoc \
		--proto_path=${API_DIR} \
		--proto_path=. \
		--proto_path=$GOPATH/src/github.com/ligato/vpp-agent/vendor \
		--gogo_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:$GOPATH/src \
		"${API_DIR}/$proto";
done
