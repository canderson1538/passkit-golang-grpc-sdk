// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/common/personal.proto

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

// Gender, as per government issued id.
type Gender int32

const (
	// Use only where gender is not known
	Gender_NOT_KNOWN Gender = 0
	// Male
	Gender_MALE Gender = 1
	// Female
	Gender_FEMALE Gender = 2
)

var Gender_name = map[int32]string{
	0: "NOT_KNOWN",
	1: "MALE",
	2: "FEMALE",
}

var Gender_value = map[string]int32{
	"NOT_KNOWN": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_206c14567a70d5f6, []int{0}
}

// A person represents a single, identifiable individual.
type Person struct {
	// Surname / Family name.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required,alpha"
	Surname string `protobuf:"bytes,1,opt,name=surname,proto3" json:"surname,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required,alpha"`
	// Forename / Given name.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required,alpha"
	Forename string `protobuf:"bytes,2,opt,name=forename,proto3" json:"forename,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required,alpha"`
	// Other names.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty,dive,alpha"
	OtherNames []string `protobuf:"bytes,3,rep,name=otherNames,proto3" json:"otherNames,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty,dive,alpha"`
	// Salutation or title.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"
	Salutation string `protobuf:"bytes,4,opt,name=salutation,proto3" json:"salutation,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"`
	// Suffix. For multiple suffixes, separate with spaces.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"
	Suffix string `protobuf:"bytes,5,opt,name=suffix,proto3" json:"suffix,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"`
	// If required, a string representing the user's preferred designation.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"isdefault"
	DisplayName string `protobuf:"bytes,6,opt,name=displayName,proto3" json:"displayName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"isdefault"`
	// Gender, as per government issued id.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required"
	Gender Gender `protobuf:"varint,7,opt,name=gender,proto3,enum=io.Gender" json:"gender,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"required"`
	// Date of birth.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"
	DateOfBirth *Date `protobuf:"bytes,8,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty"`
	// Email address.
	// @tag: validateGeneric:"omitempty,email" validateCreate:"omitempty,email" validateUpdate:"omitempty,email" validateFlight:"omitempty,email"
	EmailAddress string `protobuf:"bytes,9,opt,name=emailAddress,proto3" json:"emailAddress,omitempty" validateGeneric:"omitempty,email" validateCreate:"omitempty,email" validateUpdate:"omitempty,email" validateFlight:"omitempty,email"`
	// Telephone number. Enter in the format +{countryCode}{telephoneNumber} with no commas, spaces or other delimiters.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty,email"
	MobileNumber string `protobuf:"bytes,10,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateFlight:"omitempty,email"`
	// External id. Not editable.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	ExternalId           string   `protobuf:"bytes,11,opt,name=externalId,proto3" json:"externalId,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_206c14567a70d5f6, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Person) GetForename() string {
	if m != nil {
		return m.Forename
	}
	return ""
}

func (m *Person) GetOtherNames() []string {
	if m != nil {
		return m.OtherNames
	}
	return nil
}

func (m *Person) GetSalutation() string {
	if m != nil {
		return m.Salutation
	}
	return ""
}

func (m *Person) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *Person) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Person) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_NOT_KNOWN
}

func (m *Person) GetDateOfBirth() *Date {
	if m != nil {
		return m.DateOfBirth
	}
	return nil
}

func (m *Person) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Person) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

func (m *Person) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

