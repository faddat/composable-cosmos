// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: composable/ibctransfermiddleware/v1beta1/ibctransfermiddleware.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params holds parameters for the ibctransfermiddleware module.
type Params struct {
	ChannelFees []*ChannelFee `protobuf:"bytes,1,rep,name=channel_fees,json=channelFees,proto3" json:"channel_fees,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1193893bc248bc1b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetChannelFees() []*ChannelFee {
	if m != nil {
		return m.ChannelFees
	}
	return nil
}

type ChannelFee struct {
	Channel             string      `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	AllowedTokens       []*CoinItem `protobuf:"bytes,2,rep,name=allowed_tokens,json=allowedTokens,proto3" json:"allowed_tokens,omitempty"`
	FeeAddress          string      `protobuf:"bytes,3,opt,name=fee_address,json=feeAddress,proto3" json:"fee_address,omitempty"`
	MinTimeoutTimestamp int64       `protobuf:"varint,4,opt,name=min_timeout_timestamp,json=minTimeoutTimestamp,proto3" json:"min_timeout_timestamp,omitempty"`
}

func (m *ChannelFee) Reset()         { *m = ChannelFee{} }
func (m *ChannelFee) String() string { return proto.CompactTextString(m) }
func (*ChannelFee) ProtoMessage()    {}
func (*ChannelFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_1193893bc248bc1b, []int{1}
}
func (m *ChannelFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChannelFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChannelFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChannelFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelFee.Merge(m, src)
}
func (m *ChannelFee) XXX_Size() int {
	return m.Size()
}
func (m *ChannelFee) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelFee.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelFee proto.InternalMessageInfo

func (m *ChannelFee) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChannelFee) GetAllowedTokens() []*CoinItem {
	if m != nil {
		return m.AllowedTokens
	}
	return nil
}

func (m *ChannelFee) GetFeeAddress() string {
	if m != nil {
		return m.FeeAddress
	}
	return ""
}

func (m *ChannelFee) GetMinTimeoutTimestamp() int64 {
	if m != nil {
		return m.MinTimeoutTimestamp
	}
	return 0
}

