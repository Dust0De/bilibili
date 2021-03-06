// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/share/api/api.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type AddShareRequest struct {
	Oid  int64  `protobuf:"varint,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Mid  int64  `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	Type int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Ip   string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (m *AddShareRequest) Reset()         { *m = AddShareRequest{} }
func (m *AddShareRequest) String() string { return proto.CompactTextString(m) }
func (*AddShareRequest) ProtoMessage()    {}
func (*AddShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_84381552fbdb86d8, []int{0}
}

func (m *AddShareRequest) GetOid() int64 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *AddShareRequest) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *AddShareRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *AddShareRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type AddShareReply struct {
	Shares int64 `protobuf:"varint,1,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (m *AddShareReply) Reset()         { *m = AddShareReply{} }
func (m *AddShareReply) String() string { return proto.CompactTextString(m) }
func (*AddShareReply) ProtoMessage()    {}
func (*AddShareReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_84381552fbdb86d8, []int{1}
}

func (m *AddShareReply) GetShares() int64 {
	if m != nil {
		return m.Shares
	}
	return 0
}

func init() {
	proto.RegisterType((*AddShareRequest)(nil), "share.service.v1.AddShareRequest")
	proto.RegisterType((*AddShareReply)(nil), "share.service.v1.AddShareReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShareClient is the client API for Share service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareClient interface {
	AddShare(ctx context.Context, in *AddShareRequest, opts ...grpc.CallOption) (*AddShareReply, error)
}

type shareClient struct {
	cc *grpc.ClientConn
}

func NewShareClient(cc *grpc.ClientConn) ShareClient {
	return &shareClient{cc}
}

func (c *shareClient) AddShare(ctx context.Context, in *AddShareRequest, opts ...grpc.CallOption) (*AddShareReply, error) {
	out := new(AddShareReply)
	err := c.cc.Invoke(ctx, "/share.service.v1.Share/AddShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareServer is the server API for Share service.
type ShareServer interface {
	AddShare(context.Context, *AddShareRequest) (*AddShareReply, error)
}

func RegisterShareServer(s *grpc.Server, srv ShareServer) {
	s.RegisterService(&_Share_serviceDesc, srv)
}

func _Share_AddShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareServer).AddShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.service.v1.Share/AddShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareServer).AddShare(ctx, req.(*AddShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Share_serviceDesc = grpc.ServiceDesc{
	ServiceName: "share.service.v1.Share",
	HandlerType: (*ShareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddShare",
			Handler:    _Share_AddShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/main/share/api/api.proto",
}

func (m *AddShareRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddShareRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Oid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Oid))
	}
	if m.Mid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
	}
	if m.Type != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Type))
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	return i, nil
}

func (m *AddShareReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddShareReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Shares != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Shares))
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AddShareRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Oid != 0 {
		n += 1 + sovApi(uint64(m.Oid))
	}
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	if m.Type != 0 {
		n += 1 + sovApi(uint64(m.Type))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *AddShareReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Shares != 0 {
		n += 1 + sovApi(uint64(m.Shares))
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddShareRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: AddShareRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddShareRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oid", wireType)
			}
			m.Oid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Oid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *AddShareReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: AddShareReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddShareReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			m.Shares = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shares |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
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
				next, err := skipApi(dAtA[start:])
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
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/main/share/api/api.proto", fileDescriptor_api_84381552fbdb86d8)
}

var fileDescriptor_api_84381552fbdb86d8 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x4b, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xe9, 0x64, 0x66, 0xd0, 0x02, 0x35, 0xf4, 0x42, 0xc2, 0x2c, 0x62, 0x9c, 0x8d, 0x59,
	0x68, 0x37, 0xd1, 0x13, 0xe8, 0x01, 0x5c, 0xc4, 0x85, 0xe8, 0x2e, 0x8f, 0x36, 0x53, 0x30, 0xb1,
	0xcb, 0xbc, 0x20, 0x37, 0x74, 0xe9, 0x11, 0x24, 0x27, 0x91, 0xd4, 0x44, 0x84, 0x59, 0xb8, 0xfb,
	0xff, 0xea, 0xaf, 0x3f, 0xaa, 0xe0, 0x3a, 0x25, 0xd2, 0x8d, 0xa9, 0x7b, 0xcc, 0x8d, 0xae, 0x52,
	0x7c, 0xd7, 0xcd, 0x36, 0xad, 0x8d, 0x4e, 0x09, 0x75, 0x59, 0x53, 0xae, 0xfb, 0x78, 0xca, 0x8a,
	0x6a, 0xdb, 0x5a, 0xe9, 0xf1, 0xa3, 0x9a, 0x79, 0xd5, 0xc7, 0xeb, 0x9b, 0x12, 0xdb, 0x6d, 0x97,
	0xa9, 0xdc, 0x56, 0xba, 0xb4, 0xa5, 0xd5, 0x0c, 0x66, 0xdd, 0x1b, 0x37, 0x2e, 0x9c, 0xf6, 0x82,
	0xcd, 0x0b, 0x9c, 0xdd, 0x17, 0xc5, 0xd3, 0x64, 0x49, 0xcc, 0x47, 0x67, 0x9a, 0x56, 0x7a, 0xe0,
	0x5a, 0x2c, 0x7c, 0x11, 0x8a, 0xc8, 0x4d, 0xa6, 0x38, 0x4d, 0x2a, 0x2c, 0x7c, 0x67, 0x3f, 0xa9,
	0xb0, 0x90, 0x12, 0x16, 0xed, 0x40, 0xc6, 0x77, 0x43, 0x11, 0x2d, 0x13, 0xce, 0xf2, 0x14, 0x1c,
	0x24, 0x7f, 0x11, 0x8a, 0xe8, 0x38, 0x71, 0x90, 0x36, 0x57, 0x70, 0xf2, 0xa7, 0xa6, 0xdd, 0x20,
	0xcf, 0x61, 0xc5, 0xeb, 0x36, 0xb3, 0x7b, 0x6e, 0xb7, 0xcf, 0xb0, 0x64, 0x4a, 0x3e, 0xc2, 0xd1,
	0xef, 0x0f, 0x79, 0xa9, 0x0e, 0x4f, 0x53, 0x07, 0x8b, 0xae, 0x2f, 0xfe, 0x43, 0x68, 0x37, 0x3c,
	0x78, 0x9f, 0x63, 0x20, 0xbe, 0xc6, 0x40, 0x7c, 0x8f, 0x81, 0x78, 0x75, 0xfa, 0x38, 0x5b, 0xf1,
	0xd5, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x7c, 0x30, 0x07, 0x66, 0x01, 0x00, 0x00,
}
