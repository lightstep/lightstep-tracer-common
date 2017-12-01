// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

/*
	Package lightsteppb is a generated protocol buffer package.

	It is generated from these files:
		github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto

	It has these top-level messages:
		BinaryCarrier
		BasicTracerCarrier
*/
package lightsteppb // import "github.com/lightstep/lightstep-tracer-common/golang/gogo/lightsteppb"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The standard carrier for binary context propagation into LightStep.
type BinaryCarrier struct {
	// "text_ctx" was deprecated following lightstep-tracer-cpp-0.36
	DeprecatedTextCtx [][]byte `protobuf:"bytes,1,rep,name=deprecated_text_ctx,json=deprecatedTextCtx" json:"deprecated_text_ctx,omitempty"`
	// The Opentracing "basictracer" proto.
	BasicCtx BasicTracerCarrier `protobuf:"bytes,2,opt,name=basic_ctx,json=basicCtx" json:"basic_ctx"`
}

func (m *BinaryCarrier) Reset()                    { *m = BinaryCarrier{} }
func (m *BinaryCarrier) String() string            { return proto.CompactTextString(m) }
func (*BinaryCarrier) ProtoMessage()               {}
func (*BinaryCarrier) Descriptor() ([]byte, []int) { return fileDescriptorLightstep, []int{0} }

func (m *BinaryCarrier) GetDeprecatedTextCtx() [][]byte {
	if m != nil {
		return m.DeprecatedTextCtx
	}
	return nil
}

func (m *BinaryCarrier) GetBasicCtx() BasicTracerCarrier {
	if m != nil {
		return m.BasicCtx
	}
	return BasicTracerCarrier{}
}

