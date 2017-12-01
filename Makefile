# Note the protoc-generated code in this directory is a sanity check
# of the protos here, not generated code for import.
#
.PHONY: default build test

default: build

build: test

PWD = $(shell pwd)
TMPDIR = $(PWD)/tmpgen

GOLANG = golang
PBUF = protobuf
GOGO = gogo

PKG_PREFIX = github.com/lightstep/lightstep-tracer-common

TEST_SOURCES = \
	golang/protobuf_test.go \
	golang/gogo_test.go

PROTO_SOURCES = \
	collector.proto \
	carrier.proto

GOGO_GENRULES = $(foreach proto,$(PROTO_SOURCES),$(GOLANG)-$(GOGO)-$(basename $(proto)))
PBUF_GENRULES = $(foreach proto,$(PROTO_SOURCES),$(GOLANG)-$(PBUF)-$(basename $(proto)))

.PHONY: $(GOGO_GENRULES) $(PBUF_GENRULES)

$(GOGO_GENRULES): $(GOLANG)-$(GOGO)-%: %.proto
	@echo compiling $^ [gogoproto] ...
	@docker run --rm -v $(PWD):/input:ro -v $(TMPDIR):/output \
		lightstep/gogoprotoc:latest \
			protoc -I/input/third_party/googleapis \
				--gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:/output \
				--proto_path=/input:. \
				/input/$^
	@mkdir -p $(GOLANG)/$(GOGO)/$(basename $^)pb
	@sed 's@package $(basename $^)pb@package $(basename $^)pb // import "$(PKG_PREFIX)/golang/gogo/$(basename $^)pb"@' < $(TMPDIR)/$(basename $^).pb.go > $(GOLANG)/$(GOGO)/$(basename $^)pb/$(basename $^).pb.go
	@rm $(TMPDIR)/$(basename $^).pb.go

$(PBUF_GENRULES): $(GOLANG)-$(PBUF)-%: %.proto
	@echo compiling $^ [protobuf] ...
	@docker run --rm -v $(PWD):/input:ro -v $(TMPDIR):/output \
		lightstep/gogoprotoc:latest \
			protoc -I/input/third_party/googleapis \
				--go_out=plugins=grpc:/output \
				--proto_path=/input:. \
				/input/$^
	@mkdir -p $(GOLANG)/$(PBUF)/$(basename $^)pb
	@sed 's@package $(basename $^)pb@package $(basename $^)pb // import "$(PKG_PREFIX)/golang/protobuf/$(basename $^)pb"@' < $(TMPDIR)/$(basename $^).pb.go > $(GOLANG)/$(PBUF)/$(basename $^)pb/$(basename $^).pb.go
	@rm $(TMPDIR)/$(basename $^).pb.go

test: $(GOGO_GENRULES) $(PBUF_GENRULES) $(TEST_SOURCES)
	go test -v ./golang
