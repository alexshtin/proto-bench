// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structdef-go-v1.proto

package protobench

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

type TypeV1 int32

const (
	TypeV1_TYPEV1_UNSPECIFIED TypeV1 = 0
	TypeV1_TYPEV1_R           TypeV1 = 1
	TypeV1_TYPEV1_S           TypeV1 = 2
)

var TypeV1_name = map[int32]string{
	0: "TYPEV1_UNSPECIFIED",
	1: "TYPEV1_R",
	2: "TYPEV1_S",
}

var TypeV1_value = map[string]int32{
	"TYPEV1_UNSPECIFIED": 0,
	"TYPEV1_R":           1,
	"TYPEV1_S":           2,
}

func (x TypeV1) String() string {
	return proto.EnumName(TypeV1_name, int32(x))
}

func (TypeV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9e762851f95667d1, []int{0}
}

type GoV1 struct {
	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BirthDay int64   `protobuf:"varint,2,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Siblings int32   `protobuf:"varint,4,opt,name=siblings,proto3" json:"siblings,omitempty"`
	Spouse   bool    `protobuf:"varint,5,opt,name=spouse,proto3" json:"spouse,omitempty"`
	Money    float64 `protobuf:"fixed64,6,opt,name=money,proto3" json:"money,omitempty"`
	Type     TypeV1  `protobuf:"varint,7,opt,name=type,proto3,enum=protobench.TypeV1" json:"type,omitempty"`
	// Types that are valid to be assigned to Values:
	//	*GoV1_ValueS
	//	*GoV1_ValueI
	//	*GoV1_ValueD
	Values               isGoV1_Values `protobuf_oneof:"values"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GoV1) Reset()         { *m = GoV1{} }
func (m *GoV1) String() string { return proto.CompactTextString(m) }
func (*GoV1) ProtoMessage()    {}
func (*GoV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e762851f95667d1, []int{0}
}

func (m *GoV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoV1.Unmarshal(m, b)
}
func (m *GoV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoV1.Marshal(b, m, deterministic)
}
func (m *GoV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoV1.Merge(m, src)
}
func (m *GoV1) XXX_Size() int {
	return xxx_messageInfo_GoV1.Size(m)
}
func (m *GoV1) XXX_DiscardUnknown() {
	xxx_messageInfo_GoV1.DiscardUnknown(m)
}

var xxx_messageInfo_GoV1 proto.InternalMessageInfo

func (m *GoV1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GoV1) GetBirthDay() int64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *GoV1) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *GoV1) GetSiblings() int32 {
	if m != nil {
		return m.Siblings
	}
	return 0
}

func (m *GoV1) GetSpouse() bool {
	if m != nil {
		return m.Spouse
	}
	return false
}

func (m *GoV1) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *GoV1) GetType() TypeV1 {
	if m != nil {
		return m.Type
	}
	return TypeV1_TYPEV1_UNSPECIFIED
}

type isGoV1_Values interface {
	isGoV1_Values()
}

type GoV1_ValueS struct {
	ValueS string `protobuf:"bytes,8,opt,name=value_s,json=valueS,proto3,oneof"`
}

type GoV1_ValueI struct {
	ValueI int32 `protobuf:"varint,9,opt,name=value_i,json=valueI,proto3,oneof"`
}

type GoV1_ValueD struct {
	ValueD float64 `protobuf:"fixed64,10,opt,name=value_d,json=valueD,proto3,oneof"`
}

func (*GoV1_ValueS) isGoV1_Values() {}

func (*GoV1_ValueI) isGoV1_Values() {}

func (*GoV1_ValueD) isGoV1_Values() {}

func (m *GoV1) GetValues() isGoV1_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *GoV1) GetValueS() string {
	if x, ok := m.GetValues().(*GoV1_ValueS); ok {
		return x.ValueS
	}
	return ""
}

func (m *GoV1) GetValueI() int32 {
	if x, ok := m.GetValues().(*GoV1_ValueI); ok {
		return x.ValueI
	}
	return 0
}

func (m *GoV1) GetValueD() float64 {
	if x, ok := m.GetValues().(*GoV1_ValueD); ok {
		return x.ValueD
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GoV1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GoV1_ValueS)(nil),
		(*GoV1_ValueI)(nil),
		(*GoV1_ValueD)(nil),
	}
}

