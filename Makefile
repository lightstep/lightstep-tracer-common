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

PROTO_SOURCES = collector.proto carrier.proto

GOGO_GENRULES = $(foreach proto,$(PROTO_SOURCES),$(GOLANG)-$(GOGO)-$(basename $(proto)))
.PHONY: $(GOGO_GENRULES)

$(GOGO_GENRULES): $(GOLANG)-$(GOGO)-%: %.proto
	@echo compiling $^ [gogo] ...
	@docker run --rm -v $(PWD):/input:ro -v $(TMPDIR):/output \
		lightstep/gogoprotoc:latest \
			protoc -I/input/third_party/googleapis \
				--gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:/output \
				--proto_path=/input:. \
				/input/$^
	@mkdir -p $(GOLANG)/$(GOGO)/$(basename $^)pb
	@mv $(TMPDIR)/$(basename $^).pb.go $(GOLANG)/$(GOGO)/$(basename $^)pb

test: $(GOGO_GENRULES) golang/proto_test.go
	go test -v ./golang
