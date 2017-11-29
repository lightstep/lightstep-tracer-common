package test

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-common/test/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/test/lightsteppb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c lightsteppb.BinaryCarrier
	_ = c.ProtoSize()
}