type Address struct {
	AddressLine1         string   `protobuf:"bytes,1,opt,name=addressLine1,proto3" json:"addressLine1,omitempty"`
	AddressLine2         string   `protobuf:"bytes,2,opt,name=addressLine2,proto3" json:"addressLine2,omitempty"`
	AddressLine3         string   `protobuf:"bytes,3,opt,name=addressLine3,proto3" json:"addressLine3,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	ZipCode              string   `protobuf:"bytes,6,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	CountryCode          string   `protobuf:"bytes,7,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_206c14567a70d5f6, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddressLine1() string {
	if m != nil {
		return m.AddressLine1
	}
	return ""
}

func (m *Address) GetAddressLine2() string {
	if m != nil {
		return m.AddressLine2
	}
	return ""
}

func (m *Address) GetAddressLine3() string {
	if m != nil {
		return m.AddressLine3
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Address) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func init() {
	proto.RegisterEnum("io.Gender", Gender_name, Gender_value)
	proto.RegisterType((*Person)(nil), "io.Person")
	proto.RegisterType((*Address)(nil), "io.Address")
}

func init() {
	proto.RegisterFile("io/common/personal.proto", fileDescriptor_206c14567a70d5f6)
}

var fileDescriptor_206c14567a70d5f6 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xe7, 0x24, 0x75, 0x92, 0x93, 0xfd, 0x29, 0x62, 0x0c, 0xd1, 0x8b, 0x62, 0xc2, 0x2e,
	0x42, 0x61, 0x36, 0x4b, 0xee, 0x07, 0xed, 0xd6, 0x8d, 0x92, 0xce, 0x09, 0x61, 0x30, 0xd8, 0x4d,
	0x51, 0xec, 0x93, 0x46, 0xab, 0xe5, 0x63, 0x24, 0x19, 0x9a, 0x3d, 0xd2, 0x5e, 0x69, 0x2f, 0xb1,
	0x47, 0x18, 0x96, 0x9d, 0xd6, 0xe9, 0xae, 0xac, 0xef, 0xf7, 0x7d, 0x92, 0xd1, 0xf9, 0x04, 0x5c,
	0x52, 0x94, 0x90, 0x52, 0x94, 0x47, 0x05, 0x6a, 0x43, 0xb9, 0xc8, 0xc2, 0x42, 0x93, 0x25, 0xd6,
	0x91, 0x74, 0x72, 0xfa, 0xe8, 0xd6, 0x9f, 0x1b, 0x5a, 0xff, 0xc4, 0xc4, 0x9a, 0x3a, 0x33, 0xfe,
	0xdb, 0x01, 0x7f, 0xe9, 0xb6, 0x31, 0x0e, 0x7d, 0x53, 0xea, 0x5c, 0x28, 0xe4, 0x5e, 0xe0, 0x4d,
	0x86, 0xab, 0xbd, 0x64, 0x27, 0x30, 0xd8, 0x90, 0x46, 0x67, 0x75, 0x9c, 0xf5, 0xa0, 0xd9, 0x29,
	0x00, 0xd9, 0x2d, 0xea, 0x58, 0x28, 0x34, 0xbc, 0x1b, 0x74, 0x27, 0xc3, 0x55, 0x8b, 0x54, 0xbe,
	0x11, 0x59, 0x69, 0x85, 0x95, 0x94, 0xf3, 0x9e, 0xdb, 0xdd, 0x22, 0xec, 0x0d, 0xf8, 0xa6, 0xdc,
	0x6c, 0xe4, 0x3d, 0x3f, 0x72, 0x5e, 0xa3, 0x58, 0x00, 0xa3, 0x54, 0x9a, 0x22, 0x13, 0xbb, 0xea,
	0x1c, 0xee, 0x3b, 0xb3, 0x8d, 0xd8, 0x18, 0xfc, 0x5b, 0xcc, 0x53, 0xd4, 0xbc, 0x1f, 0x78, 0x93,
	0x97, 0x53, 0x08, 0x25, 0x85, 0x5f, 0x1c, 0x59, 0x35, 0x0e, 0x3b, 0x83, 0x51, 0x2a, 0x2c, 0x2e,
	0x36, 0x17, 0x52, 0xdb, 0x2d, 0x1f, 0x04, 0xde, 0x64, 0x34, 0x1d, 0x54, 0xc1, 0x4f, 0xc2, 0xe2,
	0xaa, 0x6d, 0xb2, 0x31, 0x3c, 0x47, 0x25, 0x64, 0x76, 0x9e, 0xa6, 0x1a, 0x8d, 0xe1, 0x43, 0xf7,
	0xcb, 0x03, 0x56, 0x65, 0x14, 0xad, 0x65, 0x86, 0x71, 0xa9, 0xd6, 0xa8, 0x39, 0xd4, 0x99, 0x36,
	0xab, 0x6e, 0x8c, 0xf7, 0x16, 0x75, 0x2e, 0xb2, 0xab, 0x94, 0x8f, 0xea, 0x1b, 0x3f, 0x92, 0xf1,
	0x1f, 0x0f, 0xfa, 0xad, 0xf3, 0x44, 0xbd, 0xbc, 0x96, 0x39, 0xbe, 0x6f, 0x06, 0x7f, 0xc0, 0x9e,
	0x64, 0xa6, 0x4d, 0x03, 0x07, 0xec, 0x49, 0x66, 0xc6, 0xbb, 0xff, 0x65, 0x66, 0x8c, 0x41, 0x2f,
	0x91, 0x76, 0xd7, 0x74, 0xe0, 0xd6, 0xec, 0x35, 0x1c, 0x19, 0x2b, 0x2c, 0x36, 0xc3, 0xaf, 0x45,
	0xf5, 0x12, 0x7e, 0xc9, 0xe2, 0x23, 0xa5, 0xfb, 0xb9, 0xef, 0x65, 0xd5, 0x4a, 0x42, 0x65, 0x6e,
	0xf5, 0xce, 0xb9, 0xfd, 0xba, 0x95, 0x16, 0x3a, 0x7b, 0x07, 0x7e, 0xdd, 0x01, 0x7b, 0x01, 0xc3,
	0x78, 0xf1, 0xed, 0x66, 0x1e, 0x2f, 0xbe, 0xc7, 0xc7, 0xcf, 0xd8, 0x00, 0x7a, 0x5f, 0xcf, 0xaf,
	0x2f, 0x8f, 0x3d, 0x06, 0xe0, 0x7f, 0xbe, 0x74, 0xeb, 0xce, 0xc5, 0x07, 0x78, 0x25, 0x29, 0x2c,
	0x84, 0x31, 0x77, 0xd2, 0x86, 0xcb, 0xb9, 0xa4, 0x1f, 0x6f, 0x8d, 0x15, 0x66, 0xfb, 0xc0, 0x12,
	0x52, 0x91, 0xa4, 0x48, 0x51, 0x8a, 0x59, 0x64, 0xd2, 0xbb, 0xe8, 0x96, 0x22, 0x49, 0xbf, 0x3b,
	0xbd, 0xe5, 0xfc, 0x8a, 0xd6, 0xbe, 0x7b, 0xc6, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0x6d, 0x6d, 0x83, 0x06, 0x03, 0x00, 0x00,
}
