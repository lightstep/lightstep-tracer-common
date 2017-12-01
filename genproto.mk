GOLANG = golang
PBUF   = protobuf
GOGO   = gogo

PWD    = $(shell pwd)
TMPDIR = $(PWD)/tmpgen

# List of standard protoc options
PROTOC_OPTS = plugins=grpc

# These flags manage mapping the google-standard protobuf types (e.g., Timestamp)
# into the annotated versions supplied with Gogo.  The trailing `,` matters.
GOGO_OPTS = Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,

define protos_to_gogo_targets
$(foreach proto,$(1),$(GOLANG)-$(GOGO)-$(basename $(proto)))
endef

define protos_to_protobuf_targets
$(foreach proto,$(1),$(GOLANG)-$(PBUF)-$(basename $(proto)))
endef

define gen_gogo_target
$(call gen_protoc_target,$(1),$(GOLANG)/$(GOGO)/$(basename $(1))pb/$(basename $(1)).pb.go,$(GOGO),--gogofaster_out=$(GOGO_OPTS)$(PROTOC_OPTS))
endef

define gen_protobuf_target
$(call gen_protoc_target,$(1),$(GOLANG)/$(PBUF)/$(basename $(1))pb/$(basename $(1)).pb.go,$(PBUF),--go_out=$(PROTOC_OPTS))
endef

# $(1) = .proto input
# $(2) = .pb.go output
# $(3) = gogo or protobuf
# $(4) = protoc-output spec
#
# Note: the --proto_path include "." below references the
# docker image's $(GOPATH)/src.
define gen_protoc_target
  @echo compiling $(1) [$(3)]
  @mkdir -p $(TMPDIR) 
  @docker run --rm \
    -v $(PWD):/input:ro \
    -v $(TMPDIR):/output \
    lightstep/gogoprotoc:latest \
    protoc \
    -I/input/third_party/googleapis \
    $(4):/output \
    --proto_path=/input:. \
    /input/$(1)
  @mkdir -p $(GOLANG)/$(3)/$(basename $(1))pb/$(basename $(1))pbfakes
  @sed 's@package $(basename $(1))pb@package $(basename $(1))pb // import "$(PKG_PREFIX)/golang/$(3)/$(basename $(1))pb"@' < $(TMPDIR)/$(basename $(1)).pb.go > $(GOLANG)/$(3)/$(basename $(1))pb/$(basename $(1)).pb.go
  @rm $(TMPDIR)/$(basename $(1)).pb.go
endef

define clean_protoc_targets
  @rm -rf $(foreach target,$(1),$(subst -,/,$(target)pb))
endef

# generate_fake: runs counterfeiter in docker container to generate fake classes
# $(1) output file path
# $(2) input file path
# $(3) class name
define generate_fake
  @docker run --rm \
	-v $(GOPATH):/usergo \
	lightstep/gobuild:latest \
	/bin/bash -c "cd /usergo/src/$(PKG_PREFIX) && counterfeiter -o $(1) $(2) $(3)"
endef
