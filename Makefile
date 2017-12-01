include genproto.mk

PKG_PREFIX = github.com/lightstep/lightstep-tracer-common

PROTO_SOURCES = \
	collector.proto \
	lightstep.proto

TEST_SOURCES = \
	$(GOLANG)/gogo_test.go \
	$(GOLANG)/protobuf_test.go

GOGO_GENTGTS = $(call protos_to_gogo_targets,$(PROTO_SOURCES))
PBUF_GENTGTS = $(call protos_to_protobuf_targets,$(PROTO_SOURCES))

FAKES = \
	golang/gogo/collectorpb/collectorpbfakes/fake_collector_service_client.go \
	golang/protobuf/collectorpb/collectorpbfakes/fake_collector_service_client.go

.PHONY: default build test clean $(GOGO_GENTGTS) $(PBUF_GENTGTS)

default: build

build: test

test: $(GOGO_GENTGTS) $(PBUF_GENTGTS) $(FAKES) $(TEST_SOURCES)
	go test -v ./golang

clean: 
	$(call clean_protoc_targets,$(GOGO_GENTGTS) $(PBUF_GENTGTS))

$(GOGO_GENTGTS): $(GOLANG)-$(GOGO)-%: %.proto
	$(call gen_gogo_target,$<)

$(PBUF_GENTGTS): $(GOLANG)-$(PBUF)-%: %.proto
	$(call gen_protobuf_target,$<)

golang/gogo/collectorpb/collectorpbfakes/fake_collector_service_client.go: golang/gogo/collectorpb/collector.pb.go
	$(call generate_fake,$@,$<,CollectorServiceClient)

golang/protobuf/collectorpb/collectorpbfakes/fake_collector_service_client.go: golang/protobuf/collectorpb/collector.pb.go
	$(call generate_fake,$@,$<,CollectorServiceClient)
