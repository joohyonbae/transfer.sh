// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/interaction_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enum describing possible interaction types.
type InteractionTypeEnum_InteractionType int32

const (
	// Not specified.
	InteractionTypeEnum_UNSPECIFIED InteractionTypeEnum_InteractionType = 0
	// Used for return value only. Represents value unknown in this version.
	InteractionTypeEnum_UNKNOWN InteractionTypeEnum_InteractionType = 1
	// Calls.
	InteractionTypeEnum_CALLS InteractionTypeEnum_InteractionType = 8000
)

var InteractionTypeEnum_InteractionType_name = map[int32]string{
	0:    "UNSPECIFIED",
	1:    "UNKNOWN",
	8000: "CALLS",
}
var InteractionTypeEnum_InteractionType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CALLS":       8000,
}

func (x InteractionTypeEnum_InteractionType) String() string {
	return proto.EnumName(InteractionTypeEnum_InteractionType_name, int32(x))
}
func (InteractionTypeEnum_InteractionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_interaction_type_c1d6ddb209db8963, []int{0, 0}
}

// Container for enum describing possible interaction types.
type InteractionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractionTypeEnum) Reset()         { *m = InteractionTypeEnum{} }
func (m *InteractionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*InteractionTypeEnum) ProtoMessage()    {}
func (*InteractionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_interaction_type_c1d6ddb209db8963, []int{0}
}
func (m *InteractionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractionTypeEnum.Unmarshal(m, b)
}
func (m *InteractionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *InteractionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractionTypeEnum.Merge(dst, src)
}
func (m *InteractionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_InteractionTypeEnum.Size(m)
}
func (m *InteractionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_InteractionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InteractionTypeEnum)(nil), "google.ads.googleads.v0.enums.InteractionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.InteractionTypeEnum_InteractionType", InteractionTypeEnum_InteractionType_name, InteractionTypeEnum_InteractionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/interaction_type.proto", fileDescriptor_interaction_type_c1d6ddb209db8963)
}

var fileDescriptor_interaction_type_c1d6ddb209db8963 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xf8,
	0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x52, 0xbd, 0xc4,
	0x94, 0x62, 0x3d, 0xb8, 0x2e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x2e, 0xa5, 0x20, 0x2e, 0x61, 0x4f,
	0x84, 0xc6, 0x90, 0xca, 0x82, 0x54, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x6b, 0x2e, 0x7e, 0x34, 0x61,
	0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01,
	0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x2e,
	0x2e, 0x56, 0x67, 0x47, 0x1f, 0x9f, 0x60, 0x81, 0x03, 0x76, 0x4e, 0x2f, 0x19, 0xb9, 0x14, 0x93,
	0xf3, 0x73, 0xf5, 0xf0, 0xda, 0xec, 0x24, 0x82, 0x66, 0x41, 0x00, 0xc8, 0xb9, 0x01, 0x8c, 0x51,
	0x4e, 0x50, 0x6d, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9,
	0x79, 0x60, 0xcf, 0xc0, 0xbc, 0x5d, 0x90, 0x59, 0x8c, 0x23, 0x14, 0xac, 0xc1, 0xe4, 0x22, 0x26,
	0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2, 0xee, 0x10, 0xa3, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c,
	0x10, 0x2b, 0xcc, 0x40, 0x0f, 0xe4, 0xc7, 0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7, 0x94, 0xe2, 0x18,
	0xb8, 0x7c, 0x4c, 0x98, 0x41, 0x0c, 0x58, 0xfe, 0x15, 0x93, 0x22, 0x44, 0xd0, 0xca, 0xca, 0x31,
	0xa5, 0xd8, 0xca, 0x0a, 0xae, 0xc2, 0xca, 0x2a, 0xcc, 0xc0, 0xca, 0x0a, 0xac, 0x26, 0x89, 0x0d,
	0xec, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x07, 0xdd, 0xc0, 0x9d, 0x01, 0x00,
	0x00,
}