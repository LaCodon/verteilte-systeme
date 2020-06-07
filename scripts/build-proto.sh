#!/usr/bin/env bash

proto_file=$1

SRC_DIR=$(pwd)
DST_DIR=${SRC_DIR}

protoc -I=$SRC_DIR --go_out=$DST_DIR --go-grpc_out=$DST_DIR $SRC_DIR/$proto_file