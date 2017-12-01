// Code generated by protoc-gen-go. DO NOT EDIT.
// source: collector.proto

/*
Package collectorpb is a generated protocol buffer package.

It is generated from these files:
	collector.proto

It has these top-level messages:
	SpanContext
	KeyValue
	Log
	Reference
	Span
	Reporter
	MetricsSample
	InternalMetrics
	Auth
	ReportRequest
	Command
	ReportResponse
*/
package collectorpb // import "github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Reference_Relationship int32

const (
	Reference_CHILD_OF     Reference_Relationship = 0
	Reference_FOLLOWS_FROM Reference_Relationship = 1
)

var Reference_Relationship_name = map[int32]string{
	0: "CHILD_OF",
	1: "FOLLOWS_FROM",
}
var Reference_Relationship_value = map[string]int32{
	"CHILD_OF":     0,
	"FOLLOWS_FROM": 1,
}

func (x Reference_Relationship) String() string {
	return proto.EnumName(Reference_Relationship_name, int32(x))
}
func (Reference_Relationship) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type SpanContext struct {
	TraceId uint64            `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId  uint64            `protobuf:"varint,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Baggage map[string]string `protobuf:"bytes,3,rep,name=baggage" json:"baggage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SpanContext) Reset()                    { *m = SpanContext{} }
func (m *SpanContext) String() string            { return proto.CompactTextString(m) }
func (*SpanContext) ProtoMessage()               {}
func (*SpanContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpanContext) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *SpanContext) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanContext) GetBaggage() map[string]string {
	if m != nil {
		return m.Baggage
	}
	return nil
}

// Represent both tags and log fields.
type KeyValue struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*KeyValue_StringValue
	//	*KeyValue_IntValue
	//	*KeyValue_DoubleValue
	//	*KeyValue_BoolValue
	//	*KeyValue_JsonValue
	Value isKeyValue_Value `protobuf_oneof:"value"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isKeyValue_Value interface {
	isKeyValue_Value()
}

type KeyValue_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,oneof"`
}
type KeyValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,oneof"`
}
type KeyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,oneof"`
}
type KeyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,oneof"`
}
type KeyValue_JsonValue struct {
	JsonValue string `protobuf:"bytes,6,opt,name=json_value,json=jsonValue,oneof"`
}

func (*KeyValue_StringValue) isKeyValue_Value() {}
func (*KeyValue_IntValue) isKeyValue_Value()    {}
func (*KeyValue_DoubleValue) isKeyValue_Value() {}
func (*KeyValue_BoolValue) isKeyValue_Value()   {}
func (*KeyValue_JsonValue) isKeyValue_Value()   {}

func (m *KeyValue) GetValue() isKeyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetStringValue() string {
	if x, ok := m.GetValue().(*KeyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *KeyValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*KeyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *KeyValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*KeyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *KeyValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*KeyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *KeyValue) GetJsonValue() string {
	if x, ok := m.GetValue().(*KeyValue_JsonValue); ok {
		return x.JsonValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeyValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeyValue_OneofMarshaler, _KeyValue_OneofUnmarshaler, _KeyValue_OneofSizer, []interface{}{
		(*KeyValue_StringValue)(nil),
		(*KeyValue_IntValue)(nil),
		(*KeyValue_DoubleValue)(nil),
		(*KeyValue_BoolValue)(nil),
		(*KeyValue_JsonValue)(nil),
	}
}

func _KeyValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *KeyValue_IntValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *KeyValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *KeyValue_JsonValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonValue)
	case nil:
	default:
		return fmt.Errorf("KeyValue.Value has unexpected type %T", x)
	}
	return nil
}

