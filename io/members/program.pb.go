// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/member/program.proto

package members

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProgramAutoUpDownGradeTierType int32

const (
	// Auto up / down grades based on points.
	ProgramAutoUpDownGradeTierType_UPDOWNGRADE_POINTS ProgramAutoUpDownGradeTierType = 0
	// Auto up / down grades based on secondary points.
	ProgramAutoUpDownGradeTierType_UPDOWNGRADE_SECONDARY_POINTS ProgramAutoUpDownGradeTierType = 1
	// Auto up / down grades based on tier points.
	ProgramAutoUpDownGradeTierType_UPDOWNGRADE_TIER_POINTS ProgramAutoUpDownGradeTierType = 2
	// Auto up / down grades based on # of visits.
	ProgramAutoUpDownGradeTierType_UPDOWNGRADE_VISITS ProgramAutoUpDownGradeTierType = 3
)

var ProgramAutoUpDownGradeTierType_name = map[int32]string{
	0: "UPDOWNGRADE_POINTS",
	1: "UPDOWNGRADE_SECONDARY_POINTS",
	2: "UPDOWNGRADE_TIER_POINTS",
	3: "UPDOWNGRADE_VISITS",
}

var ProgramAutoUpDownGradeTierType_value = map[string]int32{
	"UPDOWNGRADE_POINTS":           0,
	"UPDOWNGRADE_SECONDARY_POINTS": 1,
	"UPDOWNGRADE_TIER_POINTS":      2,
	"UPDOWNGRADE_VISITS":           3,
}

func (x ProgramAutoUpDownGradeTierType) String() string {
	return proto.EnumName(ProgramAutoUpDownGradeTierType_name, int32(x))
}

func (ProgramAutoUpDownGradeTierType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{0}
}

type BalanceType int32

const (
	BalanceType_BALANCE_TYPE_STRING BalanceType = 0
	BalanceType_BALANCE_TYPE_INT    BalanceType = 1
	BalanceType_BALANCE_TYPE_DOUBLE BalanceType = 2
	BalanceType_BALANCE_TYPE_MONEY  BalanceType = 3
)

var BalanceType_name = map[int32]string{
	0: "BALANCE_TYPE_STRING",
	1: "BALANCE_TYPE_INT",
	2: "BALANCE_TYPE_DOUBLE",
	3: "BALANCE_TYPE_MONEY",
}

var BalanceType_value = map[string]int32{
	"BALANCE_TYPE_STRING": 0,
	"BALANCE_TYPE_INT":    1,
	"BALANCE_TYPE_DOUBLE": 2,
	"BALANCE_TYPE_MONEY":  3,
}

func (x BalanceType) String() string {
	return proto.EnumName(BalanceType_name, int32(x))
}

func (BalanceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{1}
}

type ProfileImageSetting int32

const (
	ProfileImageSetting_PROFILE_IMAGE_NONE     ProfileImageSetting = 0
	ProfileImageSetting_PROFILE_IMAGE_OPTIONAL ProfileImageSetting = 1
	ProfileImageSetting_PROFILE_IMAGE_REQUIRED ProfileImageSetting = 2
)

var ProfileImageSetting_name = map[int32]string{
	0: "PROFILE_IMAGE_NONE",
	1: "PROFILE_IMAGE_OPTIONAL",
	2: "PROFILE_IMAGE_REQUIRED",
}

var ProfileImageSetting_value = map[string]int32{
	"PROFILE_IMAGE_NONE":     0,
	"PROFILE_IMAGE_OPTIONAL": 1,
	"PROFILE_IMAGE_REQUIRED": 2,
}

func (x ProfileImageSetting) String() string {
	return proto.EnumName(ProfileImageSetting_name, int32(x))
}

func (ProfileImageSetting) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{2}
}

// Pass fields that are protocol specific.
type ProtocolPassFields int32

const (
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_DO_NOT_USE           ProtocolPassFields = 0
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_NUMBER               ProtocolPassFields = 1
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_DETAILS              ProtocolPassFields = 2
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_TIER_NAME            ProtocolPassFields = 3
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_TIER_INFORMATION     ProtocolPassFields = 4
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_POINTS               ProtocolPassFields = 5
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_SECONDARY_POINTS     ProtocolPassFields = 6
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_GENERIC_PROGRAM_INFO ProtocolPassFields = 7
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_LOCATIONS            ProtocolPassFields = 8
	ProtocolPassFields_PROTOCOL_MEMBERSHIP_LATEST_TRANSACTIONS  ProtocolPassFields = 9
)

