// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/user/user.proto

package io

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NewUserResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CompanyId            string   `protobuf:"bytes,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	EmailAddress         string   `protobuf:"bytes,4,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	Secret               string   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	Certificate          string   `protobuf:"bytes,6,opt,name=certificate,proto3" json:"certificate,omitempty"`
	PrivateKey           string   `protobuf:"bytes,7,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserResponse) Reset()         { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()    {}
func (*NewUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{0}
}

func (m *NewUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserResponse.Unmarshal(m, b)
}
func (m *NewUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserResponse.Marshal(b, m, deterministic)
}
func (m *NewUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserResponse.Merge(m, src)
}
func (m *NewUserResponse) XXX_Size() int {
	return xxx_messageInfo_NewUserResponse.Size(m)
}
func (m *NewUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserResponse proto.InternalMessageInfo

func (m *NewUserResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NewUserResponse) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *NewUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewUserResponse) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *NewUserResponse) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *NewUserResponse) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *NewUserResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type NewUser struct {
	// @tag: validateCreate:"required,alphanum,gte=6"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" validateCreate:"required,alphanum,gte=6"`
	// @tag: validateCreate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" validateCreate:"required"`
	// @tag: validateCreate:"required"
	EmailAddress string `protobuf:"bytes,3,opt,name=emailAddress,proto3" json:"emailAddress,omitempty" validateCreate:"required"`
	// @tag: validateCreate:"omitempty"
	SendEmail bool `protobuf:"varint,4,opt,name=sendEmail,proto3" json:"sendEmail,omitempty" validateCreate:"omitempty"`
	// If a company name is provided, then a new company will be created, otherwise, the user will inherit the company
	// of the parent.  Only system admin users can create new companies.
	// @tag: validateCreate:"omitempty"
	CompanyName string `protobuf:"bytes,5,opt,name=companyName,proto3" json:"companyName,omitempty" validateCreate:"omitempty"`
	// @tag: validateCreate:"omitempty"
	MobileNumber string `protobuf:"bytes,6,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty" validateCreate:"omitempty"`
	// @tag: validateCreate:"omitempty"
	TwoFactorAuth bool `protobuf:"varint,7,opt,name=twoFactorAuth,proto3" json:"twoFactorAuth,omitempty" validateCreate:"omitempty"`
	// @tag: validateCreate:"omitempty"
	IsOwner bool `protobuf:"varint,8,opt,name=isOwner,proto3" json:"isOwner,omitempty" validateCreate:"omitempty"`
	// @tag: validateCreate:"omitempty"
	ReadOnly             bool     `protobuf:"varint,9,opt,name=readOnly,proto3" json:"readOnly,omitempty" validateCreate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUser) Reset()         { *m = NewUser{} }
func (m *NewUser) String() string { return proto.CompactTextString(m) }
func (*NewUser) ProtoMessage()    {}
func (*NewUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{1}
}

func (m *NewUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUser.Unmarshal(m, b)
}
func (m *NewUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUser.Marshal(b, m, deterministic)
}
func (m *NewUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUser.Merge(m, src)
}
func (m *NewUser) XXX_Size() int {
	return xxx_messageInfo_NewUser.Size(m)
}
func (m *NewUser) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUser.DiscardUnknown(m)
}

var xxx_messageInfo_NewUser proto.InternalMessageInfo

func (m *NewUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *NewUser) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *NewUser) GetSendEmail() bool {
	if m != nil {
		return m.SendEmail
	}
	return false
}

func (m *NewUser) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *NewUser) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

func (m *NewUser) GetTwoFactorAuth() bool {
	if m != nil {
		return m.TwoFactorAuth
	}
	return false
}

func (m *NewUser) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

func (m *NewUser) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

// response used by getUser
type GetUserResponse struct {
	Email                string               `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	CompanyId            string               `protobuf:"bytes,3,opt,name=companyId,proto3" json:"companyId,omitempty"`
	CompanyName          string               `protobuf:"bytes,4,opt,name=companyName,proto3" json:"companyName,omitempty"`
	CompanyStatus        uint64               `protobuf:"varint,5,opt,name=companyStatus,proto3" json:"companyStatus,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{2}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserResponse) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *GetUserResponse) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *GetUserResponse) GetCompanyStatus() uint64 {
	if m != nil {
		return m.CompanyStatus
	}
	return 0
}

func (m *GetUserResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type BillingMeta struct {
	CompanyName          string   `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	TaxId                string   `protobuf:"bytes,2,opt,name=taxId,proto3" json:"taxId,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	BillingAddress       *Address `protobuf:"bytes,4,opt,name=billingAddress,proto3" json:"billingAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillingMeta) Reset()         { *m = BillingMeta{} }
