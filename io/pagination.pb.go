// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/pagination.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Pagination struct {
	// Limit the number of records returned. If not specified, a default of 25 is used.  Enter -1 for all records.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Allows you to offset the first record returned by the limit.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Array of column names to filter results by.
	FilterField []string `protobuf:"bytes,3,rep,name=filterField,proto3" json:"filterField,omitempty"`
	// Array of values to test against the filter fields.
	FilterValue []string `protobuf:"bytes,4,rep,name=filterValue,proto3" json:"filterValue,omitempty"`
	// will be whitelisted operators in io core
	FilterOperator []string `protobuf:"bytes,5,rep,name=filterOperator,proto3" json:"filterOperator,omitempty"`
	// Field to order results by.
	OrderBy string `protobuf:"bytes,6,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	// Will return in ascending order if true, or descending order if false.
	OrderAsc             bool     `protobuf:"varint,7,opt,name=orderAsc,proto3" json:"orderAsc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_3022e52cf9a6ff9e, []int{0}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Pagination) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Pagination) GetFilterField() []string {
	if m != nil {
		return m.FilterField
	}
	return nil
}

func (m *Pagination) GetFilterValue() []string {
	if m != nil {
		return m.FilterValue
	}
	return nil
}

func (m *Pagination) GetFilterOperator() []string {
	if m != nil {
		return m.FilterOperator
	}
	return nil
}

func (m *Pagination) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *Pagination) GetOrderAsc() bool {
	if m != nil {
		return m.OrderAsc
	}
	return false
}

type Filter struct {
	// Array of column names to filter results by.
	FilterField []string `protobuf:"bytes,1,rep,name=filterField,proto3" json:"filterField,omitempty"`
	// Array of values to test against the filter fields.
	FilterValue []string `protobuf:"bytes,2,rep,name=filterValue,proto3" json:"filterValue,omitempty"`
	// will be whitelisted operators in io core
	FilterOperator       []string `protobuf:"bytes,3,rep,name=filterOperator,proto3" json:"filterOperator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3022e52cf9a6ff9e, []int{1}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetFilterField() []string {
	if m != nil {
		return m.FilterField
	}
	return nil
}

func (m *Filter) GetFilterValue() []string {
	if m != nil {
		return m.FilterValue
	}
	return nil
}

func (m *Filter) GetFilterOperator() []string {
	if m != nil {
		return m.FilterOperator
	}
	return nil
}

func init() {
	proto.RegisterType((*Pagination)(nil), "io.Pagination")
	proto.RegisterType((*Filter)(nil), "io.Filter")
}

func init() {
	proto.RegisterFile("io/common/pagination.proto", fileDescriptor_3022e52cf9a6ff9e)
}

var fileDescriptor_3022e52cf9a6ff9e = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x69, 0xb7, 0x75, 0x5b, 0xfe, 0xf0, 0x17, 0x8a, 0x48, 0xd8, 0x29, 0x0c, 0x91, 0x22,
	0xae, 0x01, 0xbd, 0x29, 0x08, 0xdb, 0x61, 0x30, 0x06, 0x3a, 0x7a, 0xf0, 0xe0, 0x2d, 0x6b, 0xb3,
	0x1a, 0x97, 0xe6, 0x2d, 0xc9, 0x3b, 0xc4, 0xaf, 0xe1, 0xc7, 0x10, 0x3f, 0xa4, 0x2c, 0xd3, 0x6d,
	0x0c, 0x41, 0x6f, 0xcf, 0xf3, 0xe3, 0x79, 0x4b, 0x7f, 0x84, 0xf4, 0x14, 0xf0, 0x1c, 0xaa, 0x0a,
	0x0c, 0xaf, 0x45, 0xa9, 0x8c, 0x40, 0x05, 0x26, 0xad, 0x2d, 0x20, 0xc4, 0xa1, 0x82, 0xde, 0x85,
	0x8f, 0xf9, 0xa0, 0x94, 0x66, 0xe0, 0x5e, 0x44, 0x59, 0x4a, 0xcb, 0xa1, 0x5e, 0x8f, 0x1c, 0x17,
	0xc6, 0x00, 0xfa, 0x03, 0xb7, 0xb9, 0xe8, 0x7f, 0x84, 0x84, 0xcc, 0xb6, 0x9f, 0x89, 0x8f, 0x49,
	0x4b, 0xab, 0x4a, 0x21, 0x0d, 0x58, 0x90, 0xb4, 0xb2, 0x4d, 0x89, 0x4f, 0x48, 0x04, 0x8b, 0x85,
	0x93, 0x48, 0x43, 0x8f, 0xbf, 0x5a, 0xcc, 0xc8, 0xbf, 0x85, 0xd2, 0x28, 0xed, 0x58, 0x49, 0x5d,
	0xd0, 0x06, 0x6b, 0x24, 0xdd, 0x6c, 0x1f, 0xed, 0x16, 0x0f, 0x42, 0xaf, 0x24, 0x6d, 0xee, 0x2f,
	0x3c, 0x8a, 0xcf, 0xc8, 0xff, 0x4d, 0xbd, 0xaf, 0xa5, 0x15, 0x08, 0x96, 0xb6, 0xfc, 0xe8, 0x80,
	0xc6, 0x94, 0xb4, 0xc1, 0x16, 0xd2, 0x8e, 0x5e, 0x69, 0xc4, 0x82, 0xa4, 0x9b, 0x7d, 0xd7, 0xb8,
	0x47, 0x3a, 0x3e, 0x0e, 0x5d, 0x4e, 0xdb, 0x2c, 0x48, 0x3a, 0xd9, 0xb6, 0x5f, 0xdf, 0xbd, 0x0d,
	0xa7, 0x64, 0x72, 0xbe, 0xa7, 0x78, 0x79, 0xb3, 0xcb, 0x0c, 0xe6, 0xcf, 0x32, 0x47, 0xb6, 0x72,
	0xb2, 0x60, 0x08, 0xcc, 0xd5, 0x5a, 0x21, 0xcb, 0xc1, 0xa0, 0x34, 0xc8, 0x94, 0x59, 0xa3, 0x4a,
	0x68, 0x2d, 0x2d, 0xab, 0x85, 0x45, 0x97, 0xf6, 0x91, 0x44, 0x63, 0xff, 0x5f, 0x87, 0xee, 0xc1,
	0xaf, 0xee, 0xe1, 0x5f, 0xdc, 0x1b, 0x3f, 0xb9, 0x8f, 0x6e, 0xc9, 0x91, 0x82, 0xb4, 0x16, 0xce,
	0x2d, 0x15, 0xa6, 0xb3, 0xa9, 0x82, 0xc7, 0x53, 0x87, 0xc2, 0x3d, 0x6d, 0x59, 0x0e, 0x15, 0x57,
	0xc0, 0x2b, 0x28, 0xa4, 0xe6, 0xae, 0x58, 0xf2, 0x12, 0xb8, 0x82, 0xf7, 0xb0, 0x39, 0x9b, 0x4e,
	0x60, 0x1e, 0xf9, 0xb7, 0xbe, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe5, 0x27, 0xaa, 0x3b,
	0x02, 0x00, 0x00,
}
