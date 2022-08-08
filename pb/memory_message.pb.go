// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memory_message.proto

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

type Memory_Unit int32

const (
	Memory_UNIT_BYTE      Memory_Unit = 0
	Memory_UNIT_KILOBYTE  Memory_Unit = 1
	Memory_UNIT_MEGABYTE  Memory_Unit = 2
	Memory_UNIT_GIGABYTE  Memory_Unit = 3
	Memory_UNIT_TERABYTE  Memory_Unit = 4
	Memory_UNIT_PETABYTE  Memory_Unit = 5
	Memory_UNIT_EXABYTE   Memory_Unit = 6
	Memory_UNIT_ZETTABYTE Memory_Unit = 7
	Memory_UNIT_YOTTABYTE Memory_Unit = 8
	Memory_UNKNOW         Memory_Unit = 9
)

var Memory_Unit_name = map[int32]string{
	0: "UNIT_BYTE",
	1: "UNIT_KILOBYTE",
	2: "UNIT_MEGABYTE",
	3: "UNIT_GIGABYTE",
	4: "UNIT_TERABYTE",
	5: "UNIT_PETABYTE",
	6: "UNIT_EXABYTE",
	7: "UNIT_ZETTABYTE",
	8: "UNIT_YOTTABYTE",
	9: "UNKNOW",
}

var Memory_Unit_value = map[string]int32{
	"UNIT_BYTE":      0,
	"UNIT_KILOBYTE":  1,
	"UNIT_MEGABYTE":  2,
	"UNIT_GIGABYTE":  3,
	"UNIT_TERABYTE":  4,
	"UNIT_PETABYTE":  5,
	"UNIT_EXABYTE":   6,
	"UNIT_ZETTABYTE": 7,
	"UNIT_YOTTABYTE": 8,
	"UNKNOW":         9,
}

func (x Memory_Unit) String() string {
	return proto.EnumName(Memory_Unit_name, int32(x))
}

func (Memory_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0, 0}
}

type Memory struct {
	Value                uint64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 Memory_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=example.com.Memory_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0}
}

func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Memory) GetUnit() Memory_Unit {
	if m != nil {
		return m.Unit
	}
	return Memory_UNIT_BYTE
}

func init() {
	proto.RegisterEnum("example.com.Memory_Unit", Memory_Unit_name, Memory_Unit_value)
	proto.RegisterType((*Memory)(nil), "example.com.Memory")
}

func init() {
	proto.RegisterFile("memory_message.proto", fileDescriptor_c0c7f919ccc765da)
}

var fileDescriptor_c0c7f919ccc765da = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcd, 0xcd,
	0x2f, 0xaa, 0x8c, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x4e, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x4b, 0xce, 0xcf, 0x55, 0xea, 0x64,
	0xe2, 0x62, 0xf3, 0x05, 0xab, 0x12, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x70, 0x84, 0x74, 0xb8, 0x58, 0x4a, 0xf3, 0x32, 0x4b, 0x24,
	0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x24, 0xf4, 0x90, 0x34, 0xeb, 0x41, 0x34, 0xea, 0x85, 0xe6,
	0x65, 0x96, 0x04, 0x81, 0x55, 0x29, 0xed, 0x62, 0xe4, 0x62, 0x01, 0x71, 0x85, 0x78, 0xb9, 0x38,
	0x43, 0xfd, 0x3c, 0x43, 0xe2, 0x9d, 0x22, 0x43, 0x5c, 0x05, 0x18, 0x84, 0x04, 0xb9, 0x78, 0xc1,
	0x5c, 0x6f, 0x4f, 0x1f, 0x7f, 0xb0, 0x10, 0x23, 0x5c, 0xc8, 0xd7, 0xd5, 0xdd, 0x11, 0x2c, 0xc4,
	0x04, 0x17, 0x72, 0xf7, 0x84, 0x0a, 0x31, 0xc3, 0x85, 0x42, 0x5c, 0x83, 0x20, 0x42, 0x2c, 0x70,
	0xa1, 0x00, 0xd7, 0x10, 0x88, 0x10, 0xab, 0x90, 0x00, 0x17, 0x0f, 0x58, 0xc8, 0x35, 0x02, 0x22,
	0xc2, 0x26, 0x24, 0xc4, 0xc5, 0x07, 0x16, 0x89, 0x72, 0x0d, 0x81, 0xaa, 0x62, 0x87, 0x8b, 0x45,
	0xfa, 0xc3, 0xc4, 0x38, 0x84, 0xb8, 0xb8, 0xd8, 0x42, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0x05, 0x38,
	0x9d, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0xa1, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x08, 0xf9, 0x9a, 0x3d, 0x01, 0x00, 0x00,
}