.PHONY: run

pkg/rpc/service.pb.go: proto/service.proto
	./scripts/build-proto.sh $<

run:
	go build -o bin/smkvs ./cmd/smkvs/
	./scripts/run.sh

build:
	go build -o bin/smkvs ./cmd/smkvs/

proto: pkg/rpc/service.pb.go