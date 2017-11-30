package golang

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/lightsteppb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c lightsteppb.BinaryCarrier
	_ = c.ProtoSize()
}