// Copy of https://github.com/opentracing/basictracer-go/blob/master/wire/wire.proto
type BasicTracerCarrier struct {
	TraceId      uint64            `protobuf:"fixed64,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId       uint64            `protobuf:"fixed64,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled,proto3" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items,json=baggageItems" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *BasicTracerCarrier) Reset()                    { *m = BasicTracerCarrier{} }
func (m *BasicTracerCarrier) String() string            { return proto.CompactTextString(m) }
func (*BasicTracerCarrier) ProtoMessage()               {}
func (*BasicTracerCarrier) Descriptor() ([]byte, []int) { return fileDescriptorLightstep, []int{1} }

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
func (m *BinaryCarrier) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BinaryCarrier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeprecatedTextCtx) > 0 {
		for _, b := range m.DeprecatedTextCtx {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintLightstep(dAtA, i, uint64(m.BasicCtx.ProtoSize()))
	n1, err := m.BasicCtx.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *BasicTracerCarrier) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicTracerCarrier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TraceId != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.TraceId))
		i += 8
	}
	if m.SpanId != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.SpanId))
		i += 8
	}
	if m.Sampled {
		dAtA[i] = 0x18
		i++
		if m.Sampled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.BaggageItems) > 0 {
		for k, _ := range m.BaggageItems {
			dAtA[i] = 0x22
			i++
			v := m.BaggageItems[k]
			mapSize := 1 + len(k) + sovLightstep(uint64(len(k))) + 1 + len(v) + sovLightstep(uint64(len(v)))
			i = encodeVarintLightstep(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintLightstep(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintLightstep(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BinaryCarrier) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.DeprecatedTextCtx) > 0 {
		for _, b := range m.DeprecatedTextCtx {
			l = len(b)
			n += 1 + l + sovLightstep(uint64(l))
		}
	}
	l = m.BasicCtx.ProtoSize()
	n += 1 + l + sovLightstep(uint64(l))
	return n
}

func (m *BasicTracerCarrier) ProtoSize() (n int) {
	var l int
	_ = l
	if m.TraceId != 0 {
		n += 9
	}
	if m.SpanId != 0 {
		n += 9
	}
	if m.Sampled {
		n += 2
	}
	if len(m.BaggageItems) > 0 {
		for k, v := range m.BaggageItems {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLightstep(uint64(len(k))) + 1 + len(v) + sovLightstep(uint64(len(v)))
			n += mapEntrySize + 1 + sovLightstep(uint64(mapEntrySize))
		}
	}
	return n
}

func sovLightstep(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLightstep(x uint64) (n int) {
	return sovLightstep(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BinaryCarrier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLightstep
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BinaryCarrier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BinaryCarrier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeprecatedTextCtx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeprecatedTextCtx = append(m.DeprecatedTextCtx, make([]byte, postIndex-iNdEx))
			copy(m.DeprecatedTextCtx[len(m.DeprecatedTextCtx)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasicCtx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BasicCtx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLightstep(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLightstep
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BasicTracerCarrier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLightstep
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BasicTracerCarrier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicTracerCarrier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.SpanId = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sampled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sampled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaggageItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLightstep
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaggageItems == nil {
				m.BaggageItems = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLightstep
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLightstep
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLightstep
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLightstep
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLightstep
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLightstep(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthLightstep
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.BaggageItems[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLightstep(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLightstep
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLightstep(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLightstep
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLightstep
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLightstep
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLightstep
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLightstep(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLightstep = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLightstep   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/tmpgen/lightstep.proto", fileDescriptorLightstep)
}

var fileDescriptorLightstep = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbd, 0x6e, 0xe2, 0x40,
	0x14, 0x85, 0x19, 0xcc, 0x02, 0x1e, 0x40, 0x5a, 0x66, 0x57, 0x5a, 0x2f, 0xd2, 0x7a, 0x1d, 0x2a,
	0x37, 0xd8, 0x12, 0x69, 0xa2, 0x34, 0x89, 0x6c, 0xa5, 0xa0, 0xb5, 0xa8, 0xd2, 0x58, 0x63, 0x7b,
	0x32, 0x58, 0xc1, 0x3f, 0x1a, 0x5f, 0x22, 0x53, 0xe6, 0x0d, 0xf2, 0x08, 0x79, 0x1c, 0xca, 0x3c,
	0x41, 0x14, 0xc1, 0x53, 0xa4, 0x8b, 0x3c, 0x24, 0x18, 0x09, 0x29, 0xdd, 0xfd, 0xee, 0x39, 0xc7,
	0x3e, 0x57, 0x83, 0x5d, 0x1e, 0xc3, 0x62, 0x15, 0x58, 0x61, 0x96, 0xd8, 0xcb, 0x98, 0x2f, 0xa0,
	0x00, 0x96, 0xd7, 0xd3, 0x04, 0x04, 0x0d, 0x99, 0x98, 0x84, 0x59, 0x92, 0x64, 0xa9, 0x0d, 0x49,
	0xce, 0x59, 0x5a, 0xcb, 0x56, 0x2e, 0x32, 0xc8, 0x88, 0x7a, 0x58, 0x8c, 0x26, 0x47, 0xdf, 0xe3,
	0x19, 0xcf, 0x6c, 0xe9, 0x08, 0x56, 0x77, 0x92, 0x24, 0xc8, 0x69, 0x9f, 0x1c, 0x3f, 0x22, 0x3c,
	0x70, 0xe2, 0x94, 0x8a, 0xb5, 0x4b, 0x85, 0x88, 0x99, 0x20, 0x16, 0xfe, 0x15, 0xb1, 0x5c, 0xb0,
	0x90, 0x02, 0x8b, 0x7c, 0x60, 0x25, 0xf8, 0x21, 0x94, 0x1a, 0x32, 0x14, 0xb3, 0xef, 0x0d, 0x6b,
	0x69, 0xce, 0x4a, 0x70, 0xa1, 0x24, 0xd7, 0x58, 0x0d, 0x68, 0x11, 0x87, 0xd2, 0xd5, 0x34, 0x90,
	0xd9, 0x9b, 0xfe, 0xb3, 0xea, 0x82, 0x4e, 0xa5, 0xcd, 0xe5, 0x11, 0x9f, 0x7f, 0x70, 0x5a, 0x9b,
	0xd7, 0xff, 0x0d, 0xaf, 0x2b, 0x53, 0x2e, 0x94, 0xe3, 0x77, 0x84, 0xc9, 0xa9, 0x8d, 0xfc, 0xc5,
	0x5d, 0x79, 0xbc, 0x1f, 0x47, 0x1a, 0x32, 0x90, 0xd9, 0xf6, 0x3a, 0x92, 0x67, 0x11, 0xf9, 0x83,
	0x3b, 0x45, 0x4e, 0xd3, 0x4a, 0x69, 0x4a, 0xa5, 0x5d, 0xe1, 0x2c, 0x22, 0x1a, 0xee, 0x14, 0x34,
	0xc9, 0x97, 0x2c, 0xd2, 0x14, 0x03, 0x99, 0x5d, 0xef, 0x0b, 0xc9, 0x1c, 0x0f, 0x02, 0xca, 0x39,
	0xe5, 0xcc, 0x8f, 0x81, 0x25, 0x85, 0xd6, 0x32, 0x14, 0xb3, 0x37, 0xb5, 0xbf, 0xad, 0x6a, 0x39,
	0xfb, 0xc8, 0xac, 0x4a, 0xdc, 0xa4, 0x20, 0xd6, 0x5e, 0x3f, 0x38, 0x5a, 0x8d, 0xae, 0xf0, 0xf0,
	0xc4, 0x42, 0x7e, 0x62, 0xe5, 0x9e, 0xad, 0x65, 0x67, 0xd5, 0xab, 0x46, 0xf2, 0x1b, 0xff, 0x78,
	0xa0, 0xcb, 0x15, 0x93, 0x6d, 0x55, 0x6f, 0x0f, 0x97, 0xcd, 0x0b, 0xe4, 0x9c, 0x6d, 0xb6, 0x3a,
	0x7a, 0xd9, 0xea, 0xe8, 0x6d, 0xab, 0x37, 0x9e, 0x76, 0x7a, 0xe3, 0x79, 0xa7, 0xa3, 0xdb, 0xde,
	0xa1, 0x54, 0x1e, 0x04, 0x6d, 0xf9, 0x52, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x4f,
	0x71, 0x41, 0x2a, 0x02, 0x00, 0x00,
}
