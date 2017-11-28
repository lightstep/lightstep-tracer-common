package lightstep

import (
	"testing"

	"github.com/lightstep/lightstep-tracer-go/lightstep-tracer-common/collectorpb"
	"github.com/lightstep/lightstep-tracer-go/lightstep-tracer-common/lightsteppb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c lightsteppb.BinaryCarrier
	_ = c.ProtoSize()
}
