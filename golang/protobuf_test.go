package golang

import (
	"reflect"
	"testing"

	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/carrierpb"
	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb"
	"github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb/collectorpbfakes"
)

func TestProtoIsProtobuf(t *testing.T) {
	var r collectorpb.ReportRequest
	if ptag := reflect.ValueOf(r).Type().Field(0).Tag.Get("protobuf"); ptag == "" {
		panic("Not a protobuf!")
	}

	var c carrierpb.BinaryCarrier
	if ptag := reflect.ValueOf(c).Type().Field(0).Tag.Get("protobuf"); ptag == "" {
		panic("Not a protobuf!")
	}

	_ = collectorpbfakes.FakeCollectorServiceClient{}
}
