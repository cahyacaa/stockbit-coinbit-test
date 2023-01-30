// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deposit.proto

package deposit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Deposit struct {
	WalletID             string   `protobuf:"bytes,1,opt,name=walletID,proto3" json:"walletID,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d647de60f1ae88, []int{0}
}

func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deposit.Unmarshal(m, b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return xxx_messageInfo_Deposit.Size(m)
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetWalletID() string {
	if m != nil {
		return m.WalletID
	}
	return ""
}

func (m *Deposit) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Balance struct {
	WalletID             string   `protobuf:"bytes,1,opt,name=walletID,proto3" json:"walletID,omitempty"`
	Amount               float32  `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Balance              float32  `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	IsAboveThreshold     bool     `protobuf:"varint,4,opt,name=isAboveThreshold,proto3" json:"isAboveThreshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d647de60f1ae88, []int{1}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetWalletID() string {
	if m != nil {
		return m.WalletID
	}
	return ""
}

func (m *Balance) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Balance) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Balance) GetIsAboveThreshold() bool {
	if m != nil {
		return m.IsAboveThreshold
	}
	return false
}

type DepositFlagger struct {
	WalletID             string                 `protobuf:"bytes,1,opt,name=walletID,proto3" json:"walletID,omitempty"`
	Amount               float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	TimeWindowBalance    float32                `protobuf:"fixed32,3,opt,name=timeWindowBalance,proto3" json:"timeWindowBalance,omitempty"`
	IsAboveThreshold     bool                   `protobuf:"varint,4,opt,name=isAboveThreshold,proto3" json:"isAboveThreshold,omitempty"`
	TimeExpired          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=TimeExpired,proto3" json:"TimeExpired,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DepositFlagger) Reset()         { *m = DepositFlagger{} }
func (m *DepositFlagger) String() string { return proto.CompactTextString(m) }
func (*DepositFlagger) ProtoMessage()    {}
func (*DepositFlagger) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d647de60f1ae88, []int{2}
}

func (m *DepositFlagger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepositFlagger.Unmarshal(m, b)
}
func (m *DepositFlagger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepositFlagger.Marshal(b, m, deterministic)
}
func (m *DepositFlagger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositFlagger.Merge(m, src)
}
func (m *DepositFlagger) XXX_Size() int {
	return xxx_messageInfo_DepositFlagger.Size(m)
}
func (m *DepositFlagger) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositFlagger.DiscardUnknown(m)
}

var xxx_messageInfo_DepositFlagger proto.InternalMessageInfo

func (m *DepositFlagger) GetWalletID() string {
	if m != nil {
		return m.WalletID
	}
	return ""
}

func (m *DepositFlagger) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DepositFlagger) GetTimeWindowBalance() float32 {
	if m != nil {
		return m.TimeWindowBalance
	}
	return 0
}

func (m *DepositFlagger) GetIsAboveThreshold() bool {
	if m != nil {
		return m.IsAboveThreshold
	}
	return false
}

func (m *DepositFlagger) GetTimeExpired() *timestamppb.Timestamp {
	if m != nil {
		return m.TimeExpired
	}
	return nil
}

func init() {
	proto.RegisterType((*Deposit)(nil), "Deposit")
	proto.RegisterType((*Balance)(nil), "Balance")
	proto.RegisterType((*DepositFlagger)(nil), "DepositFlagger")
}

func init() {
	proto.RegisterFile("deposit.proto", fileDescriptor_c2d647de60f1ae88)
}

var fileDescriptor_c2d647de60f1ae88 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0xc8,
	0x2f, 0xce, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73,
	0x0b, 0x20, 0x0a, 0x94, 0x6c, 0xb9, 0xd8, 0x5d, 0x20, 0x3a, 0x84, 0xa4, 0xb8, 0x38, 0xca, 0x13,
	0x73, 0x72, 0x52, 0x4b, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21,
	0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xa6, 0x20,
	0x28, 0x4f, 0xa9, 0x99, 0x91, 0x8b, 0xdd, 0x29, 0x31, 0x27, 0x31, 0x2f, 0x39, 0x95, 0x1c, 0xfd,
	0x42, 0x12, 0x5c, 0xec, 0x49, 0x10, 0xed, 0x12, 0xcc, 0x60, 0x09, 0x18, 0x57, 0x48, 0x8b, 0x4b,
	0x20, 0xb3, 0xd8, 0x31, 0x29, 0xbf, 0x2c, 0x35, 0x24, 0xa3, 0x28, 0xb5, 0x38, 0x23, 0x3f, 0x27,
	0x45, 0x82, 0x45, 0x81, 0x51, 0x83, 0x23, 0x08, 0x43, 0x5c, 0xe9, 0x0e, 0x23, 0x17, 0x1f, 0xd4,
	0x17, 0x6e, 0x39, 0x89, 0xe9, 0xe9, 0xa9, 0x45, 0x64, 0x39, 0x46, 0x87, 0x4b, 0x10, 0x14, 0x3c,
	0xe1, 0x99, 0x79, 0x29, 0xf9, 0xe5, 0x4e, 0x28, 0xce, 0xc2, 0x94, 0x20, 0xc5, 0x81, 0x42, 0x36,
	0x5c, 0xdc, 0x21, 0x99, 0xb9, 0xa9, 0xae, 0x15, 0x05, 0x99, 0x45, 0xa9, 0x29, 0x12, 0xac, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x90, 0xc8, 0xd1, 0x83, 0x45, 0x8e, 0x5e, 0x08, 0x2c, 0x72,
	0x82, 0x90, 0x95, 0x27, 0xb1, 0x81, 0x15, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0x82,
	0x98, 0xc9, 0xdc, 0x01, 0x00, 0x00,
}