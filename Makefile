# Note the protoc-generated code in this directory is a sanity check
# of the protos here, not generated code for import.
#
.PHONY: default build test

default: build

build: test

PWD = $(shell pwd)

TEST_OUTPUT_PREFIX = test
TEST_PROTO_GEN = \
	$(TEST_OUTPUT_PREFIX)/lightsteppb/lightstep_carrier.pb.go \
	$(TEST_OUTPUT_PREFIX)/collectorpb/collector.pb.go

$(TEST_OUTPUT_PREFIX)/collectorpb/collector.pb.go: collector.proto
	mkdir -p $(TEST_OUTPUT_PREFIX)/collectorpb && cd test && \
	docker run --rm -v $(PWD):/input:ro -v $(PWD)/$(TEST_OUTPUT_PREFIX)/collectorpb:/output \
	  lightstep/gogoprotoc:latest \
		protoc -I/input/third_party/googleapis --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:/output --proto_path=/input:. /input/collector.proto

$(TEST_OUTPUT_PREFIX)/lightsteppb/lightstep_carrier.pb.go: lightstep_carrier.proto
	mkdir -p $(TEST_OUTPUT_PREFIX)/lightsteppb && cd test && \
	docker run --rm -v $(PWD):/input:ro -v $(PWD)/$(TEST_OUTPUT_PREFIX)/lightsteppb:/output \
	  lightstep/gogoprotoc:latest \
	  	protoc --gogofaster_out=plugins=grpc:/output --proto_path=/input:. /input/lightstep_carrier.proto

test: $(TEST_PROTO_GEN) test/proto_test.go
	go test -v ./test
