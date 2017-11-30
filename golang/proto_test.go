package golang

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/gogo/carrierpb"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c carrierpb.BinaryCarrier
	_ = c.ProtoSize()
}
