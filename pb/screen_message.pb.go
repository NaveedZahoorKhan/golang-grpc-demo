// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package pb

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

type Screen_Panel int32

const (
	Screen_UNKNOW Screen_Panel = 0
	Screen_IPS    Screen_Panel = 1
	Screen_OLED   Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOW",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOW": 0,
	"IPS":    1,
	"OLED":   2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=example.com.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOW
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("example.com.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "example.com.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "example.com.Screen.Resolution")
}

func init() {
	proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7)
}

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xdd, 0x34, 0x89, 0x71, 0x4a, 0x25, 0x0c, 0x45, 0xa2, 0x42, 0x09, 0x3d, 0x48, 0x4e,
	0x11, 0xea, 0xcd, 0x83, 0x07, 0xd1, 0x43, 0x51, 0xda, 0xb2, 0x45, 0x04, 0x2f, 0x25, 0x8d, 0x43,
	0x77, 0x21, 0xd9, 0x0d, 0xd9, 0x0d, 0x8a, 0x8f, 0xe2, 0xd3, 0x4a, 0x13, 0xd1, 0x1c, 0x7a, 0x9c,
	0x8f, 0xdf, 0xf7, 0x87, 0x81, 0xb1, 0xc9, 0x6b, 0x22, 0xb5, 0x29, 0xc9, 0x98, 0x6c, 0x47, 0x69,
	0x55, 0x6b, 0xab, 0x71, 0x48, 0x9f, 0x59, 0x59, 0x15, 0x94, 0xe6, 0xba, 0x9c, 0x7e, 0x3b, 0xe0,
	0xaf, 0x5b, 0x0a, 0x2f, 0xe1, 0xc4, 0xc8, 0x2f, 0xda, 0x48, 0x95, 0x8b, 0x88, 0xc5, 0x2c, 0x71,
	0x78, 0xb0, 0x17, 0xe6, 0x2a, 0x17, 0x78, 0x07, 0x50, 0x93, 0xd1, 0x45, 0x63, 0xa5, 0x56, 0x91,
	0x13, 0xb3, 0x64, 0x38, 0x9b, 0xa4, 0xbd, 0xa4, 0xb4, 0x4b, 0x49, 0xf9, 0x1f, 0xc5, 0x7b, 0x0e,
	0xbc, 0x06, 0xaf, 0xca, 0x14, 0x15, 0xd1, 0x20, 0x66, 0xc9, 0xe9, 0xec, 0xfc, 0x90, 0x75, 0xb5,
	0x07, 0x78, 0xc7, 0xe1, 0x04, 0xa0, 0x6c, 0x0a, 0x2b, 0xad, 0x6e, 0x72, 0x11, 0xb9, 0x31, 0x4b,
	0x02, 0xde, 0x53, 0x2e, 0x6e, 0x01, 0xfe, 0xab, 0x70, 0x0c, 0xde, 0x87, 0x7c, 0xb7, 0xdd, 0xee,
	0x11, 0xef, 0x0e, 0x3c, 0x03, 0x5f, 0x90, 0xdc, 0x09, 0xdb, 0x0e, 0x1e, 0xf1, 0xdf, 0x6b, 0x7a,
	0x05, 0x5e, 0xdb, 0x85, 0x00, 0xfe, 0xcb, 0xe2, 0x69, 0xb1, 0x7c, 0x0d, 0x8f, 0xf0, 0x18, 0x06,
	0xf3, 0xd5, 0x3a, 0x64, 0x18, 0x80, 0xbb, 0x7c, 0x7e, 0x7c, 0x08, 0x9d, 0x7b, 0xf7, 0xcd, 0xa9,
	0xb6, 0x5b, 0xbf, 0x7d, 0xdb, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x58, 0x09, 0x97,
	0x4e, 0x01, 0x00, 0x00,
}
