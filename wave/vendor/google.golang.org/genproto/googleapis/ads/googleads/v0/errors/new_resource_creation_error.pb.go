// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/new_resource_creation_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum_NewResourceCreationError int32

const (
	// Enum unspecified.
	NewResourceCreationErrorEnum_UNSPECIFIED NewResourceCreationErrorEnum_NewResourceCreationError = 0
	// The received error code is not known in this version.
	NewResourceCreationErrorEnum_UNKNOWN NewResourceCreationErrorEnum_NewResourceCreationError = 1
	// Do not set the id field while creating new resources.
	NewResourceCreationErrorEnum_CANNOT_SET_ID_FOR_CREATE NewResourceCreationErrorEnum_NewResourceCreationError = 2
	// Creating more than one resource with the same temp ID is not allowed.
	NewResourceCreationErrorEnum_DUPLICATE_TEMP_IDS NewResourceCreationErrorEnum_NewResourceCreationError = 3
	// Parent resource with specified temp ID failed validation, so no
	// validation will be done for this child resource.
	NewResourceCreationErrorEnum_TEMP_ID_RESOURCE_HAD_ERRORS NewResourceCreationErrorEnum_NewResourceCreationError = 4
)

var NewResourceCreationErrorEnum_NewResourceCreationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_SET_ID_FOR_CREATE",
	3: "DUPLICATE_TEMP_IDS",
	4: "TEMP_ID_RESOURCE_HAD_ERRORS",
}
var NewResourceCreationErrorEnum_NewResourceCreationError_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"CANNOT_SET_ID_FOR_CREATE":    2,
	"DUPLICATE_TEMP_IDS":          3,
	"TEMP_ID_RESOURCE_HAD_ERRORS": 4,
}

func (x NewResourceCreationErrorEnum_NewResourceCreationError) String() string {
	return proto.EnumName(NewResourceCreationErrorEnum_NewResourceCreationError_name, int32(x))
}
func (NewResourceCreationErrorEnum_NewResourceCreationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_new_resource_creation_error_c53e81c213476e90, []int{0, 0}
}

// Container for enum describing possible new resource creation errors.
type NewResourceCreationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewResourceCreationErrorEnum) Reset()         { *m = NewResourceCreationErrorEnum{} }
func (m *NewResourceCreationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NewResourceCreationErrorEnum) ProtoMessage()    {}
func (*NewResourceCreationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_new_resource_creation_error_c53e81c213476e90, []int{0}
}
func (m *NewResourceCreationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Unmarshal(m, b)
}
func (m *NewResourceCreationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Marshal(b, m, deterministic)
}
func (dst *NewResourceCreationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewResourceCreationErrorEnum.Merge(dst, src)
}
func (m *NewResourceCreationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NewResourceCreationErrorEnum.Size(m)
}
func (m *NewResourceCreationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NewResourceCreationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NewResourceCreationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewResourceCreationErrorEnum)(nil), "google.ads.googleads.v0.errors.NewResourceCreationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.NewResourceCreationErrorEnum_NewResourceCreationError", NewResourceCreationErrorEnum_NewResourceCreationError_name, NewResourceCreationErrorEnum_NewResourceCreationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/new_resource_creation_error.proto", fileDescriptor_new_resource_creation_error_c53e81c213476e90)
}

