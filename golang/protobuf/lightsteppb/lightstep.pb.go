// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

/*
Package lightsteppb is a generated protocol buffer package.

It is generated from these files:
	github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

It has these top-level messages:
	BinaryCarrier
	BasicTracerCarrier
*/
package lightsteppb // import "github.com/lightstep/lightstep-tracer-common/golang/protobuf/lightsteppb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The standard carrier for binary context propagation into LightStep.
type BinaryCarrier struct {
	// "text_ctx" was deprecated following lightstep-tracer-cpp-0.36
	DeprecatedTextCtx [][]byte `protobuf:"bytes,1,rep,name=deprecated_text_ctx,json=deprecatedTextCtx,proto3" json:"deprecated_text_ctx,omitempty"`
	// The Opentracing "basictracer" proto.
	BasicCtx *BasicTracerCarrier `protobuf:"bytes,2,opt,name=basic_ctx,json=basicCtx" json:"basic_ctx,omitempty"`
}

func (m *BinaryCarrier) Reset()                    { *m = BinaryCarrier{} }
func (m *BinaryCarrier) String() string            { return proto.CompactTextString(m) }
func (*BinaryCarrier) ProtoMessage()               {}
func (*BinaryCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BinaryCarrier) GetDeprecatedTextCtx() [][]byte {
	if m != nil {
		return m.DeprecatedTextCtx
	}
	return nil
}

func (m *BinaryCarrier) GetBasicCtx() *BasicTracerCarrier {
	if m != nil {
		return m.BasicCtx
	}
	return nil
}

// Copy of https://github.com/opentracing/basictracer-go/blob/master/wire/wire.proto
type BasicTracerCarrier struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BasicTracerCarrier) Reset()                    { *m = BasicTracerCarrier{} }
func (m *BasicTracerCarrier) String() string            { return proto.CompactTextString(m) }
func (*BasicTracerCarrier) ProtoMessage()               {}
func (*BasicTracerCarrier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BasicTracerCarrier) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *BasicTracerCarrier) GetSampled() bool {
	if m != nil {
		return m.Sampled
	}
	return false
}

func (m *BasicTracerCarrier) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*BinaryCarrier)(nil), "lightstep.BinaryCarrier")
	proto.RegisterType((*BasicTracerCarrier)(nil), "lightstep.BasicTracerCarrier")
}

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x1d, 0xe3, 0x55, 0x33, 0x2a, 0x5c, 0xe7, 0x5e, 0x68, 0x2a, 0x98, 0x06, 0x57, 0xd9,
	0x98, 0x80, 0xdd, 0x94, 0x6e, 0x5a, 0x12, 0xba, 0x70, 0x1b, 0x5c, 0x75, 0x13, 0x26, 0xc9, 0x74,
	0x0c, 0x35, 0x7f, 0x98, 0x1c, 0x4b, 0x5c, 0xf6, 0x4d, 0xfa, 0x38, 0x2e, 0xfb, 0x04, 0xa5, 0xd8,
	0xa7, 0xe8, 0xae, 0x64, 0x6c, 0x8d, 0x20, 0x74, 0x77, 0x7e, 0xe7, 0xfb, 0xbe, 0xe4, 0x3b, 0x0c,
	0x76, 0x79, 0x0c, 0xcb, 0x75, 0x60, 0x85, 0x59, 0x62, 0xaf, 0x62, 0xbe, 0x84, 0x02, 0x58, 0x5e,
	0x4f, 0x53, 0x10, 0x34, 0x64, 0x62, 0x1a, 0x66, 0x49, 0x92, 0xa5, 0x36, 0x24, 0x39, 0x67, 0x69,
	0x2d, 0x5b, 0xb9, 0xc8, 0x20, 0x23, 0xea, 0x61, 0x31, 0x9a, 0x1e, 0x7d, 0x8f, 0x67, 0x3c, 0xb3,
	0xa5, 0x23, 0x58, 0x3f, 0x48, 0x92, 0x20, 0xa7, 0x7d, 0x72, 0xf2, 0x8c, 0xf0, 0xc0, 0x89, 0x53,
	0x2a, 0x36, 0x2e, 0x15, 0x22, 0x66, 0x82, 0x58, 0xf8, 0x5f, 0xc4, 0x72, 0xc1, 0x42, 0x0a, 0x2c,
	0xf2, 0x81, 0x95, 0xe0, 0x87, 0x50, 0x6a, 0xc8, 0x50, 0xcc, 0xbe, 0x37, 0xac, 0xa5, 0x05, 0x2b,
	0xc1, 0x85, 0x92, 0xdc, 0x62, 0x35, 0xa0, 0x45, 0x1c, 0x4a, 0x57, 0xd3, 0x40, 0x66, 0x6f, 0x36,
	0xb6, 0xea, 0x82, 0x4e, 0xa5, 0x2d, 0xe4, 0x11, 0xdf, 0x7f, 0x70, 0x5a, 0xdb, 0xb7, 0x8b, 0x86,
	0xd7, 0x95, 0x29, 0x17, 0xca, 0xc9, 0x27, 0xc2, 0xe4, 0xd4, 0x46, 0xce, 0x71, 0x57, 0x1e, 0xef,
	0xc7, 0x91, 0x86, 0x0c, 0x64, 0xb6, 0xbd, 0x8e, 0xe4, 0x79, 0x44, 0xce, 0x70, 0xa7, 0xc8, 0x69,
	0x5a, 0x29, 0x4d, 0xa9, 0xb4, 0x2b, 0x9c, 0x47, 0x44, 0xc3, 0x9d, 0x82, 0x26, 0xf9, 0x8a, 0x45,
	0x9a, 0x62, 0x20, 0xb3, 0xeb, 0xfd, 0x20, 0x59, 0xe0, 0x41, 0x40, 0x39, 0xa7, 0x9c, 0xf9, 0x31,
	0xb0, 0xa4, 0xd0, 0x5a, 0x86, 0x62, 0xf6, 0x66, 0xf6, 0xaf, 0x55, 0x2d, 0x67, 0x1f, 0x99, 0x57,
	0x89, 0xbb, 0x14, 0xc4, 0xc6, 0xeb, 0x07, 0x47, 0xab, 0xd1, 0x0d, 0x1e, 0x9e, 0x58, 0xc8, 0x5f,
	0xac, 0x3c, 0xb2, 0x8d, 0xec, 0xac, 0x7a, 0xd5, 0x48, 0xfe, 0xe3, 0x3f, 0x4f, 0x74, 0xb5, 0x66,
	0xb2, 0xad, 0xea, 0xed, 0xe1, 0xba, 0x79, 0x85, 0x9c, 0xf1, 0x76, 0xa7, 0xa3, 0xd7, 0x9d, 0x8e,
	0xde, 0x77, 0x7a, 0xe3, 0xe5, 0x43, 0x47, 0xf7, 0xbd, 0x43, 0xa1, 0x3c, 0x08, 0xda, 0xf2, 0x95,
	0x2e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xac, 0x69, 0x4d, 0x26, 0x02, 0x00, 0x00,
}