func _KeyValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeyValue)
	switch tag {
	case 2: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &KeyValue_StringValue{x}
		return true, err
	case 3: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_IntValue{int64(x)}
		return true, err
	case 4: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &KeyValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 5: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_BoolValue{x != 0}
		return true, err
	case 6: // value.json_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &KeyValue_JsonValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _KeyValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *KeyValue_IntValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		n += proto.SizeVarint(4<<3 | proto.WireFixed64)
		n += 8
	case *KeyValue_BoolValue:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *KeyValue_JsonValue:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.JsonValue)))
		n += len(x.JsonValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Fields    []*KeyValue                 `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Log) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Log) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Reference struct {
	Relationship Reference_Relationship `protobuf:"varint,1,opt,name=relationship,enum=lightstep.collector.Reference_Relationship" json:"relationship,omitempty"`
	SpanContext  *SpanContext           `protobuf:"bytes,2,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
}

func (m *Reference) Reset()                    { *m = Reference{} }
func (m *Reference) String() string            { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()               {}
func (*Reference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Reference) GetRelationship() Reference_Relationship {
	if m != nil {
		return m.Relationship
	}
	return Reference_CHILD_OF
}

func (m *Reference) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

type Span struct {
	SpanContext    *SpanContext                `protobuf:"bytes,1,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
	OperationName  string                      `protobuf:"bytes,2,opt,name=operation_name,json=operationName" json:"operation_name,omitempty"`
	References     []*Reference                `protobuf:"bytes,3,rep,name=references" json:"references,omitempty"`
	StartTimestamp *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                      `protobuf:"varint,5,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Tags           []*KeyValue                 `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Logs           []*Log                      `protobuf:"bytes,7,rep,name=logs" json:"logs,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Span) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

func (m *Span) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Span) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *Span) GetStartTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *Span) GetDurationMicros() uint64 {
	if m != nil {
		return m.DurationMicros
	}
	return 0
}

func (m *Span) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Reporter struct {
	ReporterId uint64      `protobuf:"varint,1,opt,name=reporter_id,json=reporterId" json:"reporter_id,omitempty"`
	Tags       []*KeyValue `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Reporter) Reset()                    { *m = Reporter{} }
func (m *Reporter) String() string            { return proto.CompactTextString(m) }
func (*Reporter) ProtoMessage()               {}
func (*Reporter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Reporter) GetReporterId() uint64 {
	if m != nil {
		return m.ReporterId
	}
	return 0
}

func (m *Reporter) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type MetricsSample struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*MetricsSample_IntValue
	//	*MetricsSample_DoubleValue
	Value isMetricsSample_Value `protobuf_oneof:"value"`
}

func (m *MetricsSample) Reset()                    { *m = MetricsSample{} }
func (m *MetricsSample) String() string            { return proto.CompactTextString(m) }
func (*MetricsSample) ProtoMessage()               {}
func (*MetricsSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isMetricsSample_Value interface {
	isMetricsSample_Value()
}

type MetricsSample_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,oneof"`
}
type MetricsSample_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,oneof"`
}

func (*MetricsSample_IntValue) isMetricsSample_Value()    {}
func (*MetricsSample_DoubleValue) isMetricsSample_Value() {}

func (m *MetricsSample) GetValue() isMetricsSample_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricsSample) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsSample) GetIntValue() int64 {
	if x, ok := m.GetValue().(*MetricsSample_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *MetricsSample) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*MetricsSample_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricsSample) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricsSample_OneofMarshaler, _MetricsSample_OneofUnmarshaler, _MetricsSample_OneofSizer, []interface{}{
		(*MetricsSample_IntValue)(nil),
		(*MetricsSample_DoubleValue)(nil),
	}
}