var fileDescriptor_new_resource_creation_error_c53e81c213476e90 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x6e, 0xa3, 0x30,
	0x18, 0x86, 0x07, 0x32, 0x9a, 0x91, 0x9c, 0xc5, 0x20, 0x16, 0xa3, 0x48, 0x93, 0xc9, 0x48, 0x1c,
	0xc0, 0x20, 0xcd, 0xce, 0xb3, 0x19, 0x07, 0x9c, 0x14, 0xb5, 0x05, 0x64, 0x20, 0x95, 0x2a, 0x24,
	0x8b, 0x06, 0x0b, 0x45, 0x4a, 0x70, 0x64, 0xe7, 0xe7, 0x18, 0xbd, 0x43, 0x97, 0x5d, 0xf5, 0x1c,
	0x3d, 0x4a, 0x4f, 0xd0, 0x65, 0x05, 0x26, 0xd9, 0xa5, 0x2b, 0x5e, 0xf8, 0x1e, 0x1e, 0x7f, 0x7e,
	0xc1, 0xff, 0x5a, 0x88, 0x7a, 0xcd, 0xdd, 0xb2, 0x52, 0xae, 0x8e, 0x6d, 0x3a, 0x78, 0x2e, 0x97,
	0x52, 0x48, 0xe5, 0x36, 0xfc, 0xc8, 0x24, 0x57, 0x62, 0x2f, 0x97, 0x9c, 0x2d, 0x25, 0x2f, 0x77,
	0x2b, 0xd1, 0xb0, 0x6e, 0x08, 0xb7, 0x52, 0xec, 0x84, 0x3d, 0xd1, 0xbf, 0xc1, 0xb2, 0x52, 0xf0,
	0x6c, 0x80, 0x07, 0x0f, 0x6a, 0x83, 0xf3, 0x62, 0x80, 0x71, 0xc4, 0x8f, 0xb4, 0x97, 0xf8, 0xbd,
	0x83, 0xb4, 0x53, 0xd2, 0xec, 0x37, 0xce, 0xa3, 0x01, 0x46, 0x97, 0x00, 0xfb, 0x07, 0x18, 0xe6,
	0x51, 0x9a, 0x10, 0x3f, 0x9c, 0x85, 0x24, 0xb0, 0xbe, 0xd8, 0x43, 0xf0, 0x3d, 0x8f, 0xae, 0xa3,
	0xf8, 0x2e, 0xb2, 0x0c, 0x7b, 0x0c, 0x46, 0x3e, 0x8e, 0xa2, 0x38, 0x63, 0x29, 0xc9, 0x58, 0x18,
	0xb0, 0x59, 0x4c, 0x99, 0x4f, 0x09, 0xce, 0x88, 0x65, 0xda, 0x3f, 0x81, 0x1d, 0xe4, 0xc9, 0x4d,
	0xe8, 0xe3, 0x8c, 0xb0, 0x8c, 0xdc, 0x26, 0x2c, 0x0c, 0x52, 0x6b, 0x60, 0xff, 0x01, 0xbf, 0xfa,
	0x37, 0x46, 0x49, 0x1a, 0xe7, 0xd4, 0x27, 0xec, 0x0a, 0x07, 0x8c, 0x50, 0x1a, 0xd3, 0xd4, 0xfa,
	0x3a, 0x7d, 0x37, 0x80, 0xb3, 0x14, 0x1b, 0xf8, 0xf9, 0xcd, 0xa6, 0xbf, 0x2f, 0x6d, 0x9d, 0xb4,
	0xc5, 0x24, 0xc6, 0x7d, 0xd0, 0x0b, 0x6a, 0xb1, 0x2e, 0x9b, 0x1a, 0x0a, 0x59, 0xbb, 0x35, 0x6f,
	0xba, 0xda, 0x4e, 0x65, 0x6f, 0x57, 0xea, 0x52, 0xf7, 0xff, 0xf4, 0xe3, 0xc9, 0x1c, 0xcc, 0x31,
	0x7e, 0x36, 0x27, 0x73, 0x2d, 0xc3, 0x95, 0x82, 0x3a, 0xb6, 0x69, 0xe1, 0xc1, 0xee, 0x48, 0xf5,
	0x7a, 0x02, 0x0a, 0x5c, 0xa9, 0xe2, 0x0c, 0x14, 0x0b, 0xaf, 0xd0, 0xc0, 0x9b, 0xe9, 0xe8, 0xaf,
	0x08, 0xe1, 0x4a, 0x21, 0x74, 0x46, 0x10, 0x5a, 0x78, 0x08, 0x69, 0xe8, 0xe1, 0x5b, 0xb7, 0xdd,
	0xdf, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x79, 0x1b, 0x37, 0x18, 0x02, 0x00, 0x00,
}
