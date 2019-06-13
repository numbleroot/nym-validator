// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/coconut/scheme/proto/types.proto

package coconut

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	elgamal "github.com/nymtech/nym/crypto/elgamal"
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

// in principle should never be used as secret key would never be sent over the wire,
// but the definition is included for completion sake
type ProtoSecretKey struct {
	X                    []byte   `protobuf:"bytes,1,opt,name=X,json=x,proto3" json:"X,omitempty"`
	Y                    [][]byte `protobuf:"bytes,2,rep,name=Y,json=y,proto3" json:"Y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoSecretKey) Reset()         { *m = ProtoSecretKey{} }
func (m *ProtoSecretKey) String() string { return proto.CompactTextString(m) }
func (*ProtoSecretKey) ProtoMessage()    {}
func (*ProtoSecretKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{0}
}

func (m *ProtoSecretKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoSecretKey.Unmarshal(m, b)
}
func (m *ProtoSecretKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoSecretKey.Marshal(b, m, deterministic)
}
func (m *ProtoSecretKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoSecretKey.Merge(m, src)
}
func (m *ProtoSecretKey) XXX_Size() int {
	return xxx_messageInfo_ProtoSecretKey.Size(m)
}
func (m *ProtoSecretKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoSecretKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoSecretKey proto.InternalMessageInfo

func (m *ProtoSecretKey) GetX() []byte {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *ProtoSecretKey) GetY() [][]byte {
	if m != nil {
		return m.Y
	}
	return nil
}

type ProtoVerificationKey struct {
	G2                   []byte   `protobuf:"bytes,1,opt,name=G2,json=g2,proto3" json:"G2,omitempty"`
	Alpha                []byte   `protobuf:"bytes,2,opt,name=Alpha,json=alpha,proto3" json:"Alpha,omitempty"`
	Beta                 [][]byte `protobuf:"bytes,3,rep,name=Beta,json=beta,proto3" json:"Beta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoVerificationKey) Reset()         { *m = ProtoVerificationKey{} }
func (m *ProtoVerificationKey) String() string { return proto.CompactTextString(m) }
func (*ProtoVerificationKey) ProtoMessage()    {}
func (*ProtoVerificationKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{1}
}

func (m *ProtoVerificationKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoVerificationKey.Unmarshal(m, b)
}
func (m *ProtoVerificationKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoVerificationKey.Marshal(b, m, deterministic)
}
func (m *ProtoVerificationKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoVerificationKey.Merge(m, src)
}
func (m *ProtoVerificationKey) XXX_Size() int {
	return xxx_messageInfo_ProtoVerificationKey.Size(m)
}
func (m *ProtoVerificationKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoVerificationKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoVerificationKey proto.InternalMessageInfo

func (m *ProtoVerificationKey) GetG2() []byte {
	if m != nil {
		return m.G2
	}
	return nil
}

func (m *ProtoVerificationKey) GetAlpha() []byte {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func (m *ProtoVerificationKey) GetBeta() [][]byte {
	if m != nil {
		return m.Beta
	}
	return nil
}

type ProtoSignature struct {
	Sig1                 []byte   `protobuf:"bytes,1,opt,name=sig1,proto3" json:"sig1,omitempty"`
	Sig2                 []byte   `protobuf:"bytes,2,opt,name=sig2,proto3" json:"sig2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoSignature) Reset()         { *m = ProtoSignature{} }
func (m *ProtoSignature) String() string { return proto.CompactTextString(m) }
func (*ProtoSignature) ProtoMessage()    {}
func (*ProtoSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{2}
}

func (m *ProtoSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoSignature.Unmarshal(m, b)
}
func (m *ProtoSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoSignature.Marshal(b, m, deterministic)
}
func (m *ProtoSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoSignature.Merge(m, src)
}
func (m *ProtoSignature) XXX_Size() int {
	return xxx_messageInfo_ProtoSignature.Size(m)
}
func (m *ProtoSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoSignature proto.InternalMessageInfo

func (m *ProtoSignature) GetSig1() []byte {
	if m != nil {
		return m.Sig1
	}
	return nil
}

func (m *ProtoSignature) GetSig2() []byte {
	if m != nil {
		return m.Sig2
	}
	return nil
}

type ProtoBlindedSignature struct {
	Sig1                 []byte                   `protobuf:"bytes,1,opt,name=sig1,proto3" json:"sig1,omitempty"`
	Sig2Tilda            *elgamal.ProtoEncryption `protobuf:"bytes,2,opt,name=sig2Tilda,proto3" json:"sig2Tilda,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProtoBlindedSignature) Reset()         { *m = ProtoBlindedSignature{} }
func (m *ProtoBlindedSignature) String() string { return proto.CompactTextString(m) }
func (*ProtoBlindedSignature) ProtoMessage()    {}
func (*ProtoBlindedSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{3}
}

func (m *ProtoBlindedSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoBlindedSignature.Unmarshal(m, b)
}
func (m *ProtoBlindedSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoBlindedSignature.Marshal(b, m, deterministic)
}
func (m *ProtoBlindedSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoBlindedSignature.Merge(m, src)
}
func (m *ProtoBlindedSignature) XXX_Size() int {
	return xxx_messageInfo_ProtoBlindedSignature.Size(m)
}
func (m *ProtoBlindedSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoBlindedSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoBlindedSignature proto.InternalMessageInfo

func (m *ProtoBlindedSignature) GetSig1() []byte {
	if m != nil {
		return m.Sig1
	}
	return nil
}

func (m *ProtoBlindedSignature) GetSig2Tilda() *elgamal.ProtoEncryption {
	if m != nil {
		return m.Sig2Tilda
	}
	return nil
}

type ProtoSignerProof struct {
	C                    []byte   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	Rr                   []byte   `protobuf:"bytes,2,opt,name=rr,proto3" json:"rr,omitempty"`
	Rk                   [][]byte `protobuf:"bytes,3,rep,name=rk,proto3" json:"rk,omitempty"`
	Rm                   [][]byte `protobuf:"bytes,4,rep,name=rm,proto3" json:"rm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoSignerProof) Reset()         { *m = ProtoSignerProof{} }
func (m *ProtoSignerProof) String() string { return proto.CompactTextString(m) }
func (*ProtoSignerProof) ProtoMessage()    {}
func (*ProtoSignerProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{4}
}

func (m *ProtoSignerProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoSignerProof.Unmarshal(m, b)
}
func (m *ProtoSignerProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoSignerProof.Marshal(b, m, deterministic)
}
func (m *ProtoSignerProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoSignerProof.Merge(m, src)
}
func (m *ProtoSignerProof) XXX_Size() int {
	return xxx_messageInfo_ProtoSignerProof.Size(m)
}
func (m *ProtoSignerProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoSignerProof.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoSignerProof proto.InternalMessageInfo

func (m *ProtoSignerProof) GetC() []byte {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *ProtoSignerProof) GetRr() []byte {
	if m != nil {
		return m.Rr
	}
	return nil
}

func (m *ProtoSignerProof) GetRk() [][]byte {
	if m != nil {
		return m.Rk
	}
	return nil
}

func (m *ProtoSignerProof) GetRm() [][]byte {
	if m != nil {
		return m.Rm
	}
	return nil
}

type ProtoLambda struct {
	Cm                   []byte                     `protobuf:"bytes,1,opt,name=cm,proto3" json:"cm,omitempty"`
	Enc                  []*elgamal.ProtoEncryption `protobuf:"bytes,2,rep,name=enc,proto3" json:"enc,omitempty"`
	Proof                *ProtoSignerProof          `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ProtoLambda) Reset()         { *m = ProtoLambda{} }
func (m *ProtoLambda) String() string { return proto.CompactTextString(m) }
func (*ProtoLambda) ProtoMessage()    {}
func (*ProtoLambda) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{5}
}

func (m *ProtoLambda) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoLambda.Unmarshal(m, b)
}
func (m *ProtoLambda) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoLambda.Marshal(b, m, deterministic)
}
func (m *ProtoLambda) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoLambda.Merge(m, src)
}
func (m *ProtoLambda) XXX_Size() int {
	return xxx_messageInfo_ProtoLambda.Size(m)
}
func (m *ProtoLambda) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoLambda.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoLambda proto.InternalMessageInfo

func (m *ProtoLambda) GetCm() []byte {
	if m != nil {
		return m.Cm
	}
	return nil
}

func (m *ProtoLambda) GetEnc() []*elgamal.ProtoEncryption {
	if m != nil {
		return m.Enc
	}
	return nil
}

func (m *ProtoLambda) GetProof() *ProtoSignerProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type ProtoVerifierProof struct {
	C                    []byte   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	Rm                   [][]byte `protobuf:"bytes,2,rep,name=rm,proto3" json:"rm,omitempty"`
	Rt                   []byte   `protobuf:"bytes,3,opt,name=rt,proto3" json:"rt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoVerifierProof) Reset()         { *m = ProtoVerifierProof{} }
func (m *ProtoVerifierProof) String() string { return proto.CompactTextString(m) }
func (*ProtoVerifierProof) ProtoMessage()    {}
func (*ProtoVerifierProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{6}
}

func (m *ProtoVerifierProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoVerifierProof.Unmarshal(m, b)
}
func (m *ProtoVerifierProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoVerifierProof.Marshal(b, m, deterministic)
}
func (m *ProtoVerifierProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoVerifierProof.Merge(m, src)
}
func (m *ProtoVerifierProof) XXX_Size() int {
	return xxx_messageInfo_ProtoVerifierProof.Size(m)
}
func (m *ProtoVerifierProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoVerifierProof.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoVerifierProof proto.InternalMessageInfo

func (m *ProtoVerifierProof) GetC() []byte {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *ProtoVerifierProof) GetRm() [][]byte {
	if m != nil {
		return m.Rm
	}
	return nil
}

func (m *ProtoVerifierProof) GetRt() []byte {
	if m != nil {
		return m.Rt
	}
	return nil
}

type ProtoTheta struct {
	Kappa                []byte              `protobuf:"bytes,1,opt,name=kappa,proto3" json:"kappa,omitempty"`
	Nu                   []byte              `protobuf:"bytes,2,opt,name=nu,proto3" json:"nu,omitempty"`
	Proof                *ProtoVerifierProof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProtoTheta) Reset()         { *m = ProtoTheta{} }
func (m *ProtoTheta) String() string { return proto.CompactTextString(m) }
func (*ProtoTheta) ProtoMessage()    {}
func (*ProtoTheta) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{7}
}

func (m *ProtoTheta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoTheta.Unmarshal(m, b)
}
func (m *ProtoTheta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoTheta.Marshal(b, m, deterministic)
}
func (m *ProtoTheta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoTheta.Merge(m, src)
}
func (m *ProtoTheta) XXX_Size() int {
	return xxx_messageInfo_ProtoTheta.Size(m)
}
func (m *ProtoTheta) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoTheta.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoTheta proto.InternalMessageInfo

func (m *ProtoTheta) GetKappa() []byte {
	if m != nil {
		return m.Kappa
	}
	return nil
}

func (m *ProtoTheta) GetNu() []byte {
	if m != nil {
		return m.Nu
	}
	return nil
}

func (m *ProtoTheta) GetProof() *ProtoVerifierProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type ProtoParams struct {
	P                    []byte   `protobuf:"bytes,2,opt,name=p,proto3" json:"p,omitempty"`
	G1                   []byte   `protobuf:"bytes,3,opt,name=g1,proto3" json:"g1,omitempty"`
	G2                   []byte   `protobuf:"bytes,4,opt,name=g2,proto3" json:"g2,omitempty"`
	Hs                   [][]byte `protobuf:"bytes,5,rep,name=hs,proto3" json:"hs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoParams) Reset()         { *m = ProtoParams{} }
func (m *ProtoParams) String() string { return proto.CompactTextString(m) }
func (*ProtoParams) ProtoMessage()    {}
func (*ProtoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{8}
}

func (m *ProtoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoParams.Unmarshal(m, b)
}
func (m *ProtoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoParams.Marshal(b, m, deterministic)
}
func (m *ProtoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoParams.Merge(m, src)
}
func (m *ProtoParams) XXX_Size() int {
	return xxx_messageInfo_ProtoParams.Size(m)
}
func (m *ProtoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoParams.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoParams proto.InternalMessageInfo

func (m *ProtoParams) GetP() []byte {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *ProtoParams) GetG1() []byte {
	if m != nil {
		return m.G1
	}
	return nil
}

func (m *ProtoParams) GetG2() []byte {
	if m != nil {
		return m.G2
	}
	return nil
}

func (m *ProtoParams) GetHs() [][]byte {
	if m != nil {
		return m.Hs
	}
	return nil
}

// encapsulates everything required by IAs to issue credential
type ProtoBlindSignMaterials struct {
	Lambda               *ProtoLambda            `protobuf:"bytes,1,opt,name=lambda,proto3" json:"lambda,omitempty"`
	EgPub                *elgamal.ProtoPublicKey `protobuf:"bytes,2,opt,name=egPub,proto3" json:"egPub,omitempty"`
	PubM                 [][]byte                `protobuf:"bytes,3,rep,name=pubM,proto3" json:"pubM,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ProtoBlindSignMaterials) Reset()         { *m = ProtoBlindSignMaterials{} }
func (m *ProtoBlindSignMaterials) String() string { return proto.CompactTextString(m) }
func (*ProtoBlindSignMaterials) ProtoMessage()    {}
func (*ProtoBlindSignMaterials) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{9}
}

func (m *ProtoBlindSignMaterials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoBlindSignMaterials.Unmarshal(m, b)
}
func (m *ProtoBlindSignMaterials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoBlindSignMaterials.Marshal(b, m, deterministic)
}
func (m *ProtoBlindSignMaterials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoBlindSignMaterials.Merge(m, src)
}
func (m *ProtoBlindSignMaterials) XXX_Size() int {
	return xxx_messageInfo_ProtoBlindSignMaterials.Size(m)
}
func (m *ProtoBlindSignMaterials) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoBlindSignMaterials.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoBlindSignMaterials proto.InternalMessageInfo

func (m *ProtoBlindSignMaterials) GetLambda() *ProtoLambda {
	if m != nil {
		return m.Lambda
	}
	return nil
}

func (m *ProtoBlindSignMaterials) GetEgPub() *elgamal.ProtoPublicKey {
	if m != nil {
		return m.EgPub
	}
	return nil
}

func (m *ProtoBlindSignMaterials) GetPubM() [][]byte {
	if m != nil {
		return m.PubM
	}
	return nil
}

// encapsulates everything required by verifiers to verify credentials
type ProtoBlindVerifyMaterials struct {
	Sig                  *ProtoSignature `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	Theta                *ProtoTheta     `protobuf:"bytes,2,opt,name=theta,proto3" json:"theta,omitempty"`
	PubM                 [][]byte        `protobuf:"bytes,3,rep,name=pubM,proto3" json:"pubM,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProtoBlindVerifyMaterials) Reset()         { *m = ProtoBlindVerifyMaterials{} }
func (m *ProtoBlindVerifyMaterials) String() string { return proto.CompactTextString(m) }
func (*ProtoBlindVerifyMaterials) ProtoMessage()    {}
func (*ProtoBlindVerifyMaterials) Descriptor() ([]byte, []int) {
	return fileDescriptor_69ade651a23d8dbd, []int{10}
}

func (m *ProtoBlindVerifyMaterials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoBlindVerifyMaterials.Unmarshal(m, b)
}
func (m *ProtoBlindVerifyMaterials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoBlindVerifyMaterials.Marshal(b, m, deterministic)
}
func (m *ProtoBlindVerifyMaterials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoBlindVerifyMaterials.Merge(m, src)
}
func (m *ProtoBlindVerifyMaterials) XXX_Size() int {
	return xxx_messageInfo_ProtoBlindVerifyMaterials.Size(m)
}
func (m *ProtoBlindVerifyMaterials) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoBlindVerifyMaterials.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoBlindVerifyMaterials proto.InternalMessageInfo

func (m *ProtoBlindVerifyMaterials) GetSig() *ProtoSignature {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *ProtoBlindVerifyMaterials) GetTheta() *ProtoTheta {
	if m != nil {
		return m.Theta
	}
	return nil
}

func (m *ProtoBlindVerifyMaterials) GetPubM() [][]byte {
	if m != nil {
		return m.PubM
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoSecretKey)(nil), "coconut.ProtoSecretKey")
	proto.RegisterType((*ProtoVerificationKey)(nil), "coconut.ProtoVerificationKey")
	proto.RegisterType((*ProtoSignature)(nil), "coconut.ProtoSignature")
	proto.RegisterType((*ProtoBlindedSignature)(nil), "coconut.ProtoBlindedSignature")
	proto.RegisterType((*ProtoSignerProof)(nil), "coconut.ProtoSignerProof")
	proto.RegisterType((*ProtoLambda)(nil), "coconut.ProtoLambda")
	proto.RegisterType((*ProtoVerifierProof)(nil), "coconut.ProtoVerifierProof")
	proto.RegisterType((*ProtoTheta)(nil), "coconut.ProtoTheta")
	proto.RegisterType((*ProtoParams)(nil), "coconut.ProtoParams")
	proto.RegisterType((*ProtoBlindSignMaterials)(nil), "coconut.ProtoBlindSignMaterials")
	proto.RegisterType((*ProtoBlindVerifyMaterials)(nil), "coconut.ProtoBlindVerifyMaterials")
}

func init() {
	proto.RegisterFile("crypto/coconut/scheme/proto/types.proto", fileDescriptor_69ade651a23d8dbd)
}

var fileDescriptor_69ade651a23d8dbd = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x51, 0x6f, 0x94, 0x4c,
	0x14, 0x0d, 0xec, 0xd2, 0xef, 0x73, 0xda, 0x34, 0xcd, 0x58, 0xb3, 0x54, 0x5f, 0x36, 0xbc, 0x68,
	0xcd, 0x0a, 0x59, 0x4c, 0x8c, 0xaf, 0x6e, 0x62, 0x4c, 0xac, 0x4d, 0x10, 0x1b, 0xa3, 0xbe, 0x0d,
	0xb3, 0x53, 0x98, 0x2c, 0x03, 0x64, 0x18, 0x12, 0xf1, 0xdd, 0x17, 0x7f, 0xb5, 0xb9, 0x97, 0xa1,
	0x2b, 0xb6, 0xf6, 0x69, 0xef, 0x61, 0xcf, 0xbd, 0x9c, 0x7b, 0xce, 0x0d, 0xe4, 0x29, 0xd7, 0x7d,
	0x63, 0xea, 0x88, 0xd7, 0xbc, 0xae, 0x3a, 0x13, 0xb5, 0xbc, 0x10, 0x4a, 0x44, 0x8d, 0xae, 0x4d,
	0x1d, 0x99, 0xbe, 0x11, 0x6d, 0x88, 0x35, 0xfd, 0xcf, 0x32, 0x1e, 0x2f, 0x6d, 0x87, 0x28, 0x73,
	0xa6, 0x58, 0x79, 0x9b, 0x1a, 0xac, 0xc8, 0x71, 0x02, 0xc5, 0x27, 0xc1, 0xb5, 0x30, 0x17, 0xa2,
	0xa7, 0x47, 0xc4, 0xf9, 0xe2, 0x3b, 0x4b, 0xe7, 0xd9, 0x51, 0xea, 0x7c, 0x07, 0xf4, 0xd5, 0x77,
	0x97, 0x33, 0x40, 0x7d, 0x90, 0x90, 0x53, 0x64, 0x7f, 0x16, 0x5a, 0x5e, 0x4b, 0xce, 0x8c, 0xac,
	0x2b, 0xe8, 0x39, 0x26, 0xee, 0xbb, 0xd8, 0x36, 0xb9, 0x79, 0x4c, 0x4f, 0x89, 0xf7, 0xa6, 0x6c,
	0x0a, 0xe6, 0xbb, 0xf8, 0xc8, 0x63, 0x00, 0x28, 0x25, 0xf3, 0x8d, 0x30, 0xcc, 0x9f, 0xe1, 0xb8,
	0x79, 0x26, 0x0c, 0x0b, 0x5e, 0x8f, 0xef, 0x97, 0x79, 0xc5, 0x4c, 0xa7, 0x05, 0xb0, 0x5a, 0x99,
	0xaf, 0xed, 0x34, 0xac, 0xed, 0xb3, 0xd8, 0x8e, 0xc3, 0x3a, 0xe0, 0xe4, 0x11, 0x76, 0x6e, 0x4a,
	0x59, 0x6d, 0xc5, 0xf6, 0xfe, 0x01, 0xaf, 0xc8, 0x03, 0x68, 0xba, 0x92, 0xe5, 0x76, 0x10, 0x75,
	0x18, 0xfb, 0xa1, 0x75, 0x25, 0xc4, 0x31, 0x6f, 0x2b, 0xf4, 0x4a, 0xd6, 0x55, 0xba, 0xa7, 0x06,
	0x09, 0x39, 0xb9, 0x91, 0x27, 0x74, 0xa2, 0xeb, 0xfa, 0x1a, 0x2c, 0xe1, 0xa3, 0x41, 0x1c, 0x56,
	0xd7, 0xda, 0x0a, 0x73, 0xb5, 0x46, 0xbc, 0xb3, 0x2b, 0xba, 0x7a, 0x87, 0x58, 0xf9, 0x73, 0x8b,
	0x55, 0xf0, 0x83, 0x1c, 0xe2, 0xc4, 0x0f, 0x4c, 0x65, 0x5b, 0x06, 0x7f, 0x73, 0x35, 0x3a, 0xc7,
	0x15, 0x7d, 0x4e, 0x66, 0xa2, 0xe2, 0xe8, 0xf8, 0x7d, 0x12, 0x81, 0x44, 0x23, 0xe2, 0x35, 0xa0,
	0xc8, 0x9f, 0xe1, 0x42, 0x67, 0xa1, 0x8d, 0x3d, 0xfc, 0x5b, 0x72, 0x3a, 0xf0, 0x82, 0x0d, 0xa1,
	0x7f, 0xc4, 0xf7, 0xef, 0x7d, 0x94, 0x4d, 0xdc, 0xd5, 0x0a, 0xb1, 0xc1, 0x37, 0x00, 0x36, 0x81,
	0x20, 0x04, 0x67, 0x5c, 0x15, 0xc2, 0x30, 0x08, 0x7a, 0xc7, 0x9a, 0x86, 0xd9, 0xfe, 0x01, 0x40,
	0x4f, 0xd5, 0x8d, 0x9e, 0x54, 0x1d, 0x5d, 0x4f, 0x85, 0x3e, 0x99, 0x0a, 0x9d, 0xa8, 0x19, 0xa5,
	0x7e, 0xb4, 0x36, 0x25, 0x4c, 0x33, 0xd5, 0x82, 0xc6, 0xc6, 0x0e, 0x74, 0x1a, 0x98, 0x9f, 0xaf,
	0x47, 0x4d, 0xf9, 0x1a, 0x71, 0xec, 0xcf, 0x6f, 0xce, 0xef, 0x98, 0xb8, 0x45, 0xeb, 0x7b, 0xc3,
	0x0e, 0x45, 0xfb, 0x7e, 0xfe, 0xbf, 0x73, 0xe2, 0x06, 0xbf, 0x1c, 0xb2, 0xd8, 0x5f, 0x0c, 0xd8,
	0x73, 0xc9, 0x8c, 0xd0, 0x92, 0x95, 0x2d, 0x5d, 0x91, 0x83, 0x12, 0x03, 0xc1, 0x45, 0x0e, 0xe3,
	0xd3, 0xa9, 0xc4, 0x21, 0xac, 0xd4, 0x72, 0xe8, 0x0b, 0xe2, 0x89, 0x3c, 0xe9, 0x32, 0x7b, 0x49,
	0x8b, 0x69, 0x4c, 0x49, 0x97, 0x95, 0x92, 0x5f, 0x88, 0x3e, 0x1d, 0x58, 0x70, 0x90, 0x4d, 0x97,
	0x5d, 0x8e, 0x77, 0x0f, 0x75, 0xf0, 0xd3, 0x21, 0x67, 0x7b, 0x31, 0x68, 0x41, 0xbf, 0x97, 0x73,
	0x4e, 0x66, 0xad, 0xcc, 0xad, 0x96, 0xc5, 0xed, 0x5c, 0xf1, 0xd0, 0x53, 0xe0, 0xd0, 0x73, 0xe2,
	0x19, 0x88, 0xc2, 0x6a, 0x79, 0x38, 0x25, 0x63, 0x4a, 0xe9, 0xc0, 0xb8, 0x4b, 0xc7, 0x26, 0xfc,
	0xb6, 0xca, 0xa5, 0x29, 0xba, 0x2c, 0xe4, 0xb5, 0x8a, 0xaa, 0x5e, 0x19, 0xc1, 0x0b, 0xf8, 0x8d,
	0xee, 0xfc, 0xd8, 0x64, 0x07, 0xf8, 0xd9, 0x78, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd7,
	0x8f, 0x41, 0x8c, 0x04, 0x00, 0x00,
}