func (m *BillingMeta) String() string { return proto.CompactTextString(m) }
func (*BillingMeta) ProtoMessage()    {}
func (*BillingMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{3}
}

func (m *BillingMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingMeta.Unmarshal(m, b)
}
func (m *BillingMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingMeta.Marshal(b, m, deterministic)
}
func (m *BillingMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingMeta.Merge(m, src)
}
func (m *BillingMeta) XXX_Size() int {
	return xxx_messageInfo_BillingMeta.Size(m)
}
func (m *BillingMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BillingMeta proto.InternalMessageInfo

func (m *BillingMeta) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *BillingMeta) GetTaxId() string {
	if m != nil {
		return m.TaxId
	}
	return ""
}

func (m *BillingMeta) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *BillingMeta) GetBillingAddress() *Address {
	if m != nil {
		return m.BillingAddress
	}
	return nil
}

type Credentials struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{4}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type JWT struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWT) Reset()         { *m = JWT{} }
func (m *JWT) String() string { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()    {}
func (*JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{5}
}

func (m *JWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWT.Unmarshal(m, b)
}
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWT.Marshal(b, m, deterministic)
}
func (m *JWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWT.Merge(m, src)
}
func (m *JWT) XXX_Size() int {
	return xxx_messageInfo_JWT.Size(m)
}
func (m *JWT) XXX_DiscardUnknown() {
	xxx_messageInfo_JWT.DiscardUnknown(m)
}

var xxx_messageInfo_JWT proto.InternalMessageInfo

