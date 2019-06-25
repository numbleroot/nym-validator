// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tendermint/nymabci/transaction/proto/types.proto

package transaction

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	scheme "github.com/nymtech/nym/crypto/coconut/scheme"
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

type NewAccountRequest struct {
	// Public Key of the user used to derive account address and validate signature
	Address []byte `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	// represents some optional credential from an IP if required
	Credential []byte `protobuf:"bytes,2,opt,name=Credential,json=credential,proto3" json:"Credential,omitempty"`
	// Signature on request to confirm its validity + asserts knowledge of private key
	Sig                  []byte   `protobuf:"bytes,3,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAccountRequest) Reset()         { *m = NewAccountRequest{} }
func (m *NewAccountRequest) String() string { return proto.CompactTextString(m) }
func (*NewAccountRequest) ProtoMessage()    {}
func (*NewAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{0}
}

func (m *NewAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountRequest.Unmarshal(m, b)
}
func (m *NewAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountRequest.Marshal(b, m, deterministic)
}
func (m *NewAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountRequest.Merge(m, src)
}
func (m *NewAccountRequest) XXX_Size() int {
	return xxx_messageInfo_NewAccountRequest.Size(m)
}
func (m *NewAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountRequest proto.InternalMessageInfo

func (m *NewAccountRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NewAccountRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *NewAccountRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// DEBUG
type AccountTransferRequest struct {
	// Used to validate signature + determine source address
	SourceAddress []byte `protobuf:"bytes,1,opt,name=SourceAddress,json=sourceAddress,proto3" json:"SourceAddress,omitempty"`
	// Used to determine target address
	TargetAddress []byte `protobuf:"bytes,2,opt,name=TargetAddress,json=targetAddress,proto3" json:"TargetAddress,omitempty"`
	// Amount to be transferred
	Amount uint64 `protobuf:"varint,3,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// While this function will only be available in debug and hence a nonce is really not needed,
	// I figured I should include it anyway as it's a good practice + will need to figure out a proper
	// nonce system anyway.
	Nonce []byte `protobuf:"bytes,4,opt,name=Nonce,json=nonce,proto3" json:"Nonce,omitempty"`
	// Signature on request to confirm its validitiy
	Sig                  []byte   `protobuf:"bytes,5,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountTransferRequest) Reset()         { *m = AccountTransferRequest{} }
func (m *AccountTransferRequest) String() string { return proto.CompactTextString(m) }
func (*AccountTransferRequest) ProtoMessage()    {}
func (*AccountTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{1}
}

func (m *AccountTransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountTransferRequest.Unmarshal(m, b)
}
func (m *AccountTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountTransferRequest.Marshal(b, m, deterministic)
}
func (m *AccountTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountTransferRequest.Merge(m, src)
}
func (m *AccountTransferRequest) XXX_Size() int {
	return xxx_messageInfo_AccountTransferRequest.Size(m)
}
func (m *AccountTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountTransferRequest proto.InternalMessageInfo

func (m *AccountTransferRequest) GetSourceAddress() []byte {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *AccountTransferRequest) GetTargetAddress() []byte {
	if m != nil {
		return m.TargetAddress
	}
	return nil
}

func (m *AccountTransferRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AccountTransferRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *AccountTransferRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type DepositCoconutCredentialRequest struct {
	// Includes the actual credential, any public attributes
	// and theta (that include, among other things, NIZK proofs required to verify the credential)
	CryptoMaterials *scheme.ProtoTumblerBlindVerifyMaterials `protobuf:"bytes,1,opt,name=CryptoMaterials,json=cryptoMaterials,proto3" json:"CryptoMaterials,omitempty"`
	// Value of the credential. While it is included in a BIG form in CryptoMaterials, it's easier to operate on it
	// when it's an int. We can't send it as an uint64, as milagro requires a normal int argument to construct a BIG num.
	Value int64 `protobuf:"varint,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	// Address of the service provider to which the proof is bound and whose account balance will be increased
	ProviderAddress      []byte   `protobuf:"bytes,3,opt,name=ProviderAddress,json=providerAddress,proto3" json:"ProviderAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DepositCoconutCredentialRequest) Reset()         { *m = DepositCoconutCredentialRequest{} }
func (m *DepositCoconutCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*DepositCoconutCredentialRequest) ProtoMessage()    {}
func (*DepositCoconutCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{2}
}

func (m *DepositCoconutCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepositCoconutCredentialRequest.Unmarshal(m, b)
}
func (m *DepositCoconutCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepositCoconutCredentialRequest.Marshal(b, m, deterministic)
}
func (m *DepositCoconutCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositCoconutCredentialRequest.Merge(m, src)
}
func (m *DepositCoconutCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_DepositCoconutCredentialRequest.Size(m)
}
func (m *DepositCoconutCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositCoconutCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DepositCoconutCredentialRequest proto.InternalMessageInfo

func (m *DepositCoconutCredentialRequest) GetCryptoMaterials() *scheme.ProtoTumblerBlindVerifyMaterials {
	if m != nil {
		return m.CryptoMaterials
	}
	return nil
}

func (m *DepositCoconutCredentialRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *DepositCoconutCredentialRequest) GetProviderAddress() []byte {
	if m != nil {
		return m.ProviderAddress
	}
	return nil
}

type TransferToPipeAccountNotification struct {
	// Used to identify the particular watcher and to verify signature
	WatcherPublicKey []byte `protobuf:"bytes,1,opt,name=WatcherPublicKey,json=watcherPublicKey,proto3" json:"WatcherPublicKey,omitempty"`
	// Ethereum address of the client
	ClientAddress []byte `protobuf:"bytes,2,opt,name=ClientAddress,json=clientAddress,proto3" json:"ClientAddress,omitempty"`
	// While right now it's completely unrequired as there is only a single pipe account, it might be useful
	// to have this information in the future if we decided to monitor multiple chains or have multiple pipe accounts
	// for example on epoch changes.
	PipeAccountAddress []byte `protobuf:"bytes,3,opt,name=PipeAccountAddress,json=pipeAccountAddress,proto3" json:"PipeAccountAddress,omitempty"`
	// Amount transferred by the client to the pipe account.
	Amount uint64 `protobuf:"varint,4,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Hash of the transaction in which the transfer occured.
	// Used to distinguish from multiple transfers the client might have done.
	TxHash []byte `protobuf:"bytes,5,opt,name=TxHash,json=txHash,proto3" json:"TxHash,omitempty"`
	// Signature on the entire message done with the watcher's key.
	Sig                  []byte   `protobuf:"bytes,6,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferToPipeAccountNotification) Reset()         { *m = TransferToPipeAccountNotification{} }
func (m *TransferToPipeAccountNotification) String() string { return proto.CompactTextString(m) }
func (*TransferToPipeAccountNotification) ProtoMessage()    {}
func (*TransferToPipeAccountNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{3}
}

func (m *TransferToPipeAccountNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferToPipeAccountNotification.Unmarshal(m, b)
}
func (m *TransferToPipeAccountNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferToPipeAccountNotification.Marshal(b, m, deterministic)
}
func (m *TransferToPipeAccountNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferToPipeAccountNotification.Merge(m, src)
}
func (m *TransferToPipeAccountNotification) XXX_Size() int {
	return xxx_messageInfo_TransferToPipeAccountNotification.Size(m)
}
func (m *TransferToPipeAccountNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferToPipeAccountNotification.DiscardUnknown(m)
}

var xxx_messageInfo_TransferToPipeAccountNotification proto.InternalMessageInfo

func (m *TransferToPipeAccountNotification) GetWatcherPublicKey() []byte {
	if m != nil {
		return m.WatcherPublicKey
	}
	return nil
}

func (m *TransferToPipeAccountNotification) GetClientAddress() []byte {
	if m != nil {
		return m.ClientAddress
	}
	return nil
}

func (m *TransferToPipeAccountNotification) GetPipeAccountAddress() []byte {
	if m != nil {
		return m.PipeAccountAddress
	}
	return nil
}

func (m *TransferToPipeAccountNotification) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferToPipeAccountNotification) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TransferToPipeAccountNotification) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type CredentialRequest struct {
	// Ethereum address of the client
	ClientAddress []byte `protobuf:"bytes,1,opt,name=ClientAddress,json=clientAddress,proto3" json:"ClientAddress,omitempty"`
	// While right now it's completely unrequired as there is only a single pipe account, it might be useful
	// to have this information in the future if we decided to monitor multiple chains or have multiple pipe accounts
	// for example on epoch changes.
	PipeAccountAddress []byte `protobuf:"bytes,2,opt,name=PipeAccountAddress,json=pipeAccountAddress,proto3" json:"PipeAccountAddress,omitempty"`
	// All the cryptographic materials required by issuers to perform a blind sign
	CryptoMaterials *scheme.ProtoBlindSignMaterials `protobuf:"bytes,3,opt,name=CryptoMaterials,json=cryptoMaterials,proto3" json:"CryptoMaterials,omitempty"`
	// Value of the credential. While it is included in a BIG form in CryptoMaterials, it's easier to operate on it
	// when it's an int. We can't send it as an uint64, as milagro requires a normal int argument to construct a BIG num.
	Value int64 `protobuf:"varint,4,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	// Required to prevent replay attacks.
	Nonce []byte `protobuf:"bytes,5,opt,name=Nonce,json=nonce,proto3" json:"Nonce,omitempty"`
	// Signature on entire request with client's ethereum key (so that client's address could be used to verify it)
	Sig                  []byte   `protobuf:"bytes,6,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialRequest) Reset()         { *m = CredentialRequest{} }
func (m *CredentialRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialRequest) ProtoMessage()    {}
func (*CredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{4}
}

func (m *CredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialRequest.Unmarshal(m, b)
}
func (m *CredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialRequest.Marshal(b, m, deterministic)
}
func (m *CredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialRequest.Merge(m, src)
}
func (m *CredentialRequest) XXX_Size() int {
	return xxx_messageInfo_CredentialRequest.Size(m)
}
func (m *CredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialRequest proto.InternalMessageInfo

func (m *CredentialRequest) GetClientAddress() []byte {
	if m != nil {
		return m.ClientAddress
	}
	return nil
}

func (m *CredentialRequest) GetPipeAccountAddress() []byte {
	if m != nil {
		return m.PipeAccountAddress
	}
	return nil
}

func (m *CredentialRequest) GetCryptoMaterials() *scheme.ProtoBlindSignMaterials {
	if m != nil {
		return m.CryptoMaterials
	}
	return nil
}

func (m *CredentialRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CredentialRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *CredentialRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type CredentialVerificationNotification struct {
	// Used to identify the particular verifier and to verify signature
	VerifierPublicKey []byte `protobuf:"bytes,1,opt,name=VerifierPublicKey,json=verifierPublicKey,proto3" json:"VerifierPublicKey,omitempty"`
	// Address of the provider who sent the original request
	ProviderAddress []byte `protobuf:"bytes,2,opt,name=ProviderAddress,json=providerAddress,proto3" json:"ProviderAddress,omitempty"`
	// Value of the credential
	Value int64 `protobuf:"varint,3,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	// Zeta associated with the credential to unique identify it
	Zeta []byte `protobuf:"bytes,4,opt,name=Zeta,json=zeta,proto3" json:"Zeta,omitempty"`
	// Was the credential valid or not
	CredentialValidity bool `protobuf:"varint,5,opt,name=CredentialValidity,json=credentialValidity,proto3" json:"CredentialValidity,omitempty"`
	// Signature on the entire message done with the verifier's key.
	Sig                  []byte   `protobuf:"bytes,6,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialVerificationNotification) Reset()         { *m = CredentialVerificationNotification{} }
func (m *CredentialVerificationNotification) String() string { return proto.CompactTextString(m) }
func (*CredentialVerificationNotification) ProtoMessage()    {}
func (*CredentialVerificationNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{5}
}

func (m *CredentialVerificationNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialVerificationNotification.Unmarshal(m, b)
}
func (m *CredentialVerificationNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialVerificationNotification.Marshal(b, m, deterministic)
}
func (m *CredentialVerificationNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialVerificationNotification.Merge(m, src)
}
func (m *CredentialVerificationNotification) XXX_Size() int {
	return xxx_messageInfo_CredentialVerificationNotification.Size(m)
}
func (m *CredentialVerificationNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialVerificationNotification.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialVerificationNotification proto.InternalMessageInfo

func (m *CredentialVerificationNotification) GetVerifierPublicKey() []byte {
	if m != nil {
		return m.VerifierPublicKey
	}
	return nil
}

func (m *CredentialVerificationNotification) GetProviderAddress() []byte {
	if m != nil {
		return m.ProviderAddress
	}
	return nil
}

func (m *CredentialVerificationNotification) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CredentialVerificationNotification) GetZeta() []byte {
	if m != nil {
		return m.Zeta
	}
	return nil
}

func (m *CredentialVerificationNotification) GetCredentialValidity() bool {
	if m != nil {
		return m.CredentialValidity
	}
	return false
}

func (m *CredentialVerificationNotification) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type TokenRedemptionRequest struct {
	// Ethereum address of the provider wishing to move funds back to ERC20
	ProviderAddress []byte `protobuf:"bytes,1,opt,name=ProviderAddress,json=providerAddress,proto3" json:"ProviderAddress,omitempty"`
	// Amount to move back into ERC20 tokens. Provider needs to have at least that much available.
	Amount uint64 `protobuf:"varint,2,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Required to prevent replay attacks. TODO: if managing this turns out to be too difficult, for now just
	// move ALL funds from the account back to ERC20
	Nonce []byte `protobuf:"bytes,3,opt,name=Nonce,json=nonce,proto3" json:"Nonce,omitempty"`
	// Signature on entire request to confirm its validity
	Sig                  []byte   `protobuf:"bytes,4,opt,name=Sig,json=sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRedemptionRequest) Reset()         { *m = TokenRedemptionRequest{} }
func (m *TokenRedemptionRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRedemptionRequest) ProtoMessage()    {}
func (*TokenRedemptionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffb862c5130efc92, []int{6}
}

func (m *TokenRedemptionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRedemptionRequest.Unmarshal(m, b)
}
func (m *TokenRedemptionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRedemptionRequest.Marshal(b, m, deterministic)
}
func (m *TokenRedemptionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRedemptionRequest.Merge(m, src)
}
func (m *TokenRedemptionRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRedemptionRequest.Size(m)
}
func (m *TokenRedemptionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRedemptionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRedemptionRequest proto.InternalMessageInfo

func (m *TokenRedemptionRequest) GetProviderAddress() []byte {
	if m != nil {
		return m.ProviderAddress
	}
	return nil
}

func (m *TokenRedemptionRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TokenRedemptionRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TokenRedemptionRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*NewAccountRequest)(nil), "transaction.NewAccountRequest")
	proto.RegisterType((*AccountTransferRequest)(nil), "transaction.AccountTransferRequest")
	proto.RegisterType((*DepositCoconutCredentialRequest)(nil), "transaction.DepositCoconutCredentialRequest")
	proto.RegisterType((*TransferToPipeAccountNotification)(nil), "transaction.TransferToPipeAccountNotification")
	proto.RegisterType((*CredentialRequest)(nil), "transaction.CredentialRequest")
	proto.RegisterType((*CredentialVerificationNotification)(nil), "transaction.CredentialVerificationNotification")
	proto.RegisterType((*TokenRedemptionRequest)(nil), "transaction.TokenRedemptionRequest")
}

func init() {
	proto.RegisterFile("tendermint/nymabci/transaction/proto/types.proto", fileDescriptor_ffb862c5130efc92)
}

var fileDescriptor_ffb862c5130efc92 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6e, 0xd3, 0x3c,
	0x14, 0x57, 0x9a, 0xb4, 0xfb, 0xe4, 0x7d, 0x53, 0x57, 0x0b, 0x55, 0xd5, 0x2e, 0x60, 0x54, 0x48,
	0x0c, 0x84, 0x1a, 0x04, 0x42, 0x5c, 0x6f, 0xe5, 0x02, 0x81, 0x98, 0xaa, 0x34, 0x1a, 0xd2, 0x6e,
	0x90, 0xeb, 0x9c, 0xb5, 0x16, 0x89, 0x1d, 0x1c, 0x67, 0x23, 0x5c, 0xf2, 0x2c, 0x3c, 0x03, 0xaf,
	0x04, 0xf7, 0xbc, 0x00, 0xb2, 0x93, 0x34, 0x49, 0x13, 0xa6, 0x5d, 0xb5, 0xe7, 0x1c, 0xdb, 0xe7,
	0xf7, 0xe7, 0xe4, 0xa0, 0xe7, 0x0a, 0x78, 0x00, 0x32, 0x62, 0x5c, 0xb9, 0x3c, 0x8b, 0xc8, 0x8a,
	0x32, 0x57, 0x49, 0xc2, 0x13, 0x42, 0x15, 0x13, 0xdc, 0x8d, 0xa5, 0x50, 0xc2, 0x55, 0x59, 0x0c,
	0xc9, 0xcc, 0xfc, 0xc7, 0xfb, 0xb5, 0xf2, 0xd1, 0x63, 0x2a, 0xb3, 0x58, 0x09, 0x97, 0x0a, 0x2a,
	0x78, 0xaa, 0xdc, 0x84, 0x6e, 0x20, 0x82, 0xf6, 0xad, 0xa3, 0xd9, 0xad, 0x07, 0xd3, 0x68, 0x15,
	0x82, 0xac, 0x9d, 0x9f, 0x7e, 0x42, 0xa3, 0x73, 0xb8, 0x39, 0xa5, 0x54, 0xa4, 0x5c, 0x79, 0xf0,
	0x25, 0x85, 0x44, 0xe1, 0x09, 0xda, 0x3b, 0x0d, 0x02, 0x09, 0x49, 0x32, 0xb1, 0x8e, 0xad, 0x93,
	0xff, 0xbd, 0x3d, 0x92, 0x87, 0xf8, 0x3e, 0x42, 0x73, 0x09, 0x01, 0x70, 0xc5, 0x48, 0x38, 0xe9,
	0x99, 0x22, 0xa2, 0xdb, 0x0c, 0x3e, 0x44, 0xf6, 0x92, 0xad, 0x27, 0xb6, 0x29, 0xd8, 0x09, 0x5b,
	0x4f, 0x7f, 0x58, 0x68, 0x5c, 0x3c, 0xef, 0x6b, 0x42, 0x57, 0x20, 0xcb, 0x36, 0x8f, 0xd0, 0xc1,
	0x52, 0xa4, 0x92, 0x42, 0xb3, 0xd9, 0x41, 0x52, 0x4f, 0xea, 0x53, 0x3e, 0x91, 0x6b, 0x50, 0xe5,
	0xa9, 0xbc, 0xeb, 0x81, 0xaa, 0x27, 0xf1, 0x18, 0x0d, 0x4e, 0x23, 0xdd, 0xc4, 0xf4, 0x76, 0xbc,
	0x01, 0x31, 0x11, 0xbe, 0x87, 0xfa, 0xe7, 0x82, 0x53, 0x98, 0x38, 0xe6, 0x56, 0x9f, 0xeb, 0xa0,
	0x84, 0xd9, 0xaf, 0x60, 0xfe, 0xb4, 0xd0, 0x83, 0x37, 0x10, 0x8b, 0x84, 0xa9, 0x79, 0x2e, 0x5d,
	0xc5, 0xb3, 0xc4, 0xbb, 0x44, 0xc3, 0xb9, 0x51, 0xf7, 0x03, 0x51, 0x20, 0x19, 0x09, 0x73, 0xc4,
	0xfb, 0x2f, 0x9e, 0xcc, 0x0a, 0xb9, 0x67, 0x0b, 0x2d, 0xaa, 0x9f, 0xcb, 0x7c, 0x16, 0x32, 0x1e,
	0x5c, 0x80, 0x64, 0x57, 0xd9, 0xf6, 0x82, 0x37, 0xa4, 0xcd, 0x17, 0x34, 0xc0, 0x0b, 0x12, 0xa6,
	0x60, 0x68, 0xd9, 0x5e, 0xff, 0x5a, 0x07, 0xf8, 0x04, 0x0d, 0x17, 0x52, 0x5c, 0xb3, 0x00, 0x64,
	0x49, 0x3b, 0xd7, 0x74, 0x18, 0x37, 0xd3, 0xd3, 0x5f, 0x16, 0x7a, 0x58, 0x0a, 0xeb, 0x8b, 0x05,
	0x8b, 0xa1, 0x50, 0xfb, 0x5c, 0x28, 0x76, 0xc5, 0x28, 0xd1, 0xf3, 0x83, 0x9f, 0xa2, 0xc3, 0x8f,
	0x44, 0xd1, 0x0d, 0xc8, 0x45, 0xba, 0x0a, 0x19, 0x7d, 0x0f, 0x59, 0xa1, 0xf6, 0xe1, 0xcd, 0x4e,
	0x5e, 0x0b, 0x3e, 0x0f, 0x19, 0xf0, 0x5d, 0xc1, 0x69, 0x3d, 0x89, 0x67, 0x08, 0xd7, 0x9a, 0x35,
	0x41, 0xe2, 0xb8, 0x55, 0xa9, 0x19, 0xe4, 0x34, 0x0c, 0x1a, 0xa3, 0x81, 0xff, 0xf5, 0x2d, 0x49,
	0x36, 0x85, 0x1b, 0x03, 0x65, 0xa2, 0xd2, 0xa2, 0x41, 0x65, 0xd1, 0x1f, 0x0b, 0x8d, 0xda, 0xa6,
	0xb4, 0xd0, 0x5a, 0x77, 0x47, 0xdb, 0xfb, 0x27, 0xda, 0x77, 0x6d, 0xab, 0x6d, 0x63, 0xf5, 0x71,
	0xd3, 0x6a, 0xe3, 0xf1, 0x92, 0xad, 0xf9, 0x5d, 0x1c, 0x76, 0xea, 0x0e, 0x6f, 0x07, 0xb3, 0xdf,
	0x31, 0x98, 0x35, 0xd6, 0xbf, 0x2d, 0x34, 0xad, 0x58, 0x9b, 0x71, 0x2a, 0x4c, 0x6d, 0x18, 0xfc,
	0x0c, 0x8d, 0xf2, 0x5a, 0xdb, 0xe1, 0xd1, 0xf5, 0x6e, 0xa1, 0x6b, 0xbc, 0x7a, 0x9d, 0xe3, 0x55,
	0x81, 0xb7, 0xeb, 0xe0, 0x31, 0x72, 0x2e, 0x41, 0x91, 0xe2, 0xa3, 0x72, 0xbe, 0x81, 0x22, 0x5a,
	0xe2, 0x1a, 0x4e, 0x12, 0xb2, 0x80, 0xa9, 0xcc, 0xb0, 0xfb, 0xcf, 0xc3, 0xb4, 0x55, 0xe9, 0xa0,
	0xfa, 0xdd, 0x42, 0x63, 0x5f, 0x7c, 0x06, 0xee, 0x41, 0x00, 0x51, 0xac, 0x79, 0x95, 0x2e, 0x77,
	0x00, 0xb6, 0xba, 0x01, 0x57, 0x73, 0xd6, 0xeb, 0x5e, 0x04, 0x76, 0x87, 0xde, 0xce, 0x16, 0xc4,
	0xd9, 0xeb, 0xcb, 0x57, 0x6b, 0xa6, 0x36, 0xe9, 0x6a, 0x46, 0x45, 0xa4, 0x57, 0xb5, 0x02, 0xba,
	0xd1, 0xbf, 0xee, 0xed, 0x1b, 0x7c, 0x35, 0x30, 0x0b, 0xf5, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3e, 0x0d, 0x2f, 0x05, 0xea, 0x05, 0x00, 0x00,
}