func _MetricsSample_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case nil:
	default:
		return fmt.Errorf("MetricsSample.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricsSample_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricsSample)
	switch tag {
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricsSample_IntValue{int64(x)}
		return true, err
	case 3: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricsSample_DoubleValue{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _MetricsSample_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InternalMetrics struct {
	StartTimestamp *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                      `protobuf:"varint,2,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Logs           []*Log                      `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
	Counts         []*MetricsSample            `protobuf:"bytes,4,rep,name=counts" json:"counts,omitempty"`
	Gauges         []*MetricsSample            `protobuf:"bytes,5,rep,name=gauges" json:"gauges,omitempty"`
}

func (m *InternalMetrics) Reset()                    { *m = InternalMetrics{} }
func (m *InternalMetrics) String() string            { return proto.CompactTextString(m) }
func (*InternalMetrics) ProtoMessage()               {}
func (*InternalMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *InternalMetrics) GetStartTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *InternalMetrics) GetDurationMicros() uint64 {
	if m != nil {
		return m.DurationMicros
	}
	return 0
}

func (m *InternalMetrics) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *InternalMetrics) GetCounts() []*MetricsSample {
	if m != nil {
		return m.Counts
	}
	return nil
}

func (m *InternalMetrics) GetGauges() []*MetricsSample {
	if m != nil {
		return m.Gauges
	}
	return nil
}

type Auth struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Auth) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ReportRequest struct {
	Reporter              *Reporter        `protobuf:"bytes,1,opt,name=reporter" json:"reporter,omitempty"`
	Auth                  *Auth            `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	Spans                 []*Span          `protobuf:"bytes,3,rep,name=spans" json:"spans,omitempty"`
	TimestampOffsetMicros int32            `protobuf:"varint,5,opt,name=timestamp_offset_micros,json=timestampOffsetMicros" json:"timestamp_offset_micros,omitempty"`
	InternalMetrics       *InternalMetrics `protobuf:"bytes,6,opt,name=internal_metrics,json=internalMetrics" json:"internal_metrics,omitempty"`
}

func (m *ReportRequest) Reset()                    { *m = ReportRequest{} }
func (m *ReportRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()               {}
func (*ReportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReportRequest) GetReporter() *Reporter {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func (m *ReportRequest) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ReportRequest) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ReportRequest) GetTimestampOffsetMicros() int32 {
	if m != nil {
		return m.TimestampOffsetMicros
	}
	return 0
}

func (m *ReportRequest) GetInternalMetrics() *InternalMetrics {
	if m != nil {
		return m.InternalMetrics
	}
	return nil
}

type Command struct {
	Disable bool `protobuf:"varint,1,opt,name=disable" json:"disable,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Command) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type ReportResponse struct {
	Commands          []*Command                  `protobuf:"bytes,1,rep,name=commands" json:"commands,omitempty"`
	ReceiveTimestamp  *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=receive_timestamp,json=receiveTimestamp" json:"receive_timestamp,omitempty"`
	TransmitTimestamp *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=transmit_timestamp,json=transmitTimestamp" json:"transmit_timestamp,omitempty"`
	Errors            []string                    `protobuf:"bytes,4,rep,name=errors" json:"errors,omitempty"`
	Warnings          []string                    `protobuf:"bytes,5,rep,name=warnings" json:"warnings,omitempty"`
	Infos             []string                    `protobuf:"bytes,6,rep,name=infos" json:"infos,omitempty"`
}

func (m *ReportResponse) Reset()                    { *m = ReportResponse{} }
func (m *ReportResponse) String() string            { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()               {}
func (*ReportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReportResponse) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *ReportResponse) GetReceiveTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *ReportResponse) GetTransmitTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.TransmitTimestamp
	}
	return nil
}

