// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeguide/route_guide.proto

package route_guide

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

type RouteNote struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dc806a642578ec, []int{0}
}

func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (m *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(m, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RouteNote)(nil), "RouteNote")
}

func init() { proto.RegisterFile("routeguide/route_guide.proto", fileDescriptor_d1dc806a642578ec) }

var fileDescriptor_d1dc806a642578ec = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xca, 0x2f, 0x2d,
	0x49, 0x4d, 0x2f, 0xcd, 0x4c, 0x49, 0xd5, 0x07, 0x33, 0xe3, 0xc1, 0x6c, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x25, 0x47, 0x2e, 0xce, 0x20, 0x90, 0xa0, 0x5f, 0x7e, 0x49, 0xaa, 0x90, 0x14, 0x17,
	0x47, 0x4e, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0x04, 0x96,
	0x82, 0x71, 0x8d, 0xcc, 0xb9, 0xb8, 0xc0, 0x46, 0xb8, 0x83, 0x8c, 0x15, 0xd2, 0x84, 0x1a, 0xe8,
	0x9c, 0x91, 0x58, 0x22, 0xc4, 0xa5, 0x07, 0x37, 0x5c, 0x0a, 0x89, 0xad, 0xc4, 0xa0, 0xc1, 0x68,
	0xc0, 0x98, 0xc4, 0x06, 0x76, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x33, 0x65, 0xe8, 0xa7,
	0xa2, 0x00, 0x00, 0x00,
}