var ProtocolPassFields_name = map[int32]string{
	0: "PROTOCOL_MEMBERSHIP_DO_NOT_USE",
	1: "PROTOCOL_MEMBERSHIP_NUMBER",
	2: "PROTOCOL_MEMBERSHIP_DETAILS",
	3: "PROTOCOL_MEMBERSHIP_TIER_NAME",
	4: "PROTOCOL_MEMBERSHIP_TIER_INFORMATION",
	5: "PROTOCOL_MEMBERSHIP_POINTS",
	6: "PROTOCOL_MEMBERSHIP_SECONDARY_POINTS",
	7: "PROTOCOL_MEMBERSHIP_GENERIC_PROGRAM_INFO",
	8: "PROTOCOL_MEMBERSHIP_LOCATIONS",
	9: "PROTOCOL_MEMBERSHIP_LATEST_TRANSACTIONS",
}

var ProtocolPassFields_value = map[string]int32{
	"PROTOCOL_MEMBERSHIP_DO_NOT_USE":           0,
	"PROTOCOL_MEMBERSHIP_NUMBER":               1,
	"PROTOCOL_MEMBERSHIP_DETAILS":              2,
	"PROTOCOL_MEMBERSHIP_TIER_NAME":            3,
	"PROTOCOL_MEMBERSHIP_TIER_INFORMATION":     4,
	"PROTOCOL_MEMBERSHIP_POINTS":               5,
	"PROTOCOL_MEMBERSHIP_SECONDARY_POINTS":     6,
	"PROTOCOL_MEMBERSHIP_GENERIC_PROGRAM_INFO": 7,
	"PROTOCOL_MEMBERSHIP_LOCATIONS":            8,
	"PROTOCOL_MEMBERSHIP_LATEST_TRANSACTIONS":  9,
}

func (x ProtocolPassFields) String() string {
	return proto.EnumName(ProtocolPassFields_name, int32(x))
}

func (ProtocolPassFields) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{3}
}

