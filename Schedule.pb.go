// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Schedule.proto

package spectre_go_entity

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

type Schedule struct {
	Items                []*Session `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	InterestFilter       string     `protobuf:"bytes,2,opt,name=interest_filter,json=interestFilter,proto3" json:"interest_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bc31960f3ce8c15, []int{0}
}

func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetItems() []*Session {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Schedule) GetInterestFilter() string {
	if m != nil {
		return m.InterestFilter
	}
	return ""
}

func init() {
	proto.RegisterType((*Schedule)(nil), "beilmo.spectre.entity.Schedule")
}

func init() { proto.RegisterFile("Schedule.proto", fileDescriptor_9bc31960f3ce8c15) }

var fileDescriptor_9bc31960f3ce8c15 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0b, 0x4e, 0xce, 0x48,
	0x4d, 0x29, 0xcd, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0x4a, 0xcd, 0xcc,
	0xc9, 0xcd, 0xd7, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0x29, 0x4a, 0xd5, 0x4b, 0xcd, 0x2b, 0xc9, 0x2c,
	0xa9, 0x94, 0xe2, 0x0d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x83, 0xa8, 0x52, 0xca, 0xe4, 0xe2,
	0x80, 0xe9, 0x13, 0x32, 0xe1, 0x62, 0xcd, 0x2c, 0x49, 0xcd, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x36, 0x92, 0xd3, 0xc3, 0x6a, 0x82, 0x1e, 0xd4, 0x80, 0x20, 0x88, 0x62, 0x21, 0x75, 0x2e,
	0xfe, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xd4, 0xe2, 0x92, 0xf8, 0xb4, 0xcc, 0x9c, 0x92, 0xd4, 0x22,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x3e, 0x98, 0xb0, 0x1b, 0x58, 0xd4, 0xc9, 0x99, 0x4b,
	0x32, 0x39, 0x3f, 0x17, 0xbb, 0xa1, 0x51, 0xca, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x10, 0x15, 0xfa, 0x50, 0x15, 0xba, 0xe9, 0xf9, 0xba, 0x10, 0x45, 0xbb, 0x98,
	0x18, 0x92, 0xd8, 0xc0, 0xce, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xae, 0x2e, 0x90,
	0xee, 0x00, 0x00, 0x00,
}
