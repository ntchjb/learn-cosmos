// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: learncosmos/gold.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type OrderType int32

const (
	OrderType_BUY  OrderType = 0
	OrderType_SELL OrderType = 1
)

var OrderType_name = map[int32]string{
	0: "BUY",
	1: "SELL",
}

var OrderType_value = map[string]int32{
	"BUY":  0,
	"SELL": 1,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e30836582ee131cc, []int{0}
}

type OrderStatus int32

const (
	OrderStatus_PENDING OrderStatus = 0
	OrderStatus_SUCCESS OrderStatus = 1
	OrderStatus_FAILED  OrderStatus = 2
)

var OrderStatus_name = map[int32]string{
	0: "PENDING",
	1: "SUCCESS",
	2: "FAILED",
}

var OrderStatus_value = map[string]int32{
	"PENDING": 0,
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e30836582ee131cc, []int{1}
}

type GoldPool struct {
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *GoldPool) Reset()         { *m = GoldPool{} }
func (m *GoldPool) String() string { return proto.CompactTextString(m) }
func (*GoldPool) ProtoMessage()    {}
func (*GoldPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e30836582ee131cc, []int{0}
}
func (m *GoldPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoldPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoldPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoldPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoldPool.Merge(m, src)
}
func (m *GoldPool) XXX_Size() int {
	return m.Size()
}
func (m *GoldPool) XXX_DiscardUnknown() {
	xxx_messageInfo_GoldPool.DiscardUnknown(m)
}

var xxx_messageInfo_GoldPool proto.InternalMessageInfo

func (m *GoldPool) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type OwnedGold struct {
	Owner  string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *OwnedGold) Reset()         { *m = OwnedGold{} }
func (m *OwnedGold) String() string { return proto.CompactTextString(m) }
func (*OwnedGold) ProtoMessage()    {}
func (*OwnedGold) Descriptor() ([]byte, []int) {
	return fileDescriptor_e30836582ee131cc, []int{1}
}
func (m *OwnedGold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnedGold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnedGold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnedGold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnedGold.Merge(m, src)
}
func (m *OwnedGold) XXX_Size() int {
	return m.Size()
}
func (m *OwnedGold) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnedGold.DiscardUnknown(m)
}

var xxx_messageInfo_OwnedGold proto.InternalMessageInfo

