// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/links.proto

package io

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

// Used to specify the type of link for link field. Each type has different icon on Google Pay.
type LinkType int32

const (
	// A link to website.
	LinkType_URI_DO_NOT_USE LinkType = 0
	// A link to website.
	LinkType_URI_WEB LinkType = 1
	// A phone number.
	LinkType_URI_TEL LinkType = 2
	// An email address.
	LinkType_URI_EMAIL LinkType = 3
	// A location address.
	LinkType_URI_LOCATION LinkType = 4
	// A calendar event.
	LinkType_URI_CALENDAR LinkType = 5
)

var LinkType_name = map[int32]string{
	0: "URI_DO_NOT_USE",
	1: "URI_WEB",
	2: "URI_TEL",
	3: "URI_EMAIL",
	4: "URI_LOCATION",
	5: "URI_CALENDAR",
}

var LinkType_value = map[string]int32{
	"URI_DO_NOT_USE": 0,
	"URI_WEB":        1,
	"URI_TEL":        2,
	"URI_EMAIL":      3,
	"URI_LOCATION":   4,
	"URI_CALENDAR":   5,
}

func (x LinkType) String() string {
	return proto.EnumName(LinkType_name, int32(x))
}

func (LinkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e346cf9813fdc45c, []int{0}
}

// Used to specify links put on the back of the pass.
type Link struct {
	// Link Id. Not writable.
	// @tag: validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"omitempty,uuidCompressedString" validateUpdate:"required,uuidCompressedString"`
	// A link text.
	// @tag: validateGeneric:"required_without=Id,uri"
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" validateGeneric:"required_without=Id,uri"`
	// Title for the link.
	// @tag: validateGeneric:"required_without=Id"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" validateGeneric:"required_without=Id"`
	// Type of link data.
	// @tag: validateGeneric:"min=0,max=4"
	Type LinkType `protobuf:"varint,4,opt,name=type,proto3,enum=io.LinkType" json:"type,omitempty" validateGeneric:"min=0,max=4"`
	// This customises link text for different languages. Ignored by Google Pay passes.
	LocalizedLink *LocalizedString `protobuf:"bytes,5,opt,name=localizedLink,proto3" json:"localizedLink,omitempty"`
	// This translates link title in different languages.
	LocalizedTitle *LocalizedString `protobuf:"bytes,6,opt,name=localizedTitle,proto3" json:"localizedTitle,omitempty"`
	// Indicates which wallets the link should be rendered for (this allows to hide certain link on Apple Wallet, and vise versa).
	// @tag: validateGeneric:"required,min=1"
	Usage []UsageType `protobuf:"varint,7,rep,packed,name=usage,proto3,enum=io.UsageType" json:"usage,omitempty" validateGeneric:"required,min=1"`
	// Links will be rendered in order of their position, from lowest to highest. Position can be any value, E.g. 3 links with positions 30, 10, 20 would render 10 first, 20 second and 30 third.  If no position is provided, the order of the links cannot be guaranteed.
	Position             uint32   `protobuf:"varint,8,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_e346cf9813fdc45c, []int{0}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Link) GetType() LinkType {
	if m != nil {
		return m.Type
	}
	return LinkType_URI_DO_NOT_USE
}

func (m *Link) GetLocalizedLink() *LocalizedString {
	if m != nil {
		return m.LocalizedLink
	}
	return nil
}

func (m *Link) GetLocalizedTitle() *LocalizedString {
	if m != nil {
		return m.LocalizedTitle
	}
	return nil
}

func (m *Link) GetUsage() []UsageType {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Link) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

type DbLink struct {
	Links                []*Link  `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DbLink) Reset()         { *m = DbLink{} }
func (m *DbLink) String() string { return proto.CompactTextString(m) }
func (*DbLink) ProtoMessage()    {}
func (*DbLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_e346cf9813fdc45c, []int{1}
}

func (m *DbLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbLink.Unmarshal(m, b)
}
func (m *DbLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbLink.Marshal(b, m, deterministic)
}
func (m *DbLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbLink.Merge(m, src)
}
func (m *DbLink) XXX_Size() int {
	return xxx_messageInfo_DbLink.Size(m)
}
func (m *DbLink) XXX_DiscardUnknown() {
	xxx_messageInfo_DbLink.DiscardUnknown(m)
}

var xxx_messageInfo_DbLink proto.InternalMessageInfo

func (m *DbLink) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.LinkType", LinkType_name, LinkType_value)
	proto.RegisterType((*Link)(nil), "io.Link")
	proto.RegisterType((*DbLink)(nil), "io.DbLink")
}

func init() {
	proto.RegisterFile("io/common/links.proto", fileDescriptor_e346cf9813fdc45c)
}