type PointsType struct {
	// The balance type.
	BalanceType BalanceType `protobuf:"varint,1,opt,name=balanceType,proto3,enum=members.BalanceType" json:"balanceType,omitempty"`
	// The currency code; only needs to be provided if balance type equals BALANCE_TYPE_MONEY.
	CurrencyCode         string   `protobuf:"bytes,2,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PointsType) Reset()         { *m = PointsType{} }
func (m *PointsType) String() string { return proto.CompactTextString(m) }
func (*PointsType) ProtoMessage()    {}
func (*PointsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{0}
}

func (m *PointsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PointsType.Unmarshal(m, b)
}
func (m *PointsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PointsType.Marshal(b, m, deterministic)
}
func (m *PointsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PointsType.Merge(m, src)
}
func (m *PointsType) XXX_Size() int {
	return xxx_messageInfo_PointsType.Size(m)
}
func (m *PointsType) XXX_DiscardUnknown() {
	xxx_messageInfo_PointsType.DiscardUnknown(m)
}

var xxx_messageInfo_PointsType proto.InternalMessageInfo

func (m *PointsType) GetBalanceType() BalanceType {
	if m != nil {
		return m.BalanceType
	}
	return BalanceType_BALANCE_TYPE_STRING
}

func (m *PointsType) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

// The Program Details
type Program struct {
	// PassKit generated program id (22 characters).
	// @tag: validateGeneric:"required" validateCreate:"-" validateUpdate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateGeneric:"required" validateCreate:"-" validateUpdate:"required"`
	// Name of the membership programs; will be shown on the enrolment page (if set) and issuer name.
	// @tag: validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validateGeneric:"omitempty" validateCreate:"required" validateUpdate:"required"`
	// Localized name of the membership program.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedName *io.LocalizedString `protobuf:"bytes,3,opt,name=localizedName,proto3" json:"localizedName,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The program status (bitmask of ProjectStatus). Defaults to io.ProjectStatus.PROJECT_ACTIVE_FOR_OBJECT_CREATION + io.ProjectStatus.PROJECT_DRAFT.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Status []io.ProjectStatus `protobuf:"varint,4,rep,packed,name=status,proto3,enum=io.ProjectStatus" json:"status,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// The project quota.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Quota *io.Quota `protobuf:"bytes,5,opt,name=quota,proto3" json:"quota,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Leave empty for draft programs. Needs to be set for programs where status contains PROJECT_PUBLISHED.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"-"
	PassTypeIdentifier string `protobuf:"bytes,6,opt,name=passTypeIdentifier,proto3" json:"passTypeIdentifier,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"-"`
	// Contains the email & sms distribution settings for the program.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	DistributionSettings *io.DistributionSettings `protobuf:"bytes,7,opt,name=distributionSettings,proto3" json:"distributionSettings,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Auto deletes the member after `autoDeleteDaysAfterExpiry` days.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	AutoDeleteDaysAfterExpiry uint32 `protobuf:"varint,8,opt,name=autoDeleteDaysAfterExpiry,proto3" json:"autoDeleteDaysAfterExpiry,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Sets how many days before expiry to send a reminder lockscreen message to the member.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	AutoReminderDaysBeforeExpiry uint32 `protobuf:"varint,9,opt,name=autoReminderDaysBeforeExpiry,proto3" json:"autoReminderDaysBeforeExpiry,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The message to send to the user `autoReminderDaysBeforeExpiry` days before expiry.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	ExpiryMessage string `protobuf:"bytes,10,opt,name=expiryMessage,proto3" json:"expiryMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Localized expiry message.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	LocalizedExpiryMessage *io.LocalizedString `protobuf:"bytes,11,opt,name=localizedExpiryMessage,proto3" json:"localizedExpiryMessage,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Auto deletes the member `autoDeleteDaysAfterNotInstalling` days after the card has not been installed.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	AutoDeleteDaysAfterNotInstalling uint32 `protobuf:"varint,12,opt,name=autoDeleteDaysAfterNotInstalling,proto3" json:"autoDeleteDaysAfterNotInstalling,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Callbacks that are defined on program events. The key is the CallbackEvent ENUM.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	Callbacks map[uint32]*io.Callback `protobuf:"bytes,14,rep,name=callbacks,proto3" json:"callbacks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Points format for primary points.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	PointsType *PointsType `protobuf:"bytes,16,opt,name=pointsType,proto3" json:"pointsType,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Points format for secondary points.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	SecondaryPointsType *PointsType `protobuf:"bytes,17,opt,name=secondaryPointsType,proto3" json:"secondaryPointsType,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// Issued, installed, uninstalled and invalidated counts. Metrics are not writable.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Metrics *io.Metrics `protobuf:"bytes,19,opt,name=metrics,proto3" json:"metrics,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// A list of dynamic fields which can be used to filter members.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	MetaFieldsList []string `protobuf:"bytes,20,rep,name=metaFieldsList,proto3" json:"metaFieldsList,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	// The date the program was created.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Created *timestamp.Timestamp `protobuf:"bytes,21,opt,name=created,proto3" json:"created,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// The date the program was updated.
	// @tag: validateGeneric:"-" validateCreate:"-" validateUpdate:"-"
	Updated *timestamp.Timestamp `protobuf:"bytes,22,opt,name=updated,proto3" json:"updated,omitempty" validateGeneric:"-" validateCreate:"-" validateUpdate:"-"`
	// Indicates if the program requires member profile images on enrol.
	// @tag: validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"
	ProfileImageSettings ProfileImageSetting `protobuf:"varint,23,opt,name=profileImageSettings,proto3,enum=members.ProfileImageSetting" json:"profileImageSettings,omitempty" validateGeneric:"omitempty" validateCreate:"omitempty" validateUpdate:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Program) Reset()         { *m = Program{} }
func (m *Program) String() string { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()    {}
func (*Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{1}
}

func (m *Program) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Program.Unmarshal(m, b)
}
func (m *Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Program.Marshal(b, m, deterministic)
}
func (m *Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Program.Merge(m, src)
}
func (m *Program) XXX_Size() int {
	return xxx_messageInfo_Program.Size(m)
}
func (m *Program) XXX_DiscardUnknown() {
	xxx_messageInfo_Program.DiscardUnknown(m)
}

var xxx_messageInfo_Program proto.InternalMessageInfo

func (m *Program) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Program) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Program) GetLocalizedName() *io.LocalizedString {
	if m != nil {
		return m.LocalizedName
	}
	return nil
}

func (m *Program) GetStatus() []io.ProjectStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Program) GetQuota() *io.Quota {
	if m != nil {
		return m.Quota
	}
	return nil
}

func (m *Program) GetPassTypeIdentifier() string {
	if m != nil {
		return m.PassTypeIdentifier
	}
	return ""
}

func (m *Program) GetDistributionSettings() *io.DistributionSettings {
	if m != nil {
		return m.DistributionSettings
	}
	return nil
}

func (m *Program) GetAutoDeleteDaysAfterExpiry() uint32 {
	if m != nil {
		return m.AutoDeleteDaysAfterExpiry
	}
	return 0
}

func (m *Program) GetAutoReminderDaysBeforeExpiry() uint32 {
	if m != nil {
		return m.AutoReminderDaysBeforeExpiry
	}
	return 0
}

func (m *Program) GetExpiryMessage() string {
	if m != nil {
		return m.ExpiryMessage
	}
	return ""
}

func (m *Program) GetLocalizedExpiryMessage() *io.LocalizedString {
	if m != nil {
		return m.LocalizedExpiryMessage
	}
	return nil
}

func (m *Program) GetAutoDeleteDaysAfterNotInstalling() uint32 {
	if m != nil {
		return m.AutoDeleteDaysAfterNotInstalling
	}
	return 0
}

func (m *Program) GetCallbacks() map[uint32]*io.Callback {
	if m != nil {
		return m.Callbacks
	}
	return nil
}

func (m *Program) GetPointsType() *PointsType {
	if m != nil {
		return m.PointsType
	}
	return nil
}

func (m *Program) GetSecondaryPointsType() *PointsType {
	if m != nil {
		return m.SecondaryPointsType
	}
	return nil
}

func (m *Program) GetMetrics() *io.Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Program) GetMetaFieldsList() []string {
	if m != nil {
		return m.MetaFieldsList
	}
	return nil
}

func (m *Program) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Program) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *Program) GetProfileImageSettings() ProfileImageSetting {
	if m != nil {
		return m.ProfileImageSettings
	}
	return ProfileImageSetting_PROFILE_IMAGE_NONE
}

// Contains payload for copying a program
type ProgramCopyRequest struct {
	// uuid of the program to copy.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// optional name for the new program.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Status to set the copied program to.
	Status []io.ProjectStatus `protobuf:"varint,3,rep,packed,name=status,proto3,enum=io.ProjectStatus" json:"status,omitempty"`
	// Can optionally be provided if setting status to published (requires a prod cert).
	PassTypeIdentifier   string   `protobuf:"bytes,4,opt,name=passTypeIdentifier,proto3" json:"passTypeIdentifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProgramCopyRequest) Reset()         { *m = ProgramCopyRequest{} }
func (m *ProgramCopyRequest) String() string { return proto.CompactTextString(m) }
func (*ProgramCopyRequest) ProtoMessage()    {}
func (*ProgramCopyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{2}
}

func (m *ProgramCopyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProgramCopyRequest.Unmarshal(m, b)
}
func (m *ProgramCopyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProgramCopyRequest.Marshal(b, m, deterministic)
}
func (m *ProgramCopyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProgramCopyRequest.Merge(m, src)
}
func (m *ProgramCopyRequest) XXX_Size() int {
	return xxx_messageInfo_ProgramCopyRequest.Size(m)
}
func (m *ProgramCopyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProgramCopyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProgramCopyRequest proto.InternalMessageInfo

func (m *ProgramCopyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProgramCopyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProgramCopyRequest) GetStatus() []io.ProjectStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ProgramCopyRequest) GetPassTypeIdentifier() string {
	if m != nil {
		return m.PassTypeIdentifier
	}
	return ""
}

// Contains an array of programs.
type ListProgramsResponse struct {
	// An array of programs.
	Programs             []*Program `protobuf:"bytes,1,rep,name=programs,proto3" json:"programs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListProgramsResponse) Reset()         { *m = ListProgramsResponse{} }
func (m *ListProgramsResponse) String() string { return proto.CompactTextString(m) }
func (*ListProgramsResponse) ProtoMessage()    {}
func (*ListProgramsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{3}
}

func (m *ListProgramsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProgramsResponse.Unmarshal(m, b)
}
func (m *ListProgramsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProgramsResponse.Marshal(b, m, deterministic)
}
func (m *ListProgramsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProgramsResponse.Merge(m, src)
}
func (m *ListProgramsResponse) XXX_Size() int {
	return xxx_messageInfo_ListProgramsResponse.Size(m)
}
func (m *ListProgramsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProgramsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProgramsResponse proto.InternalMessageInfo

func (m *ListProgramsResponse) GetPrograms() []*Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

type GetMetaKeysResponse struct {
	// An array of distinct meta keys.
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMetaKeysResponse) Reset()         { *m = GetMetaKeysResponse{} }
func (m *GetMetaKeysResponse) String() string { return proto.CompactTextString(m) }
func (*GetMetaKeysResponse) ProtoMessage()    {}
func (*GetMetaKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d9cd2d3e51b663d, []int{4}
}

func (m *GetMetaKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMetaKeysResponse.Unmarshal(m, b)
}
func (m *GetMetaKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMetaKeysResponse.Marshal(b, m, deterministic)
}
func (m *GetMetaKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMetaKeysResponse.Merge(m, src)
}
func (m *GetMetaKeysResponse) XXX_Size() int {
	return xxx_messageInfo_GetMetaKeysResponse.Size(m)
}
func (m *GetMetaKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMetaKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMetaKeysResponse proto.InternalMessageInfo

func (m *GetMetaKeysResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterEnum("members.ProgramAutoUpDownGradeTierType", ProgramAutoUpDownGradeTierType_name, ProgramAutoUpDownGradeTierType_value)
	proto.RegisterEnum("members.BalanceType", BalanceType_name, BalanceType_value)
	proto.RegisterEnum("members.ProfileImageSetting", ProfileImageSetting_name, ProfileImageSetting_value)
	proto.RegisterEnum("members.ProtocolPassFields", ProtocolPassFields_name, ProtocolPassFields_value)
	proto.RegisterType((*PointsType)(nil), "members.PointsType")
	proto.RegisterType((*Program)(nil), "members.Program")
	proto.RegisterMapType((map[uint32]*io.Callback)(nil), "members.Program.CallbacksEntry")
	proto.RegisterType((*ProgramCopyRequest)(nil), "members.ProgramCopyRequest")
	proto.RegisterType((*ListProgramsResponse)(nil), "members.ListProgramsResponse")
	proto.RegisterType((*GetMetaKeysResponse)(nil), "members.GetMetaKeysResponse")
}

func init() {
	proto.RegisterFile("io/member/program.proto", fileDescriptor_1d9cd2d3e51b663d)
}

var fileDescriptor_1d9cd2d3e51b663d = []byte{
	// 1272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdf, 0x72, 0xdb, 0xc4,
	0x17, 0xae, 0xe2, 0xfc, 0xf3, 0xa6, 0xc9, 0xa8, 0x9b, 0xfc, 0x12, 0xfd, 0xd2, 0xd0, 0x08, 0x4f,
	0x01, 0x37, 0xa4, 0xf6, 0x4c, 0xca, 0x30, 0xd0, 0x81, 0x0b, 0xc5, 0x56, 0x83, 0x12, 0x5b, 0x52,
	0x57, 0x0a, 0x4c, 0xb9, 0xf1, 0xac, 0xe5, 0x8d, 0xb3, 0x44, 0xd2, 0xaa, 0xda, 0x35, 0xc5, 0x3c,
	0x01, 0x57, 0x30, 0x03, 0x6f, 0xc1, 0xe3, 0xf4, 0x89, 0x18, 0xad, 0x24, 0xc7, 0x4e, 0x6c, 0xca,
	0x55, 0x36, 0xe7, 0xfb, 0xce, 0xb7, 0xbb, 0xe7, 0xec, 0xf9, 0x2c, 0xb0, 0x47, 0x59, 0x33, 0x22,
	0x51, 0x9f, 0xa4, 0xcd, 0x24, 0x65, 0xc3, 0x14, 0x47, 0x8d, 0x24, 0x65, 0x82, 0xc1, 0xb5, 0x3c,
	0xca, 0xf7, 0x0f, 0x87, 0x8c, 0x0d, 0x43, 0xd2, 0x94, 0xe1, 0xfe, 0xe8, 0xaa, 0x29, 0x68, 0x44,
	0xb8, 0xc0, 0x51, 0x92, 0x33, 0xf7, 0x35, 0xca, 0x9a, 0x01, 0x8b, 0x22, 0x16, 0x37, 0x03, 0x1c,
	0x86, 0x7d, 0x1c, 0xdc, 0x14, 0xc8, 0xc1, 0x2d, 0x32, 0xa0, 0x5c, 0xa4, 0xb4, 0x3f, 0x12, 0x94,
	0xc5, 0xf7, 0xd1, 0x90, 0x05, 0x38, 0xa4, 0xbf, 0xe2, 0x29, 0x74, 0xef, 0x16, 0x8d, 0x88, 0x48,
	0x69, 0xc0, 0xef, 0x03, 0x49, 0xca, 0x7e, 0x22, 0x81, 0xb8, 0x0f, 0xf4, 0x69, 0x18, 0xd2, 0x78,
	0x58, 0x00, 0xc7, 0xf2, 0x4f, 0xf0, 0x7c, 0x48, 0xe2, 0xe7, 0xfc, 0x1d, 0x1e, 0x0e, 0x49, 0xda,
	0x64, 0x49, 0xb6, 0x17, 0x6f, 0xe2, 0x38, 0x66, 0x42, 0xee, 0x5b, 0xe8, 0xd7, 0xae, 0x01, 0x70,
	0x19, 0x8d, 0x05, 0xf7, 0xc7, 0x09, 0x81, 0x5f, 0x82, 0x8d, 0x3e, 0x0e, 0x71, 0x1c, 0x90, 0xec,
	0x5f, 0x4d, 0xd1, 0x95, 0xfa, 0xd6, 0xc9, 0x4e, 0xa3, 0x28, 0x4e, 0xe3, 0xf4, 0x16, 0x43, 0xd3,
	0x44, 0x58, 0x03, 0x0f, 0x83, 0x51, 0x9a, 0x92, 0x38, 0x18, 0xb7, 0xd8, 0x80, 0x68, 0x4b, 0xba,
	0x52, 0xaf, 0xa2, 0x99, 0x58, 0xed, 0x37, 0x00, 0xd6, 0xdc, 0xbc, 0xe8, 0x70, 0x0b, 0x2c, 0xd1,
	0x81, 0x94, 0xaf, 0xa2, 0x25, 0x3a, 0x80, 0x10, 0x2c, 0xc7, 0x38, 0x2a, 0xf3, 0xe4, 0x1a, 0x7e,
	0x0d, 0x36, 0x8b, 0x42, 0x91, 0x81, 0x9d, 0x81, 0x15, 0x5d, 0xa9, 0x6f, 0x9c, 0x6c, 0x37, 0x28,
	0x6b, 0x74, 0x4a, 0xc0, 0x13, 0x29, 0x8d, 0x87, 0x68, 0x96, 0x09, 0x9f, 0x81, 0x55, 0x2e, 0xb0,
	0x18, 0x71, 0x6d, 0x59, 0xaf, 0xd4, 0xb7, 0x4e, 0x1e, 0x65, 0x39, 0x6e, 0x5e, 0x3e, 0x4f, 0x02,
	0xa8, 0x20, 0xc0, 0x43, 0xb0, 0xf2, 0x76, 0xc4, 0x04, 0xd6, 0x56, 0xa4, 0x7a, 0x35, 0x63, 0xbe,
	0xce, 0x02, 0x28, 0x8f, 0xc3, 0x06, 0x80, 0x09, 0xe6, 0xb2, 0x3c, 0xd6, 0x80, 0xc4, 0x82, 0x5e,
	0x51, 0x92, 0x6a, 0xab, 0xf2, 0xa0, 0x73, 0x10, 0xd8, 0x01, 0x3b, 0xd3, 0xdd, 0xf7, 0x88, 0x10,
	0x34, 0x1e, 0x72, 0x6d, 0x4d, 0xea, 0x6b, 0x99, 0x7e, 0x7b, 0x0e, 0x8e, 0xe6, 0x66, 0xc1, 0x6f,
	0xc0, 0xff, 0xf1, 0x48, 0xb0, 0x36, 0x09, 0x89, 0x20, 0x6d, 0x3c, 0xe6, 0xc6, 0x95, 0x20, 0xa9,
	0xf9, 0x4b, 0x42, 0xd3, 0xb1, 0xb6, 0xae, 0x2b, 0xf5, 0x4d, 0xb4, 0x98, 0x00, 0x4f, 0xc1, 0x41,
	0x06, 0x22, 0x12, 0xd1, 0x78, 0x40, 0xd2, 0x0c, 0x3e, 0x25, 0x57, 0x2c, 0x25, 0x85, 0x40, 0x55,
	0x0a, 0xfc, 0x2b, 0x07, 0x3e, 0x05, 0x9b, 0x44, 0xae, 0xba, 0x84, 0x73, 0x3c, 0x24, 0x1a, 0x90,
	0x57, 0x9f, 0x0d, 0xc2, 0x0b, 0xb0, 0x3b, 0x69, 0x81, 0x39, 0x43, 0xdf, 0x58, 0xdc, 0xb5, 0x05,
	0x29, 0xf0, 0x1c, 0xe8, 0x73, 0xee, 0x64, 0x33, 0x61, 0xc5, 0x5c, 0x60, 0xf9, 0xd6, 0xb5, 0x87,
	0xf2, 0xe8, 0x1f, 0xe4, 0xc1, 0x6f, 0x41, 0xb5, 0x1c, 0x53, 0xae, 0x6d, 0xe9, 0x95, 0xfa, 0xc6,
	0xc9, 0xe1, 0xe4, 0x3d, 0x17, 0xcf, 0xb1, 0xd1, 0x2a, 0x19, 0x66, 0x2c, 0xd2, 0x31, 0xba, 0xcd,
	0x80, 0x2f, 0x00, 0x48, 0x26, 0xe3, 0xa1, 0xa9, 0xc5, 0x5d, 0x26, 0xf9, 0x13, 0x08, 0x4d, 0xd1,
	0xa0, 0x09, 0xb6, 0x39, 0x09, 0x58, 0x3c, 0xc0, 0xe9, 0xf8, 0x96, 0xa2, 0x3d, 0x5a, 0x9c, 0x3d,
	0x8f, 0x0f, 0x3f, 0x01, 0x6b, 0x85, 0x17, 0x68, 0xdb, 0x32, 0x75, 0x23, 0x2b, 0x62, 0x37, 0x0f,
	0xa1, 0x12, 0x83, 0x9f, 0x82, 0xad, 0x88, 0x08, 0xfc, 0x8a, 0x92, 0x70, 0xc0, 0x3b, 0x94, 0x0b,
	0x6d, 0x47, 0xaf, 0xd4, 0xab, 0xe8, 0x4e, 0x14, 0x7e, 0x01, 0xd6, 0x82, 0x94, 0x60, 0x41, 0x06,
	0xda, 0xff, 0xa4, 0xdc, 0x7e, 0x23, 0xf7, 0xba, 0x46, 0xe9, 0x75, 0x0d, 0xbf, 0xf4, 0x3a, 0x54,
	0x52, 0xb3, 0xac, 0x51, 0x32, 0x90, 0x59, 0xbb, 0x1f, 0xce, 0x2a, 0xa8, 0xd0, 0x05, 0x3b, 0x49,
	0xca, 0xae, 0x68, 0x48, 0xac, 0x08, 0x0f, 0xc9, 0x64, 0x08, 0xf6, 0xa4, 0xa1, 0x1c, 0x4c, 0x37,
	0xe0, 0x2e, 0x09, 0xcd, 0xcd, 0xdc, 0x3f, 0x07, 0x5b, 0xb3, 0x5d, 0x82, 0x2a, 0xa8, 0xdc, 0x90,
	0xb1, 0x34, 0x91, 0x4d, 0x94, 0x2d, 0x61, 0x0d, 0xac, 0xfc, 0x8c, 0xc3, 0x51, 0x6e, 0x23, 0x1b,
	0x27, 0x0f, 0xb3, 0x72, 0x95, 0x49, 0x28, 0x87, 0x5e, 0x2e, 0x7d, 0xa5, 0xbc, 0xe4, 0x7f, 0x1a,
	0x09, 0x88, 0x8f, 0x4a, 0x37, 0x3a, 0xf1, 0xba, 0xf9, 0x71, 0xae, 0x69, 0xa2, 0x17, 0x31, 0xae,
	0xe3, 0x30, 0x64, 0xef, 0x74, 0xac, 0x07, 0x2c, 0x4a, 0x70, 0x3c, 0xd6, 0x05, 0xd3, 0x03, 0x16,
	0x8b, 0x94, 0x85, 0xba, 0xb8, 0x26, 0x7a, 0x7f, 0xc4, 0x69, 0x4c, 0x38, 0xd7, 0x43, 0x36, 0xa4,
	0x81, 0xce, 0xae, 0xb2, 0x28, 0x4d, 0xf5, 0xf2, 0x67, 0xe5, 0xbd, 0x22, 0x5d, 0xec, 0xbd, 0x52,
	0x18, 0xcd, 0xf9, 0xf2, 0xfa, 0xa6, 0xba, 0x75, 0xbe, 0xbc, 0x0e, 0xd5, 0xed, 0xda, 0x5f, 0x0a,
	0x80, 0xc5, 0x46, 0x2d, 0x96, 0x8c, 0x11, 0x79, 0x3b, 0x22, 0x5c, 0xfc, 0x27, 0x57, 0xbc, 0xb5,
	0xb6, 0xca, 0x87, 0xac, 0x6d, 0xbe, 0x73, 0x2d, 0x2f, 0x72, 0xae, 0x5a, 0x1b, 0xec, 0x64, 0x0f,
	0xa5, 0xac, 0x00, 0x22, 0x3c, 0x61, 0x31, 0x27, 0xf0, 0x18, 0xac, 0x17, 0xb7, 0xe2, 0x9a, 0x22,
	0x27, 0x48, 0xbd, 0x3b, 0x41, 0x68, 0xc2, 0xa8, 0x3d, 0x03, 0xdb, 0x67, 0x44, 0x74, 0x89, 0xc0,
	0x17, 0x64, 0x7c, 0x2b, 0x02, 0xc1, 0xf2, 0x0d, 0x19, 0xe7, 0x02, 0x55, 0x24, 0xd7, 0x47, 0x7f,
	0x28, 0xe0, 0x49, 0x21, 0x60, 0x8c, 0x04, 0xbb, 0x4c, 0xda, 0xec, 0x5d, 0x7c, 0x96, 0xe2, 0x01,
	0xf1, 0x29, 0x49, 0xe5, 0x0c, 0xec, 0x02, 0x78, 0xe9, 0xb6, 0x9d, 0x1f, 0xec, 0x33, 0x64, 0xb4,
	0xcd, 0x9e, 0xeb, 0x58, 0xb6, 0xef, 0xa9, 0x0f, 0xa0, 0x0e, 0x0e, 0xa6, 0xe3, 0x9e, 0xd9, 0x72,
	0xec, 0xb6, 0x81, 0xde, 0x94, 0x0c, 0x05, 0x3e, 0x06, 0x7b, 0xd3, 0x0c, 0xdf, 0x32, 0x51, 0x09,
	0x2e, 0xdd, 0x95, 0xfd, 0xde, 0xf2, 0x2c, 0xdf, 0x53, 0x2b, 0x47, 0x11, 0xd8, 0x98, 0xfa, 0x8d,
	0x83, 0x7b, 0x60, 0xfb, 0xd4, 0xe8, 0x18, 0x76, 0xcb, 0xec, 0xf9, 0x6f, 0x5c, 0xb3, 0xe7, 0xf9,
	0xc8, 0xb2, 0xcf, 0xd4, 0x07, 0x70, 0x07, 0xa8, 0x33, 0x80, 0x65, 0xfb, 0xaa, 0x72, 0x8f, 0xde,
	0x76, 0x2e, 0x4f, 0x3b, 0x66, 0xbe, 0xdd, 0x0c, 0xd0, 0x75, 0x6c, 0xf3, 0x8d, 0x5a, 0x39, 0x22,
	0x60, 0x7b, 0xce, 0x04, 0x64, 0x74, 0x17, 0x39, 0xaf, 0xac, 0x8e, 0xd9, 0xb3, 0xba, 0xc6, 0x99,
	0xd9, 0xb3, 0x1d, 0xdb, 0x54, 0x1f, 0xc0, 0x7d, 0xb0, 0x3b, 0x1b, 0x77, 0x5c, 0xdf, 0x72, 0x6c,
	0xa3, 0xa3, 0x2a, 0xf7, 0x31, 0x64, 0xbe, 0xbe, 0xb4, 0x90, 0xd9, 0x56, 0x97, 0x8e, 0x7e, 0xaf,
	0xc8, 0xe7, 0x26, 0x58, 0xc0, 0x42, 0x17, 0x73, 0x9e, 0x9b, 0x02, 0xac, 0x81, 0x27, 0x2e, 0x72,
	0x7c, 0xa7, 0xe5, 0x74, 0x7a, 0x5d, 0xb3, 0x7b, 0x6a, 0x22, 0xef, 0x3b, 0xcb, 0xed, 0xb5, 0x9d,
	0x9e, 0xed, 0xf8, 0xbd, 0x4b, 0x2f, 0xdb, 0xf2, 0x09, 0xd8, 0x9f, 0xc7, 0xb1, 0x2f, 0xb3, 0xa5,
	0xaa, 0xc0, 0x43, 0xf0, 0x78, 0xae, 0x86, 0xe9, 0x1b, 0x56, 0x27, 0xab, 0xf4, 0xc7, 0xe0, 0xa3,
	0x79, 0x04, 0xd9, 0x0e, 0xdb, 0xe8, 0x9a, 0x6a, 0x05, 0xd6, 0xc1, 0xd3, 0x85, 0x14, 0xcb, 0x7e,
	0xe5, 0xa0, 0xae, 0x91, 0x5d, 0x53, 0x5d, 0x5e, 0x74, 0x9a, 0xa2, 0xad, 0x2b, 0x8b, 0x94, 0xee,
	0xbd, 0x8e, 0x55, 0x78, 0x0c, 0xea, 0xf3, 0x98, 0x67, 0xa6, 0x6d, 0x22, 0xab, 0xd5, 0x73, 0x91,
	0x73, 0x86, 0x8c, 0xae, 0xdc, 0x5e, 0x5d, 0x5b, 0x74, 0x89, 0x8e, 0xd3, 0x92, 0x27, 0xf3, 0xd4,
	0x75, 0xf8, 0x39, 0xf8, 0x6c, 0x2e, 0xc5, 0xf0, 0x4d, 0xcf, 0xef, 0xf9, 0xc8, 0xb0, 0x3d, 0xa3,
	0x95, 0x93, 0xab, 0xa7, 0x5e, 0xf6, 0x21, 0xda, 0xc8, 0x46, 0xf0, 0x86, 0x8a, 0x86, 0x7b, 0x21,
	0x6d, 0x5d, 0x0e, 0xd4, 0x8f, 0xc7, 0x5c, 0x60, 0x7e, 0x3d, 0xc1, 0x02, 0x16, 0x35, 0xb3, 0x6f,
	0x56, 0x36, 0x20, 0x61, 0x93, 0x0f, 0x6e, 0x9a, 0x43, 0xd6, 0x9c, 0x7c, 0xc3, 0xf2, 0xbf, 0x97,
	0xaa, 0xee, 0x45, 0x91, 0xd9, 0x5f, 0x95, 0x86, 0xfc, 0xe2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xee, 0xf5, 0xeb, 0xd1, 0xe4, 0x0a, 0x00, 0x00,
}