func (m *OwnedGold) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OwnedGold) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type PoolOrder struct {
	Id           string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         OrderType   `protobuf:"varint,2,opt,name=type,proto3,enum=ntchjb.learncosmos.learncosmos.OrderType" json:"type,omitempty"`
	UserAddr     string      `protobuf:"bytes,3,opt,name=user_addr,json=userAddr,proto3" json:"user_addr,omitempty"`
	PricePerUnit uint64      `protobuf:"varint,4,opt,name=price_per_unit,json=pricePerUnit,proto3" json:"price_per_unit,omitempty"`
	Amount       uint64      `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Status       OrderStatus `protobuf:"varint,6,opt,name=status,proto3,enum=ntchjb.learncosmos.learncosmos.OrderStatus" json:"status,omitempty"`
	StatusReason string      `protobuf:"bytes,7,opt,name=status_reason,json=statusReason,proto3" json:"status_reason,omitempty"`
}

func (m *PoolOrder) Reset()         { *m = PoolOrder{} }
func (m *PoolOrder) String() string { return proto.CompactTextString(m) }
func (*PoolOrder) ProtoMessage()    {}
func (*PoolOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_e30836582ee131cc, []int{2}
}
func (m *PoolOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolOrder.Merge(m, src)
}
func (m *PoolOrder) XXX_Size() int {
	return m.Size()
}
func (m *PoolOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PoolOrder proto.InternalMessageInfo

func (m *PoolOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PoolOrder) GetType() OrderType {
	if m != nil {
		return m.Type
	}
	return OrderType_BUY
}

func (m *PoolOrder) GetUserAddr() string {
	if m != nil {
		return m.UserAddr
	}
	return ""
}

func (m *PoolOrder) GetPricePerUnit() uint64 {
	if m != nil {
		return m.PricePerUnit
	}
	return 0
}

func (m *PoolOrder) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PoolOrder) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_PENDING
}

func (m *PoolOrder) GetStatusReason() string {
	if m != nil {
		return m.StatusReason
	}
	return ""
}

func init() {
	proto.RegisterEnum("ntchjb.learncosmos.learncosmos.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("ntchjb.learncosmos.learncosmos.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*GoldPool)(nil), "ntchjb.learncosmos.learncosmos.GoldPool")
	proto.RegisterType((*OwnedGold)(nil), "ntchjb.learncosmos.learncosmos.OwnedGold")
	proto.RegisterType((*PoolOrder)(nil), "ntchjb.learncosmos.learncosmos.PoolOrder")
}

func init() { proto.RegisterFile("learncosmos/gold.proto", fileDescriptor_e30836582ee131cc) }

var fileDescriptor_e30836582ee131cc = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x93, 0x6c, 0x36, 0x6d, 0xde, 0xae, 0x25, 0x0c, 0xb2, 0x04, 0x84, 0x41, 0xaa, 0x07,
	0x5d, 0x31, 0x81, 0xdd, 0x93, 0x07, 0x0f, 0xbb, 0xdd, 0xb8, 0x14, 0x42, 0x5b, 0x12, 0x7b, 0xd0,
	0x4b, 0x48, 0x33, 0x43, 0x1b, 0x49, 0x67, 0xc2, 0x64, 0x42, 0xed, 0xb7, 0xf0, 0x1b, 0xf8, 0x75,
	0x3c, 0xf6, 0xe8, 0x51, 0xda, 0x2f, 0x22, 0x99, 0x04, 0x89, 0x17, 0xf1, 0xf6, 0xde, 0x7f, 0xde,
	0xff, 0xcf, 0xef, 0x0d, 0x0f, 0xae, 0x0a, 0x9a, 0x0a, 0x96, 0xf1, 0x6a, 0xcb, 0x2b, 0x7f, 0xcd,
	0x0b, 0xe2, 0x95, 0x82, 0x4b, 0x8e, 0x30, 0x93, 0xd9, 0xe6, 0xcb, 0xca, 0xeb, 0x3d, 0xf7, 0xeb,
	0xf1, 0x18, 0x86, 0x8f, 0xbc, 0x20, 0x0b, 0xce, 0x0b, 0x74, 0x05, 0x56, 0xba, 0xe5, 0x35, 0x93,
	0xae, 0xfe, 0x5c, 0x7f, 0x65, 0x46, 0x5d, 0x37, 0x7e, 0x07, 0xf6, 0x7c, 0xc7, 0x28, 0x69, 0x06,
	0xd1, 0x53, 0x38, 0xe7, 0x3b, 0x46, 0x85, 0x9a, 0xb1, 0xa3, 0xb6, 0xe9, 0x59, 0x8d, 0xbf, 0xac,
	0xdf, 0x0d, 0xb0, 0x9b, 0xec, 0xb9, 0x20, 0x54, 0xa0, 0x11, 0x18, 0x39, 0xe9, 0x8c, 0x46, 0x4e,
	0xd0, 0x7b, 0x30, 0xe5, 0xbe, 0xa4, 0xca, 0x33, 0xba, 0x79, 0xed, 0xfd, 0x9b, 0xd5, 0x53, 0x21,
	0x1f, 0xf7, 0x25, 0x8d, 0x94, 0x0d, 0x3d, 0x03, 0xbb, 0xae, 0xa8, 0x48, 0x52, 0x42, 0x84, 0x7b,
	0xa6, 0x52, 0x87, 0x8d, 0x70, 0x47, 0x88, 0x40, 0x2f, 0x61, 0x54, 0x8a, 0x3c, 0xa3, 0x49, 0x49,
	0x45, 0x52, 0xb3, 0x5c, 0xba, 0xa6, 0x22, 0xbb, 0x54, 0xea, 0x82, 0x8a, 0x25, 0xcb, 0x65, 0x8f,
	0xfb, 0xbc, 0xcf, 0x8d, 0x26, 0x60, 0x55, 0x32, 0x95, 0x75, 0xe5, 0x5a, 0x8a, 0xed, 0xcd, 0x7f,
	0xb1, 0xc5, 0xca, 0x12, 0x75, 0x56, 0xf4, 0x02, 0x9e, 0xb4, 0x55, 0x22, 0x68, 0x5a, 0x71, 0xe6,
	0x0e, 0x14, 0xe3, 0x65, 0x2b, 0x46, 0x4a, 0xbb, 0xc6, 0x60, 0xff, 0xd9, 0x0b, 0x0d, 0xe0, 0xec,
	0x7e, 0xf9, 0xc9, 0xd1, 0xd0, 0x10, 0xcc, 0x38, 0x08, 0x43, 0x47, 0xbf, 0xbe, 0x85, 0x8b, 0x5e,
	0x36, 0xba, 0x80, 0xc1, 0x22, 0x98, 0x3d, 0x4c, 0x67, 0x8f, 0x8e, 0xd6, 0x34, 0xf1, 0x72, 0x32,
	0x09, 0xe2, 0xd8, 0xd1, 0x11, 0x80, 0xf5, 0xe1, 0x6e, 0x1a, 0x06, 0x0f, 0x8e, 0x71, 0x1f, 0xfe,
	0x38, 0x62, 0xfd, 0x70, 0xc4, 0xfa, 0xaf, 0x23, 0xd6, 0xbf, 0x9d, 0xb0, 0x76, 0x38, 0x61, 0xed,
	0xe7, 0x09, 0x6b, 0x9f, 0x6f, 0xd6, 0xb9, 0xdc, 0xd4, 0x2b, 0x2f, 0xe3, 0x5b, 0xbf, 0x5d, 0xc9,
	0x57, 0x6b, 0xbc, 0xed, 0x4e, 0xe7, 0xab, 0xdf, 0x3f, 0xa4, 0xe6, 0x9b, 0xab, 0x95, 0xa5, 0x4e,
	0xe9, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xc9, 0x72, 0x82, 0x64, 0x02, 0x00, 0x00,
}

func (m *GoldPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoldPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoldPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OwnedGold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnedGold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnedGold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGold(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StatusReason) > 0 {
		i -= len(m.StatusReason)
		copy(dAtA[i:], m.StatusReason)
		i = encodeVarintGold(dAtA, i, uint64(len(m.StatusReason)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.Amount != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if m.PricePerUnit != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.PricePerUnit))
		i--
		dAtA[i] = 0x20
	}
	if len(m.UserAddr) > 0 {
		i -= len(m.UserAddr)
		copy(dAtA[i:], m.UserAddr)
		i = encodeVarintGold(dAtA, i, uint64(len(m.UserAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintGold(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGold(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGold(dAtA []byte, offset int, v uint64) int {
	offset -= sovGold(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoldPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovGold(uint64(m.Amount))
	}
	return n
}

func (m *OwnedGold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGold(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovGold(uint64(m.Amount))
	}
	return n
}

func (m *PoolOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGold(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovGold(uint64(m.Type))
	}
	l = len(m.UserAddr)
	if l > 0 {
		n += 1 + l + sovGold(uint64(l))
	}
	if m.PricePerUnit != 0 {
		n += 1 + sovGold(uint64(m.PricePerUnit))
	}
	if m.Amount != 0 {
		n += 1 + sovGold(uint64(m.Amount))
	}
	if m.Status != 0 {
		n += 1 + sovGold(uint64(m.Status))
	}
	l = len(m.StatusReason)
	if l > 0 {
		n += 1 + l + sovGold(uint64(l))
	}
	return n
}

func sovGold(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGold(x uint64) (n int) {
	return sovGold(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoldPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGold
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
			return fmt.Errorf("proto: GoldPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoldPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGold(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGold
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
func (m *OwnedGold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGold
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
			return fmt.Errorf("proto: OwnedGold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnedGold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
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
				return ErrInvalidLengthGold
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGold(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGold
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
func (m *PoolOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGold
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
			return fmt.Errorf("proto: PoolOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
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
				return ErrInvalidLengthGold
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= OrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
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
				return ErrInvalidLengthGold
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PricePerUnit", wireType)
			}
			m.PricePerUnit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PricePerUnit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OrderStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGold
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
				return ErrInvalidLengthGold
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGold(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGold
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
func skipGold(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGold
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
					return 0, ErrIntOverflowGold
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
					return 0, ErrIntOverflowGold
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
				return 0, ErrInvalidLengthGold
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGold
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGold
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGold        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGold          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGold = fmt.Errorf("proto: unexpected end of group")
)
