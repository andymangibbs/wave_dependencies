// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/sigpb/sigpb.proto

package sigpb // import "github.com/google/trillian/crypto/sigpb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HashAlgorithm defines the approved methods for object hashing.
//
// Supported hash algorithms. The numbering space is the same as for TLS,
// given in RFC 5246 s7.4.1.4.1 and at:
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-18
type DigitallySigned_HashAlgorithm int32

const (
	// No hash algorithm is used.
	DigitallySigned_NONE DigitallySigned_HashAlgorithm = 0
	// SHA256 is used.
	DigitallySigned_SHA256 DigitallySigned_HashAlgorithm = 4
)

var DigitallySigned_HashAlgorithm_name = map[int32]string{
	0: "NONE",
	4: "SHA256",
}
var DigitallySigned_HashAlgorithm_value = map[string]int32{
	"NONE":   0,
	"SHA256": 4,
}

func (x DigitallySigned_HashAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_HashAlgorithm_name, int32(x))
}
func (DigitallySigned_HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sigpb_62a21f00d4a205f5, []int{0, 0}
}

// SignatureAlgorithm defines the algorithm used to sign the object.
//
// Supported signature algorithms. The numbering space is the same as for TLS,
// given in RFC 5246 s7.4.1.4.1 and at:
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-16
type DigitallySigned_SignatureAlgorithm int32

const (
	// Anonymous signature scheme.
	DigitallySigned_ANONYMOUS DigitallySigned_SignatureAlgorithm = 0
	// RSA signature scheme.
	DigitallySigned_RSA DigitallySigned_SignatureAlgorithm = 1
	// ECDSA signature scheme.
	DigitallySigned_ECDSA DigitallySigned_SignatureAlgorithm = 3
)

var DigitallySigned_SignatureAlgorithm_name = map[int32]string{
	0: "ANONYMOUS",
	1: "RSA",
	3: "ECDSA",
}
var DigitallySigned_SignatureAlgorithm_value = map[string]int32{
	"ANONYMOUS": 0,
	"RSA":       1,
	"ECDSA":     3,
}

func (x DigitallySigned_SignatureAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_SignatureAlgorithm_name, int32(x))
}
func (DigitallySigned_SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sigpb_62a21f00d4a205f5, []int{0, 1}
}

// Protocol buffer encoding of the TLS DigitallySigned type, from RFC 5246 ??4.7.
type DigitallySigned struct {
	// hash_algorithm contains the hash algorithm used.
	HashAlgorithm DigitallySigned_HashAlgorithm `protobuf:"varint,1,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// sig_algorithm contains the signing algorithm used.
	SignatureAlgorithm DigitallySigned_SignatureAlgorithm `protobuf:"varint,2,opt,name=signature_algorithm,json=signatureAlgorithm,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// signature contains the object signature.
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DigitallySigned) Reset()         { *m = DigitallySigned{} }
func (m *DigitallySigned) String() string { return proto.CompactTextString(m) }
func (*DigitallySigned) ProtoMessage()    {}
func (*DigitallySigned) Descriptor() ([]byte, []int) {
	return fileDescriptor_sigpb_62a21f00d4a205f5, []int{0}
}
func (m *DigitallySigned) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DigitallySigned.Unmarshal(m, b)
}
func (m *DigitallySigned) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DigitallySigned.Marshal(b, m, deterministic)
}
func (dst *DigitallySigned) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DigitallySigned.Merge(dst, src)
}
func (m *DigitallySigned) XXX_Size() int {
	return xxx_messageInfo_DigitallySigned.Size(m)
}
func (m *DigitallySigned) XXX_DiscardUnknown() {
	xxx_messageInfo_DigitallySigned.DiscardUnknown(m)
}

var xxx_messageInfo_DigitallySigned proto.InternalMessageInfo

func (m *DigitallySigned) GetHashAlgorithm() DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return DigitallySigned_NONE
}

func (m *DigitallySigned) GetSignatureAlgorithm() DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return DigitallySigned_ANONYMOUS
}