func init() {
	proto.RegisterEnum("protobench.TypeV1", TypeV1_name, TypeV1_value)
	proto.RegisterType((*GoV1)(nil), "protobench.GoV1")
}

func init() {
	proto.RegisterFile("structdef-go-v1.proto", fileDescriptor_9e762851f95667d1)
}

var fileDescriptor_9e762851f95667d1 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x59, 0x28, 0xa5, 0x4c, 0x8c, 0x21, 0x13, 0x25, 0xab, 0xa7, 0xc6, 0x83, 0x69, 0x88,
	0x94, 0x54, 0x8f, 0x7a, 0x42, 0x50, 0xb9, 0x18, 0xb2, 0x20, 0x89, 0x5e, 0x48, 0x0b, 0x2b, 0x6d,
	0x02, 0xdd, 0xa6, 0xbb, 0x25, 0xf6, 0x51, 0x7c, 0x5b, 0xd3, 0x2d, 0x69, 0x3d, 0xf5, 0xff, 0xe6,
	0xcb, 0x74, 0xfe, 0x85, 0x4b, 0xa9, 0xd2, 0x6c, 0xa3, 0xb6, 0xfc, 0x7b, 0xb8, 0x13, 0xc3, 0xa3,
	0xe7, 0x26, 0xa9, 0x50, 0x02, 0x41, 0x7f, 0x02, 0x1e, 0x6f, 0xc2, 0x9b, 0xdf, 0x26, 0x18, 0xaf,
	0x62, 0xe5, 0x21, 0x82, 0x11, 0xfb, 0x07, 0x4e, 0x89, 0x4d, 0x9c, 0x2e, 0xd3, 0x19, 0xaf, 0xc1,
	0x0a, 0xa2, 0x54, 0x85, 0x13, 0x3f, 0xa7, 0x4d, 0x9b, 0x38, 0x2d, 0x56, 0x31, 0x5e, 0x40, 0x3b,
	0x09, 0x45, 0xcc, 0x69, 0x4b, 0x2f, 0x94, 0x50, 0x6c, 0xc8, 0x28, 0xd8, 0x47, 0xf1, 0x4e, 0x52,
	0xc3, 0x26, 0x4e, 0x9b, 0x55, 0x8c, 0x7d, 0x30, 0x65, 0x22, 0x32, 0xc9, 0x69, 0xdb, 0x26, 0x8e,
	0xc5, 0x4e, 0x54, 0xfc, 0xe9, 0x20, 0x62, 0x9e, 0x53, 0xd3, 0x26, 0x0e, 0x61, 0x25, 0xe0, 0x2d,
	0x18, 0x2a, 0x4f, 0x38, 0xed, 0xd8, 0xc4, 0x39, 0xbf, 0x47, 0xb7, 0xee, 0xec, 0x2e, 0xf3, 0x84,
	0xaf, 0x3c, 0xa6, 0x3d, 0x5e, 0x41, 0xe7, 0xe8, 0xef, 0x33, 0xbe, 0x96, 0xd4, 0x2a, 0x9a, 0xbc,
	0x35, 0x98, 0xa9, 0x07, 0x8b, 0x5a, 0x45, 0xb4, 0x5b, 0x74, 0xa9, 0xd4, 0xac, 0x56, 0x5b, 0x0a,
	0xc5, 0xd5, 0x4a, 0x4d, 0xc6, 0x16, 0x94, 0x49, 0x0e, 0x9e, 0xc0, 0x2c, 0x4f, 0x61, 0x1f, 0x70,
	0xf9, 0x39, 0x9f, 0xae, 0xbc, 0xf5, 0xc7, 0xfb, 0x62, 0x3e, 0x7d, 0x9e, 0xbd, 0xcc, 0xa6, 0x93,
	0x5e, 0x03, 0xcf, 0xc0, 0x3a, 0xcd, 0x59, 0x8f, 0xfc, 0xa3, 0x45, 0xaf, 0x39, 0xbe, 0xfb, 0x1a,
	0xec, 0x22, 0x15, 0x66, 0x81, 0xbb, 0x11, 0x87, 0x91, 0xbf, 0xe7, 0x3f, 0x32, 0x54, 0x51, 0x3c,
	0xaa, 0x1f, 0xf2, 0x58, 0xc7, 0xc0, 0xd4, 0xf9, 0xe1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x17, 0x3e,
	0x1c, 0xe5, 0xb3, 0x01, 0x00, 0x00,
}