func (m *JWT) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PasswordResetInput struct {
	// @tag: validateUpdate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateUpdate:"required"`
	// @tag: validateUpdate:"required"
	RegisteredEmail string `protobuf:"bytes,2,opt,name=registeredEmail,proto3" json:"registeredEmail,omitempty" validateUpdate:"required"`
	// @tag: validateUpdate:"required"
	NewPassword string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty" validateUpdate:"required"`
	// @tag: validateUpdate:"required"
	ConfirmNewPassword   string   `protobuf:"bytes,4,opt,name=confirmNewPassword,proto3" json:"confirmNewPassword,omitempty" validateUpdate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordResetInput) Reset()         { *m = PasswordResetInput{} }
func (m *PasswordResetInput) String() string { return proto.CompactTextString(m) }
func (*PasswordResetInput) ProtoMessage()    {}
func (*PasswordResetInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{6}
}

func (m *PasswordResetInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordResetInput.Unmarshal(m, b)
}
func (m *PasswordResetInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordResetInput.Marshal(b, m, deterministic)
}
func (m *PasswordResetInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordResetInput.Merge(m, src)
}
func (m *PasswordResetInput) XXX_Size() int {
	return xxx_messageInfo_PasswordResetInput.Size(m)
}
func (m *PasswordResetInput) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordResetInput.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordResetInput proto.InternalMessageInfo

func (m *PasswordResetInput) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PasswordResetInput) GetRegisteredEmail() string {
	if m != nil {
		return m.RegisteredEmail
	}
	return ""
}

func (m *PasswordResetInput) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *PasswordResetInput) GetConfirmNewPassword() string {
	if m != nil {
		return m.ConfirmNewPassword
	}
	return ""
}

type Username struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_9480a723983e579d, []int{7}
}

func (m *Username) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Username.Unmarshal(m, b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Username.Marshal(b, m, deterministic)
}
func (m *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(m, src)
}
func (m *Username) XXX_Size() int {
	return xxx_messageInfo_Username.Size(m)
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*NewUserResponse)(nil), "io.NewUserResponse")
	proto.RegisterType((*NewUser)(nil), "io.NewUser")
	proto.RegisterType((*GetUserResponse)(nil), "io.GetUserResponse")
	proto.RegisterType((*BillingMeta)(nil), "io.BillingMeta")
	proto.RegisterType((*Credentials)(nil), "io.Credentials")
	proto.RegisterType((*JWT)(nil), "io.JWT")
	proto.RegisterType((*PasswordResetInput)(nil), "io.PasswordResetInput")
	proto.RegisterType((*Username)(nil), "io.Username")
}

func init() {
	proto.RegisterFile("io/user/user.proto", fileDescriptor_9480a723983e579d)
}

var fileDescriptor_9480a723983e579d = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6a, 0xdb, 0x4a,
	0x14, 0x45, 0xb2, 0x9d, 0xd8, 0xd7, 0x2f, 0x31, 0x0c, 0xe1, 0x21, 0xfc, 0xc2, 0x7b, 0x41, 0x84,
	0x47, 0x56, 0x12, 0x24, 0x9b, 0xae, 0x0a, 0x49, 0x49, 0x8b, 0x1b, 0xea, 0x04, 0x35, 0xa1, 0xd0,
	0xdd, 0x58, 0xba, 0x76, 0x86, 0x48, 0x73, 0xc5, 0xcc, 0xb8, 0x6e, 0xfe, 0xa4, 0xcb, 0x6e, 0xba,
	0xe9, 0x5f, 0x95, 0xfe, 0x48, 0xd1, 0x48, 0xb2, 0x2d, 0x25, 0xdd, 0x74, 0x63, 0x38, 0xe7, 0x5c,
	0x7c, 0xef, 0x39, 0x3a, 0x03, 0x4c, 0x50, 0xb8, 0xd4, 0xa8, 0xec, 0x4f, 0x90, 0x2b, 0x32, 0xc4,
	0x5c, 0x41, 0xe3, 0xff, 0x16, 0x44, 0x8b, 0x14, 0x43, 0xcb, 0xcc, 0x96, 0xf3, 0xd0, 0x88, 0x0c,
	0xb5, 0xe1, 0x59, 0x5e, 0x0e, 0x8d, 0x3d, 0x41, 0x61, 0x4c, 0x59, 0x46, 0x32, 0xcc, 0x51, 0x69,
	0x92, 0x3c, 0x2d, 0x15, 0xff, 0x87, 0x03, 0xa3, 0x29, 0xae, 0xee, 0x34, 0xaa, 0x08, 0x75, 0x4e,
	0x52, 0x23, 0xfb, 0x1b, 0x76, 0x8a, 0x05, 0x93, 0xc4, 0x73, 0x8e, 0x9c, 0x93, 0x41, 0x54, 0x21,
	0x76, 0x08, 0x83, 0x98, 0xb2, 0x9c, 0xcb, 0xc7, 0x49, 0xe2, 0xb9, 0x56, 0xda, 0x10, 0x6c, 0x0c,
	0xfd, 0x62, 0x4e, 0xf2, 0x0c, 0xbd, 0x8e, 0x15, 0xd7, 0x98, 0xf9, 0xf0, 0x17, 0x66, 0x5c, 0xa4,
	0xe7, 0x49, 0xa2, 0x50, 0x6b, 0xaf, 0x6b, 0xf5, 0x06, 0x57, 0x6c, 0xd5, 0x18, 0x2b, 0x34, 0x5e,
	0xaf, 0xdc, 0x5a, 0x22, 0x76, 0x04, 0xc3, 0x18, 0x95, 0x11, 0x73, 0x11, 0x73, 0x83, 0xde, 0x8e,
	0x15, 0xb7, 0x29, 0xf6, 0x2f, 0x40, 0xae, 0xc4, 0x27, 0x6e, 0xf0, 0x0a, 0x1f, 0xbd, 0x5d, 0x3b,
	0xb0, 0xc5, 0xf8, 0xdf, 0x5c, 0xd8, 0xad, 0x3c, 0x36, 0xae, 0x74, 0x5a, 0x57, 0x8e, 0xa1, 0x9f,
	0x73, 0xad, 0x57, 0xa4, 0x6a, 0x7b, 0x6b, 0xfc, 0xc4, 0x41, 0xe7, 0x19, 0x07, 0x87, 0x30, 0xd0,
	0x28, 0x93, 0xcb, 0x82, 0xb3, 0x16, 0xfb, 0xd1, 0x86, 0xb0, 0x3e, 0xca, 0xb0, 0xa6, 0xc5, 0xf2,
	0x5e, 0xe5, 0x63, 0x43, 0x15, 0x3b, 0x32, 0x9a, 0x89, 0x14, 0xa7, 0xcb, 0x6c, 0x86, 0xaa, 0xb2,
	0xda, 0xe0, 0xd8, 0x31, 0xec, 0x99, 0x15, 0xbd, 0xe6, 0xb1, 0x21, 0x75, 0xbe, 0x34, 0xf7, 0xd6,
	0x6e, 0x3f, 0x6a, 0x92, 0xcc, 0x83, 0x5d, 0xa1, 0xaf, 0x57, 0x12, 0x95, 0xd7, 0xb7, 0x7a, 0x0d,
	0x0b, 0x8f, 0x0a, 0x79, 0x72, 0x2d, 0xd3, 0x47, 0x6f, 0x60, 0xa5, 0x35, 0xf6, 0x7f, 0x3a, 0x30,
	0x7a, 0x83, 0xa6, 0xd1, 0x85, 0x03, 0xe8, 0x59, 0x8f, 0x55, 0x58, 0x25, 0x68, 0xa4, 0xe8, 0xb6,
	0x52, 0x6c, 0xb4, 0xa4, 0xd3, 0x6e, 0x49, 0x2b, 0x85, 0xee, 0xd3, 0x14, 0x8e, 0x61, 0xaf, 0x82,
	0xef, 0x0d, 0x37, 0x4b, 0x6d, 0x93, 0xea, 0x46, 0x4d, 0x92, 0xbd, 0x80, 0x41, 0xac, 0x90, 0x1b,
	0x4c, 0xce, 0x8d, 0x0d, 0x6a, 0x78, 0x3a, 0x0e, 0xca, 0x67, 0x10, 0xd4, 0xcf, 0x20, 0xb8, 0xad,
	0x9f, 0x41, 0xb4, 0x19, 0xf6, 0xbf, 0x38, 0x30, 0xbc, 0x10, 0x69, 0x2a, 0xe4, 0xe2, 0x1d, 0x1a,
	0xde, 0xbe, 0xc8, 0x79, 0x7a, 0xd1, 0x01, 0xf4, 0x0c, 0xff, 0xbc, 0xee, 0x7c, 0x09, 0x0a, 0x9f,
	0x0a, 0xe7, 0xa8, 0x50, 0xc6, 0x75, 0xe1, 0x37, 0x04, 0x3b, 0x83, 0xfd, 0x59, 0xb9, 0x64, 0xbb,
	0xf3, 0xc3, 0xd3, 0x61, 0x20, 0x28, 0xa8, 0xa8, 0xa8, 0x35, 0xe2, 0x5f, 0xc2, 0xf0, 0x95, 0xc2,
	0x04, 0xa5, 0x11, 0x3c, 0xd5, 0x7f, 0xda, 0x55, 0xff, 0x1f, 0xe8, 0xbc, 0xfd, 0x70, 0x6b, 0xcf,
	0xa6, 0x07, 0x94, 0xf5, 0xa7, 0xb3, 0xc0, 0xff, 0xea, 0x00, 0xbb, 0xa9, 0x26, 0x23, 0xd4, 0x68,
	0x26, 0x32, 0x5f, 0x1a, 0xb6, 0x0f, 0xae, 0xa8, 0xdf, 0xbb, 0x2b, 0x12, 0x76, 0x02, 0x23, 0x85,
	0x0b, 0xa1, 0x0d, 0x2a, 0xac, 0x1a, 0x5d, 0xae, 0x69, 0xd3, 0x45, 0x7e, 0x12, 0x57, 0xf5, 0x5f,
	0x56, 0x49, 0x6c, 0x53, 0x2c, 0x00, 0x16, 0x93, 0x9c, 0x0b, 0x95, 0x4d, 0xb7, 0x06, 0xcb, 0x4f,
	0xff, 0x8c, 0xe2, 0xff, 0x0f, 0xfd, 0xbb, 0x2d, 0x9f, 0xbf, 0xcb, 0xe0, 0xe2, 0x25, 0x8c, 0x04,
	0x05, 0x85, 0xed, 0x07, 0x61, 0x82, 0x9b, 0x2b, 0x41, 0x1f, 0x8f, 0xb5, 0xe1, 0xfa, 0x7e, 0xcd,
	0xc5, 0x94, 0x85, 0x82, 0xc2, 0x8c, 0x12, 0x4c, 0x43, 0x9d, 0x3c, 0x84, 0x0b, 0x0a, 0x05, 0x7d,
	0x77, 0xbb, 0x37, 0x57, 0x13, 0x9a, 0xed, 0xd8, 0xa2, 0x9c, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0xfe, 0xf8, 0xfd, 0x00, 0x57, 0x05, 0x00, 0x00,
}
