// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/flights/carrier.proto

package flights

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	math "math"
	io "stash.passkit.com/io/model/sdk/go/io"
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

// A carrier record contains details of the carrier.  The carrier will be the issuer of the boarding pass and may be used as a marketing or operating carrier.
type Carrier struct {
	// The IATA carrier code. If the carrier has not been issued an IATA carrier code, use YY.
	IataCarrierCode string `protobuf:"bytes,1,opt,name=iataCarrierCode,proto3" json:"iataCarrierCode,omitempty"`
	// The ICAO carrier code.
	IcaoCarrierCode string `protobuf:"bytes,2,opt,name=icaoCarrierCode,proto3" json:"icaoCarrierCode,omitempty"`
	// The IATA accounting code / AWB prefix.  If no code has been allocated, enter zero.
	IataAccountingCode int32 `protobuf:"varint,3,opt,name=iataAccountingCode,proto3" json:"iataAccountingCode,omitempty"`
	// The name of the airline. This will be printed below the logo on the Google Pay pass.
	AirlineName string `protobuf:"bytes,4,opt,name=airlineName,proto3" json:"airlineName,omitempty"`
	// The localized airline name, if applicable.
	LocalizedAirlineName *io.LocalizedString `protobuf:"bytes,5,opt,name=localizedAirlineName,proto3" json:"localizedAirlineName,omitempty"`
	// If the carrier will issue Apple Wallet passes, supply the certificate id (E.g. pass.com.passkitair). The certificate must have previously been uploaded.
	PassTypeIdentifier string `protobuf:"bytes,6,opt,name=passTypeIdentifier,proto3" json:"passTypeIdentifier,omitempty"`
	// If PassKit are generating the boarding pass barcode, and the barcode needs signing, the ID of the private key used to sign the barcode payload.
	PrivateKeyId         string   `protobuf:"bytes,7,opt,name=privateKeyId,proto3" json:"privateKeyId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Carrier) Reset()         { *m = Carrier{} }
func (m *Carrier) String() string { return proto.CompactTextString(m) }
func (*Carrier) ProtoMessage()    {}
func (*Carrier) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b8f33b242f205b, []int{0}
}

func (m *Carrier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Carrier.Unmarshal(m, b)
}
func (m *Carrier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Carrier.Marshal(b, m, deterministic)
}
func (m *Carrier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Carrier.Merge(m, src)
}
func (m *Carrier) XXX_Size() int {
	return xxx_messageInfo_Carrier.Size(m)
}
func (m *Carrier) XXX_DiscardUnknown() {
	xxx_messageInfo_Carrier.DiscardUnknown(m)
}

var xxx_messageInfo_Carrier proto.InternalMessageInfo

func (m *Carrier) GetIataCarrierCode() string {
	if m != nil {
		return m.IataCarrierCode
	}
	return ""
}

func (m *Carrier) GetIcaoCarrierCode() string {
	if m != nil {
		return m.IcaoCarrierCode
	}
	return ""
}

func (m *Carrier) GetIataAccountingCode() int32 {
	if m != nil {
		return m.IataAccountingCode
	}
	return 0
}

func (m *Carrier) GetAirlineName() string {
	if m != nil {
		return m.AirlineName
	}
	return ""
}

func (m *Carrier) GetLocalizedAirlineName() *io.LocalizedString {
	if m != nil {
		return m.LocalizedAirlineName
	}
	return nil
}

func (m *Carrier) GetPassTypeIdentifier() string {
	if m != nil {
		return m.PassTypeIdentifier
	}
	return ""
}

func (m *Carrier) GetPrivateKeyId() string {
	if m != nil {
		return m.PrivateKeyId
	}
	return ""
}

// CarrierCode is used for retrieving or deleting a port record.
type CarrierCode struct {
	// The IATA or ICAO carrier code.
	CarrierCode          string   `protobuf:"bytes,1,opt,name=carrierCode,proto3" json:"carrierCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarrierCode) Reset()         { *m = CarrierCode{} }
func (m *CarrierCode) String() string { return proto.CompactTextString(m) }
func (*CarrierCode) ProtoMessage()    {}
func (*CarrierCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_72b8f33b242f205b, []int{1}
}

func (m *CarrierCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarrierCode.Unmarshal(m, b)
}
func (m *CarrierCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarrierCode.Marshal(b, m, deterministic)
}
func (m *CarrierCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarrierCode.Merge(m, src)
}
func (m *CarrierCode) XXX_Size() int {
	return xxx_messageInfo_CarrierCode.Size(m)
}
func (m *CarrierCode) XXX_DiscardUnknown() {
	xxx_messageInfo_CarrierCode.DiscardUnknown(m)
}

var xxx_messageInfo_CarrierCode proto.InternalMessageInfo

func (m *CarrierCode) GetCarrierCode() string {
	if m != nil {
		return m.CarrierCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Carrier)(nil), "flights.Carrier")
	proto.RegisterType((*CarrierCode)(nil), "flights.CarrierCode")
}

