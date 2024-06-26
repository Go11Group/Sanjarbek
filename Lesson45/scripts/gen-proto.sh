#!/bin/zsh
CURRENT_DIR=$1
PROTO_DIR=${CURRENT_DIR}/protos

rm -rf ${CURRENT_DIR}/library

for x in $(find ${PROTO_DIR}/* -type d); do
  protoc -I=${x} -I=${PROTO_DIR} -I /usr/local/go \
    --go_out=${CURRENT_DIR} --go_opt=paths=source_relative,Mprotos/library/library.proto=path/to/your/package/protos/library \
    --go-grpc_out=${CURRENT_DIR} --go-grpc_opt=paths=source_relative,Mprotos/library/library.proto=path/to/your/package/protos/library \
    ${x}/*.proto
done
