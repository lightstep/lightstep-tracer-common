package test

import (
	"testing"

	"./collectorpb"
	"./lightsteppb"
)

func TestProtoIsGogo(t *testing.T) {
	var r collectorpb.ReportRequest
	_ = r.ProtoSize()

	var c lightsteppb.BinaryCarrier
	_ = c.ProtoSize()
}
