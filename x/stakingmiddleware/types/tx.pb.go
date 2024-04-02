// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: composable/stakingmiddleware/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
//
// Since: cosmos-sdk 0.47
type MsgUpdateEpochParams struct {
	// authority is the address that controls the module (defaults to x/gov unless
	// overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/stakingmiddleware parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateEpochParams) Reset()         { *m = MsgUpdateEpochParams{} }
func (m *MsgUpdateEpochParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateEpochParams) ProtoMessage()    {}
func (*MsgUpdateEpochParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ef8bcd16d104cd6, []int{0}
}
func (m *MsgUpdateEpochParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateEpochParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateEpochParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateEpochParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateEpochParams.Merge(m, src)
}
func (m *MsgUpdateEpochParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateEpochParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateEpochParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateEpochParams proto.InternalMessageInfo

func (m *MsgUpdateEpochParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateEpochParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
//
// Since: cosmos-sdk 0.47
type MsgUpdateParamsEpochResponse struct {
}

func (m *MsgUpdateParamsEpochResponse) Reset()         { *m = MsgUpdateParamsEpochResponse{} }
func (m *MsgUpdateParamsEpochResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsEpochResponse) ProtoMessage()    {}
func (*MsgUpdateParamsEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ef8bcd16d104cd6, []int{1}
}
func (m *MsgUpdateParamsEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsEpochResponse.Merge(m, src)
}
func (m *MsgUpdateParamsEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsEpochResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateEpochParams)(nil), "composable.stakingmiddleware.v1beta1.MsgUpdateEpochParams")
	proto.RegisterType((*MsgUpdateParamsEpochResponse)(nil), "composable.stakingmiddleware.v1beta1.MsgUpdateParamsEpochResponse")
}

func init() {
	proto.RegisterFile("composable/stakingmiddleware/v1beta1/tx.proto", fileDescriptor_4ef8bcd16d104cd6)
}

var fileDescriptor_4ef8bcd16d104cd6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xce, 0xcf, 0x2d,
	0xc8, 0x2f, 0x4e, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xcf, 0xcd,
	0x4c, 0x49, 0xc9, 0x49, 0x2d, 0x4f, 0x2c, 0x4a, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x41, 0x28, 0xd7, 0xc3, 0x50,
	0xae, 0x07, 0x55, 0x2e, 0x25, 0x9e, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x5b, 0x9c, 0xae,
	0x5f, 0x66, 0x08, 0xa2, 0x20, 0xda, 0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24,
	0x54, 0x48, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0xa0, 0xa2, 0x92, 0x10, 0x13,
	0xe2, 0x21, 0x12, 0x10, 0x0e, 0x54, 0xca, 0x86, 0x28, 0x17, 0x63, 0x3a, 0x0e, 0xac, 0x5b, 0xe9,
	0x11, 0x23, 0x97, 0x88, 0x6f, 0x71, 0x7a, 0x68, 0x41, 0x4a, 0x62, 0x49, 0xaa, 0x6b, 0x41, 0x7e,
	0x72, 0x46, 0x40, 0x62, 0x51, 0x62, 0x6e, 0xb1, 0x90, 0x19, 0x17, 0x67, 0x62, 0x69, 0x49, 0x46,
	0x7e, 0x51, 0x66, 0x49, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc4, 0xa5, 0x2d, 0xba,
	0x22, 0x50, 0xbb, 0x1d, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2,
	0x83, 0x10, 0x4a, 0x85, 0xfc, 0xb9, 0xd8, 0x0a, 0xc0, 0x26, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x1b, 0xe9, 0xe8, 0x11, 0x13, 0x44, 0x7a, 0x10, 0x5b, 0x9d, 0x38, 0x4f, 0xdc, 0x93, 0x67, 0x58,
	0xf1, 0x7c, 0x83, 0x16, 0x63, 0x10, 0xd4, 0x18, 0x2b, 0xc7, 0xa6, 0xe7, 0x1b, 0xb4, 0x10, 0x16,
	0x74, 0x3d, 0xdf, 0xa0, 0x85, 0x64, 0xa4, 0x7e, 0x05, 0x16, 0x4f, 0xc3, 0x3d, 0x04, 0x31, 0x55,
	0x49, 0x8e, 0x4b, 0x06, 0x4d, 0x08, 0xec, 0xd3, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0xa3, 0x45, 0x8c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x93, 0x19, 0xb9, 0x04, 0x31, 0x43, 0xc2,
	0x8a, 0x38, 0x1f, 0x60, 0x0b, 0x45, 0x29, 0x27, 0x12, 0xf5, 0x62, 0x71, 0x9d, 0x14, 0x6b, 0x03,
	0x28, 0x3c, 0x9c, 0x8c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x12,
	0x5b, 0x18, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x63, 0xd9, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0xc6, 0x61, 0xc6, 0xd7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateEpochParams(ctx context.Context, in *MsgUpdateEpochParams, opts ...grpc.CallOption) (*MsgUpdateParamsEpochResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateEpochParams(ctx context.Context, in *MsgUpdateEpochParams, opts ...grpc.CallOption) (*MsgUpdateParamsEpochResponse, error) {
	out := new(MsgUpdateParamsEpochResponse)
	err := c.cc.Invoke(ctx, "/composable.stakingmiddleware.v1beta1.Msg/UpdateEpochParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateEpochParams(context.Context, *MsgUpdateEpochParams) (*MsgUpdateParamsEpochResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateEpochParams(ctx context.Context, req *MsgUpdateEpochParams) (*MsgUpdateParamsEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEpochParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateEpochParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateEpochParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateEpochParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/composable.stakingmiddleware.v1beta1.Msg/UpdateEpochParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateEpochParams(ctx, req.(*MsgUpdateEpochParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "composable.stakingmiddleware.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateEpochParams",
			Handler:    _Msg_UpdateEpochParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "composable/stakingmiddleware/v1beta1/tx.proto",
}

func (m *MsgUpdateEpochParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateEpochParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateEpochParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateEpochParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateEpochParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateEpochParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateEpochParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsEpochResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)