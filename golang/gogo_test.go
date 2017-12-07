package golang

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/collectorpb/collectorpbfakes"
	"github.com/lightstep/lightstep-tracer-common/golang/gogo/lightsteppb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_, _ = r.Marshal()

	var c lightsteppb.BinaryCarrier
	_, _ = c.Marshal()

	_ = collectorpbfakes.FakeCollectorServiceClient{}
}