var fileDescriptor_e346cf9813fdc45c = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0xab, 0x9b, 0xe3, 0x1c, 0xc7, 0xaa, 0x98, 0xb6, 0x20, 0x4c, 0x09, 0x22, 0xed, 0x42,
	0x74, 0x21, 0x81, 0xbb, 0x2a, 0x85, 0x82, 0x13, 0x6b, 0x61, 0xa2, 0xda, 0x41, 0x91, 0x29, 0x74,
	0x63, 0x64, 0x4b, 0xb8, 0xa7, 0xba, 0x1c, 0xe1, 0x99, 0x2c, 0xd2, 0x37, 0xe9, 0x2b, 0xf4, 0x29,
	0xcb, 0x8c, 0x2c, 0xf7, 0x02, 0x59, 0x49, 0xff, 0xf7, 0xff, 0x47, 0xe7, 0x22, 0x78, 0x85, 0x14,
	0xee, 0xa8, 0xae, 0xa9, 0x09, 0x2b, 0x6c, 0x4a, 0x1e, 0xb4, 0x07, 0x12, 0xc4, 0x74, 0xa4, 0xc9,
	0xe5, 0x1f, 0xab, 0x7b, 0x6c, 0x68, 0xfb, 0xbd, 0xd8, 0x89, 0x63, 0x66, 0xf2, 0xfa, 0xaf, 0x52,
	0xda, 0x65, 0x15, 0xfe, 0xc8, 0x04, 0x52, 0xd3, 0xb9, 0x57, 0x3f, 0x75, 0x30, 0x63, 0x6c, 0x4a,
	0x66, 0x83, 0x8e, 0xb9, 0xab, 0x79, 0x9a, 0x7f, 0x9e, 0xe8, 0x98, 0x33, 0x07, 0x8c, 0x87, 0x43,
	0xe5, 0xea, 0x0a, 0xc8, 0x57, 0xf6, 0x12, 0x2c, 0x81, 0xa2, 0x2a, 0x5c, 0x43, 0xb1, 0x4e, 0x30,
	0x0f, 0x4c, 0xf1, 0xd8, 0x16, 0xae, 0xe9, 0x69, 0xbe, 0x3d, 0xbd, 0x08, 0x90, 0x02, 0xf9, 0xbd,
	0xf4, 0xb1, 0x2d, 0x12, 0xe5, 0xb0, 0x0f, 0x30, 0x3e, 0x36, 0x2e, 0x72, 0x69, 0xb9, 0x96, 0xa7,
	0xf9, 0xa3, 0xe9, 0x0b, 0x15, 0xed, 0x8d, 0x7b, 0x71, 0xc0, 0x66, 0x9f, 0xfc, 0x9b, 0x64, 0x1f,
	0xc1, 0x3e, 0x81, 0x54, 0xf5, 0x1e, 0x3c, 0x5d, 0xfb, 0x5f, 0x94, 0xbd, 0x01, 0xeb, 0x81, 0x67,
	0xfb, 0xc2, 0x3d, 0xf3, 0x0c, 0xdf, 0x9e, 0x8e, 0x65, 0xcd, 0x5a, 0x02, 0x35, 0x5b, 0xe7, 0xb1,
	0x09, 0x0c, 0x5b, 0xe2, 0x28, 0x2f, 0xe2, 0x0e, 0x3d, 0xcd, 0x1f, 0x27, 0x27, 0x7d, 0xe5, 0xc3,
	0x60, 0xbe, 0x55, 0x73, 0x5c, 0x82, 0xa5, 0xce, 0xee, 0x6a, 0x9e, 0xe1, 0x8f, 0xa6, 0xc3, 0x7e,
	0xcb, 0xa4, 0xc3, 0xef, 0x4a, 0x18, 0xf6, 0x4b, 0x33, 0x06, 0xf6, 0x3a, 0x59, 0x6c, 0xe6, 0xab,
	0xcd, 0x72, 0x95, 0x6e, 0xd6, 0xf7, 0x91, 0xf3, 0x8c, 0x8d, 0xe0, 0x4c, 0xb2, 0x2f, 0xd1, 0xb5,
	0xa3, 0xf5, 0x22, 0x8d, 0x62, 0x47, 0x67, 0x63, 0x38, 0x97, 0x22, 0xfa, 0x3c, 0x5b, 0xc4, 0x8e,
	0xc1, 0x1c, 0xb8, 0x90, 0x32, 0x5e, 0xdd, 0xcc, 0xd2, 0xc5, 0x6a, 0xe9, 0x98, 0x3d, 0xb9, 0x99,
	0xc5, 0xd1, 0x72, 0x3e, 0x4b, 0x1c, 0xeb, 0xfa, 0x13, 0x3c, 0x47, 0x0a, 0xda, 0x8c, 0xf3, 0x12,
	0x45, 0x70, 0x77, 0x8b, 0xf4, 0xf5, 0x2d, 0x17, 0x19, 0xff, 0x76, 0x62, 0x3b, 0xaa, 0x43, 0xa4,
	0xb0, 0xa6, 0xbc, 0xa8, 0x42, 0x9e, 0x97, 0xe1, 0x9e, 0x42, 0xa4, 0x5f, 0xba, 0x79, 0x77, 0xbb,
	0xa0, 0xed, 0x40, 0xfd, 0xf9, 0xf7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x09, 0x09, 0xe6, 0x6e,
	0x54, 0x02, 0x00, 0x00,
}