func init() { proto.RegisterFile("io/flights/carrier.proto", fileDescriptor_72b8f33b242f205b) }

var fileDescriptor_72b8f33b242f205b = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xdd, 0xda, 0x36, 0x64, 0x52, 0x51, 0xb6, 0x82, 0xa1, 0x78, 0x78, 0x04, 0x0f, 0x41,
	0x9a, 0x1d, 0x49, 0x2f, 0x2a, 0x0a, 0x6e, 0x0a, 0x4a, 0x89, 0x48, 0x49, 0x7b, 0xb1, 0xa2, 0x30,
	0x99, 0x9d, 0x6c, 0x1e, 0xd9, 0x9d, 0xb7, 0xcc, 0x4c, 0x2c, 0x51, 0xfc, 0x07, 0xbc, 0x08, 0xb9,
	0xfa, 0x1f, 0xf8, 0xe7, 0xf4, 0x0f, 0x11, 0x3c, 0x78, 0xf3, 0x20, 0xbb, 0xc9, 0xb6, 0x9b, 0x90,
	0xbd, 0x2c, 0xf3, 0x7d, 0x9f, 0xef, 0xbc, 0xe1, 0xfd, 0x60, 0x4d, 0x24, 0x3e, 0x4a, 0x30, 0x1e,
	0x3b, 0xcb, 0xa5, 0x30, 0x06, 0x95, 0x09, 0x32, 0x43, 0x8e, 0xfc, 0xda, 0x52, 0x3e, 0x78, 0x88,
	0xc4, 0x25, 0xa5, 0x29, 0x69, 0x9e, 0x90, 0x14, 0x09, 0x7e, 0x11, 0x0e, 0x49, 0x2f, 0xb0, 0x83,
	0xc3, 0xe2, 0x27, 0x3b, 0xb1, 0xd2, 0x1d, 0x7b, 0x29, 0xe2, 0x58, 0x19, 0x4e, 0x59, 0x0e, 0x58,
	0x2e, 0xb4, 0x26, 0x57, 0xc0, 0x76, 0x41, 0xb7, 0xfe, 0x6d, 0xb3, 0xda, 0xf1, 0x22, 0x8d, 0xff,
	0x92, 0xdd, 0x45, 0xe1, 0xc4, 0xf2, 0x78, 0x4c, 0x91, 0x6a, 0x7a, 0xe0, 0xb5, 0xeb, 0xbd, 0xfd,
	0x79, 0x78, 0xef, 0xbb, 0x77, 0xe7, 0xd3, 0x87, 0xb0, 0x73, 0xf1, 0xa4, 0xf3, 0xec, 0xe3, 0xd7,
	0xee, 0xb7, 0x47, 0x83, 0x75, 0xb6, 0xb0, 0x4b, 0x41, 0x55, 0xfb, 0xd6, 0x66, 0xfb, 0x51, 0x61,
	0x5f, 0x65, 0xfd, 0x17, 0xcc, 0xcf, 0x6f, 0x0c, 0xa5, 0xa4, 0xa9, 0x76, 0xa8, 0xe3, 0xe2, 0x86,
	0xdb, 0xe0, 0xb5, 0x77, 0x7a, 0x7b, 0xf3, 0xb0, 0xfe, 0xfe, 0x56, 0xfe, 0x3d, 0xfd, 0xf1, 0x6a,
	0xb0, 0x81, 0xf3, 0x81, 0x35, 0x04, 0x9a, 0x04, 0xb5, 0x7a, 0x27, 0x52, 0xd5, 0xdc, 0xce, 0x13,
	0x0f, 0xaa, 0x92, 0xff, 0x86, 0xdd, 0x5f, 0x56, 0x4b, 0x45, 0x61, 0x05, 0xdd, 0x01, 0xaf, 0xdd,
	0xe8, 0xee, 0x07, 0x48, 0xc1, 0xdb, 0x32, 0x7e, 0xe6, 0x0c, 0xea, 0x78, 0xb0, 0xd1, 0xe0, 0x07,
	0xcc, 0xcf, 0x84, 0xb5, 0xe7, 0xb3, 0x4c, 0x9d, 0x44, 0x4a, 0x3b, 0x1c, 0xa1, 0x32, 0xcd, 0xdd,
	0x22, 0xe3, 0x86, 0x88, 0xdf, 0x62, 0x7b, 0x99, 0xc1, 0xcf, 0xc2, 0xa9, 0xbe, 0x9a, 0x9d, 0x44,
	0xcd, 0x5a, 0x41, 0xae, 0x68, 0xcf, 0xff, 0x7a, 0xf3, 0xf0, 0x8f, 0xc7, 0x7e, 0x7b, 0x8f, 0xcb,
	0x6e, 0x74, 0x7f, 0x7a, 0x21, 0x2c, 0x07, 0x00, 0x8c, 0x92, 0x64, 0x22, 0x90, 0xa4, 0x9d, 0x40,
	0x6d, 0x21, 0x52, 0x4e, 0x60, 0x62, 0x81, 0x46, 0xe0, 0xc6, 0xaa, 0xc4, 0x02, 0x80, 0xf3, 0x9b,
	0x13, 0x5c, 0x62, 0x92, 0xc0, 0x50, 0x15, 0x04, 0x5a, 0x3b, 0x55, 0xa6, 0xe4, 0x87, 0x24, 0x4c,
	0x84, 0x3a, 0x86, 0xfc, 0xb5, 0x20, 0x74, 0x04, 0xa9, 0x98, 0xe5, 0xe8, 0xd4, 0xaa, 0x08, 0x84,
	0x05, 0x01, 0xa9, 0x30, 0x13, 0x95, 0x57, 0x17, 0xc8, 0x00, 0x65, 0xca, 0x88, 0xe2, 0x50, 0x66,
	0xba, 0xf2, 0xd6, 0xdb, 0x9f, 0x2b, 0xab, 0x1d, 0xbd, 0xf2, 0xfc, 0xb2, 0xfe, 0xd3, 0x54, 0x19,
	0x94, 0x0b, 0xb1, 0xda, 0x94, 0x16, 0x67, 0x8d, 0xea, 0x0c, 0x00, 0x6b, 0xc8, 0xf5, 0xe9, 0x1b,
	0x54, 0xa5, 0xde, 0x19, 0x7b, 0x80, 0x14, 0xe4, 0xef, 0x9e, 0xa0, 0x0b, 0x4e, 0xfb, 0x48, 0xc1,
	0xeb, 0xc5, 0x5a, 0x5c, 0x1c, 0x5a, 0x27, 0xec, 0xf8, 0x3a, 0x26, 0x29, 0xe5, 0x48, 0x3c, 0xa5,
	0x48, 0x25, 0xdc, 0x46, 0x13, 0x1e, 0x13, 0xbf, 0xd9, 0xad, 0x5f, 0x5b, 0xf5, 0xd3, 0xfe, 0xd2,
	0x39, 0xdc, 0x2d, 0x76, 0xe1, 0xe8, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xf7, 0xba, 0x53,
	0x7c, 0x03, 0x00, 0x00,
}