func (m *ReportResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ReportResponse) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func (m *ReportResponse) GetInfos() []string {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterType((*SpanContext)(nil), "lightstep.collector.SpanContext")
	proto.RegisterType((*KeyValue)(nil), "lightstep.collector.KeyValue")
	proto.RegisterType((*Log)(nil), "lightstep.collector.Log")
	proto.RegisterType((*Reference)(nil), "lightstep.collector.Reference")
	proto.RegisterType((*Span)(nil), "lightstep.collector.Span")
	proto.RegisterType((*Reporter)(nil), "lightstep.collector.Reporter")
	proto.RegisterType((*MetricsSample)(nil), "lightstep.collector.MetricsSample")
	proto.RegisterType((*InternalMetrics)(nil), "lightstep.collector.InternalMetrics")
	proto.RegisterType((*Auth)(nil), "lightstep.collector.Auth")
	proto.RegisterType((*ReportRequest)(nil), "lightstep.collector.ReportRequest")
	proto.RegisterType((*Command)(nil), "lightstep.collector.Command")
	proto.RegisterType((*ReportResponse)(nil), "lightstep.collector.ReportResponse")
	proto.RegisterEnum("lightstep.collector.Reference_Relationship", Reference_Relationship_name, Reference_Relationship_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CollectorService service

type CollectorServiceClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type collectorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCollectorServiceClient(cc *grpc.ClientConn) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := grpc.Invoke(ctx, "/lightstep.collector.CollectorService/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CollectorService service

type CollectorServiceServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterCollectorServiceServer(s *grpc.Server, srv CollectorServiceServer) {
	s.RegisterService(&_CollectorService_serviceDesc, srv)
}

func _CollectorService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightstep.collector.CollectorService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lightstep.collector.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _CollectorService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collector.proto",
}

func init() { proto.RegisterFile("collector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0x1b, 0xff, 0x78, 0x76, 0x63, 0x77, 0xf8, 0x51, 0xc7, 0x6a, 0x13, 0xb3, 0x01,
	0x11, 0x40, 0xd9, 0x48, 0xae, 0x10, 0x28, 0x20, 0xa0, 0x4e, 0x09, 0xb1, 0x70, 0x70, 0x59, 0x17,
	0x90, 0x7a, 0xb1, 0xd6, 0xeb, 0xf1, 0x66, 0xe9, 0xee, 0xcc, 0x32, 0x33, 0x0e, 0xe4, 0x86, 0x38,
	0x70, 0x46, 0x42, 0x48, 0x1c, 0x11, 0x37, 0xfe, 0x0d, 0x0e, 0xd0, 0x63, 0x11, 0x07, 0x6e, 0x80,
	0x02, 0x7f, 0x08, 0xda, 0x99, 0xd9, 0xb5, 0x53, 0xdc, 0xb4, 0x41, 0xdc, 0x66, 0xde, 0xfb, 0xbe,
	0x37, 0xcf, 0xef, 0x7d, 0xef, 0x79, 0xa1, 0xee, 0xd1, 0x30, 0xc4, 0x9e, 0xa0, 0xcc, 0x8e, 0x19,
	0x15, 0x14, 0x3d, 0x11, 0x06, 0xfe, 0x91, 0xe0, 0x02, 0xc7, 0x76, 0xe6, 0x6a, 0x6d, 0xfb, 0x81,
	0x38, 0x9a, 0x8d, 0x6d, 0x8f, 0x46, 0x3b, 0x3e, 0xf5, 0xe9, 0x8e, 0xc4, 0x8e, 0x67, 0x53, 0x79,
	0x93, 0x17, 0x79, 0x52, 0x31, 0x5a, 0x1b, 0x3e, 0xa5, 0x7e, 0x88, 0xe7, 0x28, 0x11, 0x44, 0x98,
	0x0b, 0x37, 0x8a, 0x35, 0xe0, 0xaa, 0x06, 0xb8, 0x71, 0xb0, 0xe3, 0x12, 0x42, 0x85, 0x2b, 0x02,
	0x4a, 0xb8, 0xf2, 0x5a, 0x3f, 0x19, 0x50, 0x1d, 0xc6, 0x2e, 0xd9, 0xa3, 0x44, 0xe0, 0xcf, 0x04,
	0x5a, 0x83, 0xb2, 0x60, 0xae, 0x87, 0x47, 0xc1, 0xa4, 0x69, 0xb4, 0x8d, 0x2d, 0xd3, 0x29, 0xc9,
	0x7b, 0x6f, 0x82, 0xae, 0x40, 0x89, 0xc7, 0x2e, 0x49, 0x3c, 0x79, 0xe9, 0x29, 0x26, 0xd7, 0xde,
	0x04, 0xbd, 0x03, 0xa5, 0xb1, 0xeb, 0xfb, 0xae, 0x8f, 0x9b, 0x85, 0x76, 0x61, 0xab, 0xda, 0xd9,
	0xb6, 0x97, 0xfc, 0x30, 0x7b, 0xe1, 0x19, 0xbb, 0xab, 0xf0, 0x6f, 0x13, 0xc1, 0x4e, 0x9c, 0x94,
	0xdd, 0xda, 0x85, 0xda, 0xa2, 0x03, 0x35, 0xa0, 0x70, 0x17, 0x9f, 0xc8, 0x3c, 0x2a, 0x4e, 0x72,
	0x44, 0x4f, 0xc2, 0xca, 0xb1, 0x1b, 0xce, 0xb0, 0xcc, 0xa0, 0xe2, 0xa8, 0xcb, 0x6e, 0xfe, 0x55,
	0xc3, 0xba, 0x6f, 0x40, 0xf9, 0x5d, 0x7c, 0xf2, 0x61, 0x62, 0x58, 0x42, 0xdc, 0x84, 0x1a, 0x17,
	0x2c, 0x20, 0xfe, 0x68, 0x81, 0x7f, 0x90, 0x73, 0xaa, 0xca, 0xaa, 0x68, 0xd7, 0xa0, 0x12, 0x10,
	0xa1, 0x11, 0x85, 0xb6, 0xb1, 0x55, 0x38, 0xc8, 0x39, 0xe5, 0x80, 0x08, 0xe5, 0xde, 0x84, 0xda,
	0x84, 0xce, 0xc6, 0x21, 0xd6, 0x08, 0xb3, 0x6d, 0x6c, 0x19, 0x49, 0x0c, 0x65, 0x55, 0xa0, 0x0d,
	0x80, 0x31, 0xa5, 0xa1, 0x86, 0xac, 0xb4, 0x8d, 0xad, 0xf2, 0x41, 0xce, 0xa9, 0x24, 0xb6, 0x0c,
	0xf0, 0x31, 0xa7, 0x44, 0x03, 0x8a, 0x3a, 0x8f, 0x4a, 0x62, 0x93, 0x80, 0x6e, 0x49, 0xff, 0x46,
	0xeb, 0x4b, 0x03, 0x0a, 0x7d, 0xea, 0xa3, 0x2e, 0x54, 0xb2, 0xa6, 0xca, 0xdf, 0x54, 0xed, 0xb4,
	0x6c, 0xd5, 0x55, 0x3b, 0x6d, 0xbb, 0x7d, 0x3b, 0x45, 0x74, 0xcb, 0xf7, 0x7e, 0xdf, 0xc8, 0x7d,
	0xf5, 0xc7, 0x86, 0xe1, 0xcc, 0x69, 0xe8, 0x35, 0x28, 0x4e, 0x03, 0x1c, 0x4e, 0x78, 0x33, 0x2f,
	0x5b, 0x74, 0x6d, 0x69, 0x8b, 0xd2, 0x02, 0x76, 0xcd, 0x24, 0x86, 0xa3, 0x29, 0xd6, 0x6f, 0x06,
	0x54, 0x1c, 0x3c, 0xc5, 0x0c, 0x13, 0x0f, 0xa3, 0x01, 0xd4, 0x18, 0x0e, 0x95, 0x8a, 0x8e, 0x02,
	0x95, 0xd1, 0x6a, 0xe7, 0xa5, 0xa5, 0x01, 0x33, 0x96, 0xed, 0x2c, 0x50, 0x9c, 0x33, 0x01, 0x50,
	0x0f, 0x6a, 0x52, 0x58, 0x9e, 0x12, 0x87, 0xec, 0x4d, 0xb5, 0xd3, 0x7e, 0x94, 0x88, 0x74, 0x92,
	0x55, 0x3e, 0x37, 0x59, 0x36, 0xd4, 0x16, 0x1f, 0x42, 0x35, 0x28, 0xef, 0x1d, 0xf4, 0xfa, 0x37,
	0x47, 0x83, 0xfd, 0x46, 0x0e, 0x35, 0xa0, 0xb6, 0x3f, 0xe8, 0xf7, 0x07, 0x1f, 0x0d, 0x47, 0xfb,
	0xce, 0xe0, 0xb0, 0x61, 0x58, 0x3f, 0x14, 0xc0, 0x4c, 0x42, 0xfe, 0x2b, 0x07, 0xe3, 0x3f, 0xe7,
	0x80, 0x9e, 0x83, 0x55, 0x1a, 0x63, 0x26, 0x93, 0x18, 0x11, 0x37, 0x4a, 0xc5, 0x7a, 0x29, 0xb3,
	0xbe, 0xe7, 0x46, 0x18, 0xdd, 0x04, 0x60, 0x69, 0x75, 0xb8, 0x1e, 0x9c, 0xf5, 0xf3, 0x8b, 0xa8,
	0x5f, 0x5b, 0xe0, 0xa1, 0x43, 0xa8, 0x73, 0xe1, 0x32, 0x31, 0x9a, 0x2b, 0xc4, 0xbc, 0x80, 0x42,
	0x56, 0x25, 0x39, 0xf3, 0xa0, 0x6d, 0xa8, 0x4f, 0x66, 0x3a, 0xf5, 0x28, 0xf0, 0x18, 0xe5, 0x52,
	0xc2, 0x66, 0xd7, 0xfc, 0x56, 0xc2, 0x53, 0xe7, 0xa1, 0xf4, 0xa1, 0x57, 0xc0, 0x14, 0xae, 0xcf,
	0x9b, 0xc5, 0xc7, 0xd7, 0x94, 0x24, 0xa0, 0x0e, 0x98, 0x21, 0xf5, 0x79, 0xb3, 0x24, 0x89, 0xcd,
	0xa5, 0xc4, 0x3e, 0xf5, 0x53, 0x4e, 0x82, 0xb5, 0x26, 0x50, 0x76, 0x70, 0x4c, 0x99, 0xc0, 0x0c,
	0x6d, 0x40, 0x95, 0xe9, 0xf3, 0x7c, 0x53, 0x41, 0x6a, 0xea, 0x4d, 0xb2, 0xcc, 0xcc, 0x0b, 0x66,
	0x66, 0xc5, 0x70, 0xe9, 0x10, 0x0b, 0x16, 0x78, 0x7c, 0xe8, 0x46, 0x71, 0x88, 0x11, 0x02, 0x53,
	0x36, 0x51, 0x2d, 0x13, 0x79, 0x3e, 0xbb, 0x28, 0xf2, 0x8f, 0x5c, 0x14, 0x85, 0x25, 0x8b, 0x62,
	0x3e, 0xe6, 0x3f, 0xe7, 0xa1, 0xde, 0x23, 0x02, 0x33, 0xe2, 0x86, 0xfa, 0xe9, 0x65, 0x6d, 0x35,
	0xfe, 0xdf, 0xb6, 0xe6, 0xcf, 0x69, 0x6b, 0xda, 0x9d, 0xc2, 0xe3, 0x77, 0x07, 0xbd, 0x05, 0x45,
	0x8f, 0xce, 0x88, 0x48, 0x4b, 0x6e, 0x2d, 0x65, 0x9d, 0x29, 0x6d, 0xba, 0x65, 0x14, 0x2f, 0x89,
	0xe0, 0xbb, 0x33, 0x1f, 0x27, 0x92, 0xbb, 0x60, 0x04, 0xc5, 0xb3, 0x5e, 0x00, 0xf3, 0xc6, 0x4c,
	0x1c, 0xa1, 0x67, 0xa0, 0xe6, 0x7a, 0x1e, 0xe6, 0x7c, 0x24, 0xe8, 0x5d, 0x4c, 0x74, 0xeb, 0xaa,
	0xca, 0x76, 0x3b, 0x31, 0x59, 0xbf, 0xe4, 0xe1, 0x92, 0x52, 0x93, 0x83, 0x3f, 0x99, 0x61, 0x2e,
	0xd0, 0x9b, 0x50, 0x4e, 0xf5, 0xa3, 0x6b, 0x7d, 0xed, 0x21, 0xd3, 0xa8, 0x40, 0xfa, 0xed, 0x8c,
	0x84, 0xae, 0x83, 0xe9, 0xce, 0xc4, 0x91, 0x5e, 0x5f, 0x6b, 0x4b, 0xc9, 0x49, 0x7a, 0x69, 0xd9,
	0x12, 0x30, 0x7a, 0x19, 0x56, 0x92, 0xdd, 0x91, 0xd6, 0x7a, 0xed, 0xa1, 0x0b, 0x47, 0xb3, 0x14,
	0x1a, 0xbd, 0x0e, 0x57, 0x32, 0x65, 0x8c, 0xe8, 0x74, 0xca, 0xb1, 0x58, 0x9c, 0xd7, 0x15, 0xdd,
	0xd8, 0xa7, 0x32, 0xd0, 0x40, 0x62, 0x74, 0x7f, 0x3f, 0x80, 0x46, 0xa0, 0x05, 0x37, 0x8a, 0x54,
	0x3d, 0xe5, 0x1f, 0x51, 0xb5, 0xf3, 0xec, 0xd2, 0xf7, 0x1f, 0x50, 0xa7, 0x4e, 0xa5, 0x1e, 0x9c,
	0x35, 0x5b, 0x9b, 0x50, 0xda, 0xa3, 0x51, 0xe4, 0x92, 0x09, 0x6a, 0x42, 0x69, 0x12, 0x70, 0x77,
	0x1c, 0xaa, 0xb9, 0x29, 0x3b, 0xe9, 0xd5, 0xfa, 0x31, 0x0f, 0xab, 0x69, 0xe1, 0x79, 0x4c, 0x09,
	0xc7, 0xe8, 0x0d, 0x28, 0x7b, 0x8a, 0xc7, 0x9b, 0x86, 0x2c, 0xc3, 0xd5, 0xa5, 0x69, 0xe8, 0xe0,
	0x69, 0xe1, 0x53, 0x0e, 0x7a, 0x1f, 0x2e, 0x33, 0xec, 0xe1, 0xe0, 0x18, 0x2f, 0x8c, 0x4b, 0xfe,
	0x02, 0xe3, 0xd2, 0xd0, 0xf4, 0xf9, 0xc0, 0x0c, 0x01, 0x09, 0xe6, 0x12, 0x1e, 0x05, 0x8b, 0x23,
	0x58, 0xb8, 0x40, 0xcc, 0xcb, 0x29, 0x7f, 0x1e, 0xf4, 0x69, 0x28, 0x62, 0xc6, 0x28, 0x53, 0x23,
	0x52, 0x71, 0xf4, 0x0d, 0xb5, 0xa0, 0xfc, 0xa9, 0xcb, 0x48, 0x40, 0x7c, 0x25, 0xfd, 0x8a, 0x93,
	0xdd, 0x93, 0x0f, 0x9e, 0x80, 0x4c, 0xa9, 0x5a, 0xb1, 0x15, 0x47, 0x5d, 0x3a, 0xdf, 0x18, 0xd0,
	0xd8, 0x4b, 0xeb, 0x32, 0xc4, 0xec, 0x38, 0xf0, 0x30, 0xfa, 0xdc, 0x80, 0xa2, 0xaa, 0x2c, 0xb2,
	0xce, 0x51, 0xae, 0xd6, 0x7b, 0x6b, 0xf3, 0x5c, 0x8c, 0x6a, 0x8d, 0xb5, 0xfd, 0xc5, 0xaf, 0x7f,
	0x7f, 0x9d, 0x7f, 0xde, 0xaa, 0xcb, 0xcf, 0xc7, 0xe3, 0xce, 0x8e, 0x52, 0x3b, 0xdf, 0x35, 0x5e,
	0xbc, 0x73, 0x19, 0x3d, 0x68, 0xed, 0xde, 0xb8, 0x77, 0xba, 0x6e, 0xdc, 0x3f, 0x5d, 0x37, 0xfe,
	0x3c, 0x5d, 0xcf, 0x7d, 0xf7, 0xd7, 0xba, 0x01, 0x6b, 0x1e, 0x8d, 0x16, 0x1e, 0x92, 0xdf, 0x92,
	0xcc, 0xf6, 0x59, 0xec, 0xdd, 0x32, 0xee, 0x54, 0xb3, 0x77, 0xe3, 0xf1, 0xf7, 0x79, 0xb3, 0x3f,
	0xbc, 0xd5, 0x1d, 0x17, 0x65, 0x55, 0xaf, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x5f, 0x52,
	0x04, 0x2d, 0x0b, 0x00, 0x00,
}
