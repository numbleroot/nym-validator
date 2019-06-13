// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/elgamal/proto/types.proto

package elgamal

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ProtoEncryption struct {
	C1                   []byte   `protobuf:"bytes,1,opt,name=C1,json=c1,proto3" json:"C1,omitempty"`
	C2                   []byte   `protobuf:"bytes,2,opt,name=C2,json=c2,proto3" json:"C2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoEncryption) Reset()         { *m = ProtoEncryption{} }
func (m *ProtoEncryption) String() string { return proto.CompactTextString(m) }
func (*ProtoEncryption) ProtoMessage()    {}
func (*ProtoEncryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4e4601cca9348f7, []int{0}
}

func (m *ProtoEncryption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoEncryption.Unmarshal(m, b)
}
func (m *ProtoEncryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoEncryption.Marshal(b, m, deterministic)
}
func (m *ProtoEncryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoEncryption.Merge(m, src)
}
func (m *ProtoEncryption) XXX_Size() int {
	return xxx_messageInfo_ProtoEncryption.Size(m)
}
func (m *ProtoEncryption) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoEncryption.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoEncryption proto.InternalMessageInfo

func (m *ProtoEncryption) GetC1() []byte {
	if m != nil {
		return m.C1
	}
	return nil
}

func (m *ProtoEncryption) GetC2() []byte {
	if m != nil {
		return m.C2
	}
	return nil
}

type ProtoPublicKey struct {
	P                    []byte   `protobuf:"bytes,1,opt,name=P,json=p,proto3" json:"P,omitempty"`
	G                    []byte   `protobuf:"bytes,2,opt,name=G,json=g,proto3" json:"G,omitempty"`
	Gamma                []byte   `protobuf:"bytes,3,opt,name=Gamma,json=gamma,proto3" json:"Gamma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoPublicKey) Reset()         { *m = ProtoPublicKey{} }
func (m *ProtoPublicKey) String() string { return proto.CompactTextString(m) }
func (*ProtoPublicKey) ProtoMessage()    {}
func (*ProtoPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4e4601cca9348f7, []int{1}
}

func (m *ProtoPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoPublicKey.Unmarshal(m, b)
}
func (m *ProtoPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoPublicKey.Marshal(b, m, deterministic)
}
func (m *ProtoPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoPublicKey.Merge(m, src)
}
func (m *ProtoPublicKey) XXX_Size() int {
	return xxx_messageInfo_ProtoPublicKey.Size(m)
}
func (m *ProtoPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoPublicKey proto.InternalMessageInfo

func (m *ProtoPublicKey) GetP() []byte {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *ProtoPublicKey) GetG() []byte {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *ProtoPublicKey) GetGamma() []byte {
	if m != nil {
		return m.Gamma
	}
	return nil
}

// just to keep it consistent when/if writing to a PEM file
type ProtoPrivateKey struct {
	D                    []byte   `protobuf:"bytes,1,opt,name=D,json=d,proto3" json:"D,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoPrivateKey) Reset()         { *m = ProtoPrivateKey{} }
func (m *ProtoPrivateKey) String() string { return proto.CompactTextString(m) }
func (*ProtoPrivateKey) ProtoMessage()    {}
func (*ProtoPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4e4601cca9348f7, []int{2}
}

func (m *ProtoPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoPrivateKey.Unmarshal(m, b)
}
func (m *ProtoPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoPrivateKey.Marshal(b, m, deterministic)
}
func (m *ProtoPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoPrivateKey.Merge(m, src)
}
func (m *ProtoPrivateKey) XXX_Size() int {
	return xxx_messageInfo_ProtoPrivateKey.Size(m)
}
func (m *ProtoPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoPrivateKey proto.InternalMessageInfo

func (m *ProtoPrivateKey) GetD() []byte {
	if m != nil {
		return m.D
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoEncryption)(nil), "elgamal.ProtoEncryption")
	proto.RegisterType((*ProtoPublicKey)(nil), "elgamal.ProtoPublicKey")
	proto.RegisterType((*ProtoPrivateKey)(nil), "elgamal.ProtoPrivateKey")
}

func init() { proto.RegisterFile("crypto/elgamal/proto/types.proto", fileDescriptor_a4e4601cca9348f7) }

var fileDescriptor_a4e4601cca9348f7 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x4f, 0xcd, 0x49, 0x4f, 0xcc, 0x4d, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x03, 0xb3, 0x85, 0xd8, 0xa1, 0x52, 0x4a, 0x86, 0x5c,
	0xfc, 0x01, 0x20, 0x11, 0xd7, 0x3c, 0xb0, 0x9e, 0xcc, 0xfc, 0x3c, 0x21, 0x3e, 0x2e, 0x26, 0x67,
	0x43, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0xa6, 0x64, 0x43, 0x30, 0xdf, 0x48, 0x82, 0x09,
	0xca, 0x37, 0x52, 0x72, 0xe2, 0xe2, 0x03, 0x6b, 0x09, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xf6, 0x4e,
	0xad, 0x14, 0xe2, 0xe1, 0x62, 0x0c, 0x80, 0x6a, 0x60, 0x2c, 0x00, 0xf1, 0xdc, 0xa1, 0xca, 0x19,
	0xd3, 0x85, 0x44, 0xb8, 0x58, 0xdd, 0x13, 0x73, 0x73, 0x13, 0x25, 0x98, 0xc1, 0x22, 0xac, 0xe9,
	0x20, 0x8e, 0x92, 0x3c, 0xd4, 0xda, 0x80, 0xa2, 0xcc, 0xb2, 0xc4, 0x92, 0x54, 0xa8, 0x21, 0x2e,
	0x30, 0x43, 0x52, 0x9c, 0xd4, 0xa3, 0x54, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xf3, 0x2a, 0x73, 0x4b, 0x52, 0x93, 0x33, 0x40, 0xb4, 0x3e, 0xaa, 0xdf, 0x92, 0xd8,
	0xc0, 0x1e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xcc, 0xd0, 0x93, 0xf4, 0x00, 0x00,
	0x00,
}
