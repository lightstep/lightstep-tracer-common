default: build

genproto.mk:
	@docker pull lightstep/gogoprotoc:latest
	@-docker rm -v lightstep-get-genproto-mk
	@docker create --name lightstep-get-genproto-mk lightstep/gogoprotoc:latest
	@docker cp lightstep-get-genproto-mk:/root/genproto.mk genproto.mk
	@docker rm -v lightstep-get-genproto-mk

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

GOGO_LINKS  = $(call protoc_targets_to_link_targets,$(GOGO_GENTGTS))
PBUF_LINKS  = $(call protoc_targets_to_link_targets,$(PBUF_GENTGTS))

FAKES = \
	golang/gogo/collectorpb/collectorpbfakes/fake_collector_service_client.go \
	golang/protobuf/collectorpb/collectorpbfakes/fake_collector_service_client.go

.PHONY: default build test clean proto-links proto
.PHONY: $(GOGO_GENTGTS) $(PBUF_GENTGTS) $(GOGO_LINKS) $(PBUF_LINKS)

build: test

proto: $(GOGO_GENTGTS) $(PBUF_GENTGTS) $(FAKES) 

test: $(TEST_SOURCES)
	$(GOPATH)/bin/dep ensure && $(GOPATH)/bin/dep prune
	go test -v ./golang

clean: 
	$(call clean_protoc_targets,$(GOGO_GENTGTS) $(PBUF_GENTGTS))

proto-links: $(GOGO_LINKS) $(PBUF_LINKS)

$(GOGO_LINKS): $(GOLANG)-$(GOGO)-%-link: %.proto
	$(call gen_protoc_link,$<,$@,$(GOGO))

$(PBUF_LINKS): $(GOLANG)-$(PBUF)-%-link: %.proto
	$(call gen_protoc_link,$<,$@,$(PBUF))

$(GOGO_GENTGTS): $(GOLANG)-$(GOGO)-%: %.proto proto-links
	$(call gen_gogo_target,$<)

$(PBUF_GENTGTS): $(GOLANG)-$(PBUF)-%: %.proto proto-links
	$(call gen_protobuf_target,$<)

golang/gogo/collectorpb/collectorpbfakes/fake_collector_service_client.go: golang/gogo/collectorpb/collector.pb.go
	$(call generate_fake,$@,$<,CollectorServiceClient)

golang/protobuf/collectorpb/collectorpbfakes/fake_collector_service_client.go: golang/protobuf/collectorpb/collector.pb.go
	$(call generate_fake,$@,$<,CollectorServiceClient)