func (m *DigitallySigned) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*DigitallySigned)(nil), "sigpb.DigitallySigned")
	proto.RegisterEnum("sigpb.DigitallySigned_HashAlgorithm", DigitallySigned_HashAlgorithm_name, DigitallySigned_HashAlgorithm_value)
	proto.RegisterEnum("sigpb.DigitallySigned_SignatureAlgorithm", DigitallySigned_SignatureAlgorithm_name, DigitallySigned_SignatureAlgorithm_value)
}

func init() { proto.RegisterFile("crypto/sigpb/sigpb.proto", fileDescriptor_sigpb_62a21f00d4a205f5) }

var fileDescriptor_sigpb_62a21f00d4a205f5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0xa6, 0xad, 0x66, 0x30, 0x35, 0x8c, 0x97, 0x1c, 0x3c, 0x94, 0xa0, 0xd8, 0x5e,
	0x12, 0xa8, 0xa8, 0xe7, 0x68, 0x0b, 0x05, 0x31, 0x81, 0x2c, 0x1e, 0xec, 0x45, 0x36, 0x35, 0xec,
	0x2e, 0x6c, 0xb3, 0x21, 0xbb, 0x3d, 0xf4, 0x9f, 0xf9, 0xf3, 0x84, 0x94, 0xda, 0x68, 0xe9, 0x65,
	0xe0, 0x3d, 0xde, 0x7c, 0x33, 0x3c, 0xf0, 0x57, 0xf5, 0xb6, 0x32, 0x2a, 0xd2, 0x82, 0x55, 0xf9,
	0x6e, 0x86, 0x55, 0xad, 0x8c, 0xc2, 0x7e, 0x23, 0x82, 0xef, 0x2e, 0x5c, 0xce, 0x04, 0x13, 0x86,
	0x4a, 0xb9, 0x25, 0x82, 0x95, 0xc5, 0x17, 0xbe, 0xc2, 0x90, 0x53, 0xcd, 0x3f, 0xa9, 0x64, 0xaa,
	0x16, 0x86, 0xaf, 0x7d, 0x6b, 0x64, 0x8d, 0x87, 0xd3, 0x9b, 0x70, 0x07, 0xf8, 0x97, 0x0f, 0x17,
	0x54, 0xf3, 0x78, 0x9f, 0xcd, 0x5c, 0xde, 0x96, 0xb8, 0x84, 0x2b, 0x2d, 0x58, 0x49, 0xcd, 0xa6,
	0x2e, 0x5a, 0xc4, 0x6e, 0x43, 0x9c, 0x9c, 0x20, 0x92, 0xfd, 0xc6, 0x01, 0x8b, 0xfa, 0xc8, 0xc3,
	0x6b, 0x70, 0x7e, 0x5d, 0xdf, 0x1e, 0x59, 0xe3, 0x8b, 0xec, 0x60, 0x04, 0xb7, 0xe0, 0xfe, 0xf9,
	0x0c, 0xcf, 0xa1, 0x97, 0xa4, 0xc9, 0xdc, 0xeb, 0x20, 0xc0, 0x80, 0x2c, 0xe2, 0xe9, 0xc3, 0xa3,
	0xd7, 0x0b, 0x9e, 0x00, 0x8f, 0xcf, 0xa1, 0x0b, 0x4e, 0x9c, 0xa4, 0xc9, 0xc7, 0x5b, 0xfa, 0x4e,
	0xbc, 0x0e, 0x9e, 0x81, 0x9d, 0x91, 0xd8, 0xb3, 0xd0, 0x81, 0xfe, 0xfc, 0x65, 0x46, 0x62, 0xcf,
	0x7e, 0x9e, 0x2c, 0xef, 0x98, 0x30, 0x7c, 0x93, 0x87, 0x2b, 0xb5, 0x8e, 0x98, 0x52, 0x4c, 0x16,
	0x91, 0xa9, 0x85, 0x94, 0x82, 0x96, 0x51, 0xbb, 0xf8, 0x7c, 0xd0, 0x74, 0x7e, 0xff, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x04, 0xa5, 0x73, 0x8f, 0x01, 0x00, 0x00,
}
