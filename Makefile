.PHONY: default build test

default: build

build: proto

test: build

PROTO_GEN = lightsteppb/lightstep_carrier.pb.go collectorpb/collector.pb.go

.PHONY: proto clean-proto

clean-proto:
	@rm -f $(PROTO_GEN)

proto: $(PROTO_GEN)

collectorpb/collector.pb.go: collector.proto
	docker run --rm -v $(shell pwd):/input:ro -v $(shell pwd)/collectorpb:/output \
	  lightstep/gogoprotoc:latest \
		protoc -I/input/third_party/googleapis --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:/output --proto_path=/input:. /input/collector.proto

lightsteppb/lightstep_carrier.pb.go: lightstep_carrier.proto
	docker run --rm -v $(shell pwd):/input:ro -v $(shell pwd)/lightsteppb:/output \
	  lightstep/gogoprotoc:latest \
	  	protoc --gogofaster_out=plugins=grpc:/output --proto_path=/input:. /input/lightstep_carrier.proto

test: $(PROTO_GEN) proto_test.go
	go test -v .
