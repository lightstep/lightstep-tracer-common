package golang

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/gogo/carrierpb"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb/collectorpbfakes"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c carrierpb.BinaryCarrier
	_ = c.ProtoSize()

	_ = collectorpbfakes.FakeCollectorServiceClient{}
}