type CoinItem struct {
	MinFee     types.Coin `protobuf:"bytes,1,opt,name=min_fee,json=minFee,proto3" json:"min_fee"`
	Percentage int64      `protobuf:"varint,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
}

func (m *CoinItem) Reset()         { *m = CoinItem{} }
func (m *CoinItem) String() string { return proto.CompactTextString(m) }
func (*CoinItem) ProtoMessage()    {}
func (*CoinItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1193893bc248bc1b, []int{2}
}
func (m *CoinItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinItem.Merge(m, src)
}
func (m *CoinItem) XXX_Size() int {
	return m.Size()
}
func (m *CoinItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinItem.DiscardUnknown(m)
}

var xxx_messageInfo_CoinItem proto.InternalMessageInfo

func (m *CoinItem) GetMinFee() types.Coin {
	if m != nil {
		return m.MinFee
	}
	return types.Coin{}
}

func (m *CoinItem) GetPercentage() int64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "composable.ibctransfermiddleware.v1beta1.Params")
	proto.RegisterType((*ChannelFee)(nil), "composable.ibctransfermiddleware.v1beta1.ChannelFee")
	proto.RegisterType((*CoinItem)(nil), "composable.ibctransfermiddleware.v1beta1.CoinItem")
}

func init() {
	proto.RegisterFile("composable/ibctransfermiddleware/v1beta1/ibctransfermiddleware.proto", fileDescriptor_1193893bc248bc1b)
}

var fileDescriptor_1193893bc248bc1b = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x27, 0x3b, 0x32, 0xeb, 0x66, 0x54, 0xb0, 0xae, 0x98, 0x5d, 0x30, 0x0e, 0x73, 0x1a,
	0x04, 0x5b, 0xa6, 0x0a, 0xe2, 0xc1, 0x83, 0xa3, 0x08, 0xde, 0xa4, 0x0e, 0x88, 0x5e, 0x4a, 0xda,
	0xbe, 0xd6, 0x60, 0x93, 0x94, 0x24, 0xee, 0xea, 0xb7, 0xf0, 0x63, 0x78, 0xf4, 0xe0, 0x87, 0xd8,
	0xe3, 0xe2, 0xc9, 0x93, 0xc8, 0xcc, 0xc1, 0xab, 0x1f, 0x41, 0x9a, 0x64, 0x1c, 0x85, 0x15, 0xf4,
	0xd2, 0xe4, 0xbd, 0x5f, 0xff, 0xef, 0xfd, 0x93, 0x3c, 0xfc, 0xa8, 0x54, 0xa2, 0x53, 0x86, 0x15,
	0x2d, 0x24, 0xbc, 0x28, 0xad, 0x66, 0xd2, 0xd4, 0xa0, 0x05, 0xaf, 0xaa, 0x16, 0x8e, 0x99, 0x86,
	0xe4, 0x68, 0x5e, 0x80, 0x65, 0xf3, 0xb3, 0x69, 0xdc, 0x69, 0x65, 0x55, 0x34, 0xdb, 0x56, 0x89,
	0xcf, 0xfe, 0x2f, 0x54, 0x39, 0xdc, 0x6f, 0x54, 0xa3, 0x9c, 0x28, 0xe9, 0x77, 0x5e, 0x7f, 0x78,
	0x50, 0x2a, 0x23, 0x94, 0xc9, 0x3d, 0xf0, 0x41, 0x40, 0xd4, 0x47, 0x49, 0xc1, 0xcc, 0xd6, 0x4b,
	0xa9, 0xb8, 0x0c, 0xfc, 0x32, 0x13, 0x5c, 0xaa, 0xc4, 0x7d, 0x43, 0xea, 0x5a, 0x90, 0x08, 0xd3,
	0x24, 0x47, 0xf3, 0x7e, 0xf1, 0x60, 0xca, 0xf0, 0xe8, 0x29, 0xd3, 0x4c, 0x98, 0xe8, 0x39, 0xbe,
	0x50, 0xbe, 0x62, 0x52, 0x42, 0x9b, 0xd7, 0x00, 0x86, 0xa0, 0xc9, 0x70, 0x36, 0x4e, 0xef, 0xc4,
	0xff, 0x7a, 0x8e, 0xf8, 0xa1, 0x57, 0x3f, 0x06, 0xc8, 0xc6, 0xe5, 0xaf, 0xbd, 0x99, 0xfe, 0x40,
	0x18, 0x6f, 0x59, 0x44, 0xf0, 0x6e, 0xa0, 0x04, 0x4d, 0xd0, 0x6c, 0x2f, 0xdb, 0x84, 0xd1, 0x0b,
	0x7c, 0x89, 0xb5, 0xad, 0x3a, 0x86, 0x2a, 0xb7, 0xea, 0x35, 0x48, 0x43, 0x76, 0x9c, 0x87, 0xf4,
	0x3f, 0x3c, 0x28, 0x2e, 0x9f, 0x58, 0x10, 0xd9, 0xc5, 0x50, 0x69, 0xe9, 0x0a, 0x45, 0xf7, 0xf0,
	0xb8, 0x06, 0xc8, 0x59, 0x55, 0x69, 0x30, 0x86, 0x0c, 0xfb, 0xc6, 0x0b, 0xf2, 0xf9, 0xd3, 0xad,
	0xfd, 0x70, 0xb3, 0x0f, 0x3c, 0x79, 0x66, 0x35, 0x97, 0x4d, 0x86, 0x6b, 0x80, 0x90, 0x89, 0x52,
	0x7c, 0x55, 0x70, 0x99, 0x5b, 0x2e, 0x40, 0xbd, 0xb1, 0x6e, 0x35, 0x96, 0x89, 0x8e, 0x9c, 0x9b,
	0xa0, 0xd9, 0x30, 0xbb, 0x22, 0xb8, 0x5c, 0x7a, 0xb6, 0xdc, 0xa0, 0x29, 0xc7, 0xe7, 0x37, 0x4e,
	0xa2, 0xfb, 0x78, 0xb7, 0xd7, 0xd7, 0x00, 0xee, 0xbc, 0xe3, 0xf4, 0x20, 0x0e, 0x3d, 0xfb, 0xf7,
	0xfb, 0xc3, 0xf9, 0x62, 0xef, 0xe4, 0xeb, 0x8d, 0xc1, 0x87, 0xef, 0x1f, 0x6f, 0xa2, 0x6c, 0x24,
	0xb8, 0xec, 0xaf, 0x8b, 0x62, 0xdc, 0x81, 0x2e, 0x41, 0x5a, 0xd6, 0x00, 0xd9, 0x71, 0x3d, 0x7f,
	0xcb, 0x2c, 0xee, 0x9e, 0xac, 0x28, 0x3a, 0x5d, 0x51, 0xf4, 0x6d, 0x45, 0xd1, 0xfb, 0x35, 0x1d,
	0x9c, 0xae, 0xe9, 0xe0, 0xcb, 0x9a, 0x0e, 0x5e, 0x5e, 0x7f, 0xfb, 0x97, 0xf1, 0xb5, 0xef, 0x3a,
	0x30, 0xc5, 0xc8, 0x0d, 0xc0, 0xed, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x03, 0x6b, 0xef,
	0xef, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelFees) > 0 {
		for iNdEx := len(m.ChannelFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChannelFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChannelFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChannelFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinTimeoutTimestamp != 0 {
		i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(m.MinTimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.FeeAddress) > 0 {
		i -= len(m.FeeAddress)
		copy(dAtA[i:], m.FeeAddress)
		i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(len(m.FeeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AllowedTokens) > 0 {
		for iNdEx := len(m.AllowedTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoinItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Percentage != 0 {
		i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(m.Percentage))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.MinFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIbctransfermiddleware(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIbctransfermiddleware(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbctransfermiddleware(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChannelFees) > 0 {
		for _, e := range m.ChannelFees {
			l = e.Size()
			n += 1 + l + sovIbctransfermiddleware(uint64(l))
		}
	}
	return n
}

func (m *ChannelFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbctransfermiddleware(uint64(l))
	}
	if len(m.AllowedTokens) > 0 {
		for _, e := range m.AllowedTokens {
			l = e.Size()
			n += 1 + l + sovIbctransfermiddleware(uint64(l))
		}
	}
	l = len(m.FeeAddress)
	if l > 0 {
		n += 1 + l + sovIbctransfermiddleware(uint64(l))
	}
	if m.MinTimeoutTimestamp != 0 {
		n += 1 + sovIbctransfermiddleware(uint64(m.MinTimeoutTimestamp))
	}
	return n
}

func (m *CoinItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinFee.Size()
	n += 1 + l + sovIbctransfermiddleware(uint64(l))
	if m.Percentage != 0 {
		n += 1 + sovIbctransfermiddleware(uint64(m.Percentage))
	}
	return n
}

func sovIbctransfermiddleware(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbctransfermiddleware(x uint64) (n int) {
	return sovIbctransfermiddleware(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbctransfermiddleware
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
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
				return ErrInvalidLengthIbctransfermiddleware
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbctransfermiddleware
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelFees = append(m.ChannelFees, &ChannelFee{})
			if err := m.ChannelFees[len(m.ChannelFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbctransfermiddleware(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbctransfermiddleware
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
func (m *ChannelFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbctransfermiddleware
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
			return fmt.Errorf("proto: ChannelFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
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
				return ErrInvalidLengthIbctransfermiddleware
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbctransfermiddleware
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
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
				return ErrInvalidLengthIbctransfermiddleware
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbctransfermiddleware
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedTokens = append(m.AllowedTokens, &CoinItem{})
			if err := m.AllowedTokens[len(m.AllowedTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
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
				return ErrInvalidLengthIbctransfermiddleware
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbctransfermiddleware
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTimeoutTimestamp", wireType)
			}
			m.MinTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTimeoutTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIbctransfermiddleware(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbctransfermiddleware
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
func (m *CoinItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbctransfermiddleware
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
			return fmt.Errorf("proto: CoinItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
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
				return ErrInvalidLengthIbctransfermiddleware
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIbctransfermiddleware
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			m.Percentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbctransfermiddleware
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Percentage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIbctransfermiddleware(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbctransfermiddleware
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
func skipIbctransfermiddleware(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbctransfermiddleware
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
					return 0, ErrIntOverflowIbctransfermiddleware
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
					return 0, ErrIntOverflowIbctransfermiddleware
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
				return 0, ErrInvalidLengthIbctransfermiddleware
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbctransfermiddleware
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbctransfermiddleware
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbctransfermiddleware        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbctransfermiddleware          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbctransfermiddleware = fmt.Errorf("proto: unexpected end of group")
)