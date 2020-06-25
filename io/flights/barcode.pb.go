// Code generated by protoc-gen-go. DO NOT EDIT.
// source: io/flights/barcode.proto

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

// Passenger Status as detailed in IATA PSC Resolution 792 Attachment 'C'. Note: all values other than 0 indicate that the passenger has checked in.
type PassengerStatus int32

const (
	// Ticket issuance/passenger not checked in
	PassengerStatus_ISSUED_NOT_CHECKED_IN PassengerStatus = 0
	// Ticket issuance/passenger checked in
	PassengerStatus_ISSUED_CHECKED_IN PassengerStatus = 1
	// Baggage checked - passenger not checked in
	PassengerStatus_BAGGAGE_CHECKED_PASSENGER_NOT_CHECKED_IN PassengerStatus = 2
	// Baggage checked - passenger checked in
	PassengerStatus_BAGGAGE_CHECKED_PASSENGER_CHECKED_IN PassengerStatus = 3
	// Passenger passed security check
	PassengerStatus_PASSENGER_PASSED_SECURITY_CHECK PassengerStatus = 4
	// Passenger passed gate and exit (coupon used)
	PassengerStatus_PASSENGER_PASSED_GATE PassengerStatus = 5
	// Transit
	PassengerStatus_TRANSIT PassengerStatus = 6
	// Standby. Seat number not printed on boarding pass at time of check in
	PassengerStatus_STANDBY PassengerStatus = 7
	// Boarding data re-validation done
	PassengerStatus_BOARDING_DATA_REVALIDATION_DONE PassengerStatus = 8
	// Original boarding line used at time of ticket issuance
	PassengerStatus_ORIGINAL_BOARDING_LINE_USED_AT_TICKET_ISSUANCE PassengerStatus = 9
	// Up or down grading required
	PassengerStatus_UP_OR_DOWN_GRADING_REQUIRED PassengerStatus = 10
)

var PassengerStatus_name = map[int32]string{
	0:  "ISSUED_NOT_CHECKED_IN",
	1:  "ISSUED_CHECKED_IN",
	2:  "BAGGAGE_CHECKED_PASSENGER_NOT_CHECKED_IN",
	3:  "BAGGAGE_CHECKED_PASSENGER_CHECKED_IN",
	4:  "PASSENGER_PASSED_SECURITY_CHECK",
	5:  "PASSENGER_PASSED_GATE",
	6:  "TRANSIT",
	7:  "STANDBY",
	8:  "BOARDING_DATA_REVALIDATION_DONE",
	9:  "ORIGINAL_BOARDING_LINE_USED_AT_TICKET_ISSUANCE",
	10: "UP_OR_DOWN_GRADING_REQUIRED",
}

var PassengerStatus_value = map[string]int32{
	"ISSUED_NOT_CHECKED_IN":                          0,
	"ISSUED_CHECKED_IN":                              1,
	"BAGGAGE_CHECKED_PASSENGER_NOT_CHECKED_IN":       2,
	"BAGGAGE_CHECKED_PASSENGER_CHECKED_IN":           3,
	"PASSENGER_PASSED_SECURITY_CHECK":                4,
	"PASSENGER_PASSED_GATE":                          5,
	"TRANSIT":                                        6,
	"STANDBY":                                        7,
	"BOARDING_DATA_REVALIDATION_DONE":                8,
	"ORIGINAL_BOARDING_LINE_USED_AT_TICKET_ISSUANCE": 9,
	"UP_OR_DOWN_GRADING_REQUIRED":                    10,
}

func (x PassengerStatus) String() string {
	return proto.EnumName(PassengerStatus_name, int32(x))
}

func (PassengerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{0}
}

// Passenger Description as detailed in IATA PSC Resolution 792 Attachment 'C'.
type PassengerDescription int32

const (
	// Adult
	PassengerDescription_ADULT PassengerDescription = 0
	// Adult male
	PassengerDescription_MALE PassengerDescription = 1
	// Adult female
	PassengerDescription_FEMALE PassengerDescription = 2
	// Child
	PassengerDescription_CHILD PassengerDescription = 3
	// Infant
	PassengerDescription_INFANT PassengerDescription = 4
	// No passenger (cabin baggage)
	PassengerDescription_NO_PASSENGER PassengerDescription = 5
	// Adult travelling with infant
	PassengerDescription_ADULT_WITH_INFANT PassengerDescription = 6
	// Unaccompanied minor
	PassengerDescription_UNACCOMPANIED_MINOR PassengerDescription = 7
)

var PassengerDescription_name = map[int32]string{
	0: "ADULT",
	1: "MALE",
	2: "FEMALE",
	3: "CHILD",
	4: "INFANT",
	5: "NO_PASSENGER",
	6: "ADULT_WITH_INFANT",
	7: "UNACCOMPANIED_MINOR",
}

var PassengerDescription_value = map[string]int32{
	"ADULT":               0,
	"MALE":                1,
	"FEMALE":              2,
	"CHILD":               3,
	"INFANT":              4,
	"NO_PASSENGER":        5,
	"ADULT_WITH_INFANT":   6,
	"UNACCOMPANIED_MINOR": 7,
}

func (x PassengerDescription) String() string {
	return proto.EnumName(PassengerDescription_name, int32(x))
}

func (PassengerDescription) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{1}
}

// Source of transaction.
type Source int32

const (
	// Web
	Source_W Source = 0
	// Airport Kiosk
	Source_K Source = 1
	// Transfer kiosk
	Source_X Source = 2
	// Remote / off-site kiosk
	Source_R Source = 3
	// Mobile device
	Source_M Source = 4
	// Airport agent
	Source_O Source = 5
	// Town agent
	Source_T Source = 6
	// Third party vendor
	Source_V Source = 7
)

var Source_name = map[int32]string{
	0: "W",
	1: "K",
	2: "X",
	3: "R",
	4: "M",
	5: "O",
	6: "T",
	7: "V",
}

var Source_value = map[string]int32{
	"W": 0,
	"K": 1,
	"X": 2,
	"R": 3,
	"M": 4,
	"O": 5,
	"T": 6,
	"V": 7,
}

func (x Source) String() string {
	return proto.EnumName(Source_name, int32(x))
}

func (Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{2}
}

// Document type that the barcode represents.
type DocType int32

const (
	// Boarding pass
	DocType_B DocType = 0
	// Itinerary receipt
	DocType_I DocType = 1
)

var DocType_name = map[int32]string{
	0: "B",
	1: "I",
}

var DocType_value = map[string]int32{
	"B": 0,
	"I": 1,
}

func (x DocType) String() string {
	return proto.EnumName(DocType_name, int32(x))
}

func (DocType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{3}
}

// International Documentation Verification.
type InternationalDocVerification int32

const (
	// No travel document verification required
	InternationalDocVerification_NOT_REQUIRED InternationalDocVerification = 0
	// Travel document verification required before boarding
	InternationalDocVerification_REQUIRED InternationalDocVerification = 1
	// Travel document verification successfully completed
	InternationalDocVerification_COMPLETED InternationalDocVerification = 2
)

var InternationalDocVerification_name = map[int32]string{
	0: "NOT_REQUIRED",
	1: "REQUIRED",
	2: "COMPLETED",
}

var InternationalDocVerification_value = map[string]int32{
	"NOT_REQUIRED": 0,
	"REQUIRED":     1,
	"COMPLETED":    2,
}

func (x InternationalDocVerification) String() string {
	return proto.EnumName(InternationalDocVerification_name, int32(x))
}

func (InternationalDocVerification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{4}
}

// Industry Discount / Airline Discount codes. See IATA Recommended Practice 1788.
type IDADIndicator int32

const (
	// No Industry discount
	IDADIndicator_NONE IDADIndicator = 0
	// Industry discount, other carrier employee/all other eligible persons, firm reservation
	IDADIndicator_IDN1 IDADIndicator = 1
	// Industry discount, other carrier employee/all other eligible persons, space available
	IDADIndicator_IDN2 IDADIndicator = 2
	// Industry discount, other carrier employee travelling on duty, firm reservation
	IDADIndicator_IDB1 IDADIndicator = 3
	// Industry discount, other carrier employee travelling on duty, pace available
	IDADIndicator_IDB2 IDADIndicator = 4
	// Ticket issued pursuant to Resolution 880
	IDADIndicator_AD IDADIndicator = 5
	// Ticket issued pursuant to Resolution 200g
	IDADIndicator_DG IDADIndicator = 6
	// Discount not covered by industry regulations (for online use only)
	IDADIndicator_DM IDADIndicator = 7
	// Ticket issued pursuant to Resolution 886
	IDADIndicator_GE IDADIndicator = 8
	// Ticket issued pursuant to Resolution 788
	IDADIndicator_IG IDADIndicator = 9
	// Ticket issued pursuant to Resolution 888
	IDADIndicator_RG IDADIndicator = 10
	// Ticket issued pursuant to Resolution 884
	IDADIndicator_UD IDADIndicator = 11
	// Industry discount ticket - no classification designator
	IDADIndicator_ID IDADIndicator = 12
	// Industry discount, member's own employee travelling on duty, firm reservation
	IDADIndicator_IDFS1 IDADIndicator = 13
	// Industry discount, member's own employee travelling on duty, space available
	IDADIndicator_IDFS2 IDADIndicator = 14
	// Industry discount, member's own employee/dependant, firm reservation
	IDADIndicator_IDR1 IDADIndicator = 15
	// Industry discount, member's own employee/dependant, space available
	IDADIndicator_IDR2 IDADIndicator = 16
)

var IDADIndicator_name = map[int32]string{
	0:  "NONE",
	1:  "IDN1",
	2:  "IDN2",
	3:  "IDB1",
	4:  "IDB2",
	5:  "AD",
	6:  "DG",
	7:  "DM",
	8:  "GE",
	9:  "IG",
	10: "RG",
	11: "UD",
	12: "ID",
	13: "IDFS1",
	14: "IDFS2",
	15: "IDR1",
	16: "IDR2",
}

var IDADIndicator_value = map[string]int32{
	"NONE":  0,
	"IDN1":  1,
	"IDN2":  2,
	"IDB1":  3,
	"IDB2":  4,
	"AD":    5,
	"DG":    6,
	"DM":    7,
	"GE":    8,
	"IG":    9,
	"RG":    10,
	"UD":    11,
	"ID":    12,
	"IDFS1": 13,
	"IDFS2": 14,
	"IDR1":  15,
	"IDR2":  16,
}

func (x IDADIndicator) String() string {
	return proto.EnumName(IDADIndicator_name, int32(x))
}

func (IDADIndicator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{5}
}

// Passenger security status. Note: in practical terms, this is always likely to be NOT_SELECTEE or TSA_PRECHECK since current DHS Pre-Departure Regulations only allow on-site printing of boarding passes. It is mandatory when US travel is involved.
type SelecteeIndicator int32

const (
	// Not selected for special security screening
	SelecteeIndicator_NOT_SELECTEE SelecteeIndicator = 0
	// Selected for special security screening
	SelecteeIndicator_SELECTEE SelecteeIndicator = 1
	// Passenger eligible for TSA PreCheck®
	SelecteeIndicator_TSA_PRECHECK SelecteeIndicator = 3
)

var SelecteeIndicator_name = map[int32]string{
	0: "NOT_SELECTEE",
	1: "SELECTEE",
	3: "TSA_PRECHECK",
}

var SelecteeIndicator_value = map[string]int32{
	"NOT_SELECTEE": 0,
	"SELECTEE":     1,
	"TSA_PRECHECK": 3,
}

func (x SelecteeIndicator) String() string {
	return proto.EnumName(SelecteeIndicator_name, int32(x))
}

func (SelecteeIndicator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{6}
}

// Passenger eligibility for Fast-track privileges.
type FastTrack int32

const (
	// Not Eligible
	FastTrack_N FastTrack = 0
	// Eligible
	FastTrack_Y FastTrack = 1
)

var FastTrack_name = map[int32]string{
	0: "N",
	1: "Y",
}

var FastTrack_value = map[string]int32{
	"N": 0,
	"Y": 1,
}

func (x FastTrack) String() string {
	return proto.EnumName(FastTrack_name, int32(x))
}

func (FastTrack) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{7}
}

// Conditional items for the flight segment as described in IATA PSC Resolution 792.  Note: items supplied or derived through other objects (E.g. Frequent Flyer information) are not supplied here.
type ConditionalItems struct {
	// Passenger Status as detailed in IATA PSC Resolution 792 Attachment 'C'. Note: all values other than 0 indicate that the passenger has checked in.
	// @tag: validateGeneric:"omitempty"
	PassengerDescription PassengerDescription `protobuf:"varint,1,opt,name=passengerDescription,proto3,enum=flights.PassengerDescription" json:"passengerDescription,omitempty" validateGeneric:"omitempty"`
	// Source of checkin as detailed in IATA PSC Resolution 792 Attachment 'C'.
	// @tag: validateGeneric:"omitempty"
	CheckInSource Source `protobuf:"varint,2,opt,name=checkInSource,proto3,enum=flights.Source" json:"checkInSource,omitempty" validateGeneric:"omitempty"`
	// Source of boarding pass issuance as detailed in IATA PSC Resolution 792 Attachment 'C'.
	// @tag: validateGeneric:"omitempty"
	BoardingPassIssuanceSource Source `protobuf:"varint,3,opt,name=boardingPassIssuanceSource,proto3,enum=flights.Source" json:"boardingPassIssuanceSource,omitempty" validateGeneric:"omitempty"`
	// Date of boarding pass issuance, local to the boarding point.
	// @tag: validateGeneric:"omitempty"
	BoardingPassIssueDate *io.Date `protobuf:"bytes,4,opt,name=boardingPassIssueDate,proto3" json:"boardingPassIssueDate,omitempty" validateGeneric:"omitempty"`
	// Document type that the barcode represents.
	// @tag: validateGeneric:"omitempty"
	DocumentType DocType `protobuf:"varint,5,opt,name=documentType,proto3,enum=flights.DocType" json:"documentType,omitempty" validateGeneric:"omitempty"`
	// IATA or ICAO designator of boarding pass issuer.
	// @tag: validateGeneric:"omitempty"
	BoardingPassIssuer string `protobuf:"bytes,6,opt,name=boardingPassIssuer,proto3" json:"boardingPassIssuer,omitempty" validateGeneric:"omitempty"`
	// This field allows carriers to populate baggage tag numbers and the number of consecutive bags. It contains 13 characters corresponding to the 10 digit bag tag number, as per IATA BCM specifications, Resolution 740 and 3 digits identifying the number of consecutive tags.\n1: leading digit – 0 for interline tag, 1 for fall-back tag, 2 for interline rush tag.\n2-4: carrier numeric code.\n5-10: carrier initial tag number (leading zeros).\n11-13: number of consecutive tags (allows for up to 999 tags).\nUp to 2 additional, non-consecutive tags can be added.
	// @tag: validateGeneric:"omitempty,dive,numeric"
	BaggageTagNumber []string `protobuf:"bytes,7,rep,name=baggageTagNumber,proto3" json:"baggageTagNumber,omitempty" validateGeneric:"omitempty,dive,numeric"`
	// Passenger security status. Note: in practical terms, this is always likely to be NOT_SELECTEE or TSA_PRECHECK since current DHS Pre-Departure Regulations only allow on-site printing of boarding passes. It is mandatory when US travel is involved.
	// @tag: validateGeneric:"omitempty"
	SelecteeIndicator SelecteeIndicator `protobuf:"varint,8,opt,name=selecteeIndicator,proto3,enum=flights.SelecteeIndicator" json:"selecteeIndicator,omitempty" validateGeneric:"omitempty"`
	// International Documentation Verification status as detailed in IATA PSC Resolution 792 Attachment 'C'.
	// @tag: validateGeneric:"omitempty"
	InternationalDocVerification InternationalDocVerification `protobuf:"varint,9,opt,name=internationalDocVerification,proto3,enum=flights.InternationalDocVerification" json:"internationalDocVerification,omitempty" validateGeneric:"omitempty"`
	// Industry Discount / Airline Discount codes. See IATA Recommended Practice 1788.
	// @tag: validateGeneric:"omitempty"
	IdadIndicator IDADIndicator `protobuf:"varint,10,opt,name=idadIndicator,proto3,enum=flights.IDADIndicator" json:"idadIndicator,omitempty" validateGeneric:"omitempty"`
	// Passenger eligibility for Fast-track privileges.
	// @tag: validateGeneric:"omitempty"
	FastTrack            FastTrack `protobuf:"varint,11,opt,name=fastTrack,proto3,enum=flights.FastTrack" json:"fastTrack,omitempty" validateGeneric:"omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConditionalItems) Reset()         { *m = ConditionalItems{} }
func (m *ConditionalItems) String() string { return proto.CompactTextString(m) }
func (*ConditionalItems) ProtoMessage()    {}
func (*ConditionalItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a35f1da235ff26, []int{0}
}

func (m *ConditionalItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConditionalItems.Unmarshal(m, b)
}
func (m *ConditionalItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConditionalItems.Marshal(b, m, deterministic)
}
func (m *ConditionalItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConditionalItems.Merge(m, src)
}
func (m *ConditionalItems) XXX_Size() int {
	return xxx_messageInfo_ConditionalItems.Size(m)
}
func (m *ConditionalItems) XXX_DiscardUnknown() {
	xxx_messageInfo_ConditionalItems.DiscardUnknown(m)
}

var xxx_messageInfo_ConditionalItems proto.InternalMessageInfo

func (m *ConditionalItems) GetPassengerDescription() PassengerDescription {
	if m != nil {
		return m.PassengerDescription
	}
	return PassengerDescription_ADULT
}

func (m *ConditionalItems) GetCheckInSource() Source {
	if m != nil {
		return m.CheckInSource
	}
	return Source_W
}

func (m *ConditionalItems) GetBoardingPassIssuanceSource() Source {
	if m != nil {
		return m.BoardingPassIssuanceSource
	}
	return Source_W
}

func (m *ConditionalItems) GetBoardingPassIssueDate() *io.Date {
	if m != nil {
		return m.BoardingPassIssueDate
	}
	return nil
}

func (m *ConditionalItems) GetDocumentType() DocType {
	if m != nil {
		return m.DocumentType
	}
	return DocType_B
}

func (m *ConditionalItems) GetBoardingPassIssuer() string {
	if m != nil {
		return m.BoardingPassIssuer
	}
	return ""
}

func (m *ConditionalItems) GetBaggageTagNumber() []string {
	if m != nil {
		return m.BaggageTagNumber
	}
	return nil
}

func (m *ConditionalItems) GetSelecteeIndicator() SelecteeIndicator {
	if m != nil {
		return m.SelecteeIndicator
	}
	return SelecteeIndicator_NOT_SELECTEE
}

func (m *ConditionalItems) GetInternationalDocVerification() InternationalDocVerification {
	if m != nil {
		return m.InternationalDocVerification
	}
	return InternationalDocVerification_NOT_REQUIRED
}

func (m *ConditionalItems) GetIdadIndicator() IDADIndicator {
	if m != nil {
		return m.IdadIndicator
	}
	return IDADIndicator_NONE
}

func (m *ConditionalItems) GetFastTrack() FastTrack {
	if m != nil {
		return m.FastTrack
	}
	return FastTrack_N
}

func init() {
	proto.RegisterEnum("flights.PassengerStatus", PassengerStatus_name, PassengerStatus_value)
	proto.RegisterEnum("flights.PassengerDescription", PassengerDescription_name, PassengerDescription_value)
	proto.RegisterEnum("flights.Source", Source_name, Source_value)
	proto.RegisterEnum("flights.DocType", DocType_name, DocType_value)
	proto.RegisterEnum("flights.InternationalDocVerification", InternationalDocVerification_name, InternationalDocVerification_value)
	proto.RegisterEnum("flights.IDADIndicator", IDADIndicator_name, IDADIndicator_value)
	proto.RegisterEnum("flights.SelecteeIndicator", SelecteeIndicator_name, SelecteeIndicator_value)
	proto.RegisterEnum("flights.FastTrack", FastTrack_name, FastTrack_value)
	proto.RegisterType((*ConditionalItems)(nil), "flights.ConditionalItems")
}

func init() {
	proto.RegisterFile("io/flights/barcode.proto", fileDescriptor_00a35f1da235ff26)
}

var fileDescriptor_00a35f1da235ff26 = []byte{
	// 1303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x6e, 0xdb, 0xc6,
	0x12, 0x36, 0x65, 0xfd, 0x58, 0x6b, 0x3b, 0xd9, 0xec, 0x49, 0x4e, 0x14, 0x9f, 0x9c, 0x66, 0xd1,
	0x1f, 0x40, 0x11, 0x6c, 0xca, 0x92, 0x13, 0x14, 0x09, 0x8a, 0xa2, 0x94, 0x48, 0xcb, 0x84, 0x65,
	0xca, 0xa1, 0xa8, 0xa4, 0xe9, 0x0d, 0xb1, 0x22, 0x57, 0xf2, 0xd6, 0x34, 0x57, 0xe0, 0x52, 0x0d,
	0xd2, 0xab, 0x3e, 0x42, 0xd1, 0x5e, 0xf5, 0x15, 0x52, 0xf4, 0xa2, 0x2f, 0xd4, 0x8b, 0xbe, 0x40,
	0x5f, 0xa1, 0xd8, 0x25, 0x2d, 0xf9, 0x2f, 0xbe, 0x9a, 0xd9, 0x99, 0x6f, 0xbe, 0xfd, 0x38, 0xbb,
	0x43, 0x12, 0xd4, 0x18, 0x6f, 0x4e, 0x22, 0x36, 0x3d, 0x49, 0x45, 0x73, 0x4c, 0x92, 0x80, 0x87,
	0x54, 0x9f, 0x25, 0x3c, 0xe5, 0xa8, 0x92, 0x87, 0xb7, 0x3e, 0x61, 0xbc, 0x19, 0xf0, 0xb3, 0x33,
	0x1e, 0xe7, 0xc6, 0xe7, 0xe3, 0xef, 0x69, 0x90, 0x8a, 0x0c, 0xb8, 0xb5, 0xad, 0x4c, 0xb0, 0x33,
	0xa5, 0xf1, 0x8e, 0x78, 0x47, 0xa6, 0x53, 0x9a, 0x34, 0xf9, 0x2c, 0x65, 0x3c, 0x16, 0x4d, 0x12,
	0xc7, 0x3c, 0x25, 0xca, 0xcf, 0xd0, 0x9f, 0xfe, 0x06, 0x00, 0xec, 0xf2, 0x38, 0x64, 0x32, 0x48,
	0x22, 0x3b, 0xa5, 0x67, 0x02, 0xbd, 0x02, 0xf7, 0x67, 0x44, 0x08, 0x1a, 0x4f, 0x69, 0x62, 0x52,
	0x11, 0x24, 0x4c, 0xd5, 0xd7, 0x34, 0xac, 0xd5, 0xef, 0xb4, 0xff, 0xaf, 0xe7, 0x52, 0xf4, 0xe3,
	0x1b, 0x40, 0xee, 0x8d, 0xa5, 0xe8, 0x39, 0xd8, 0x0c, 0x4e, 0x68, 0x70, 0x6a, 0xc7, 0x43, 0x3e,
	0x4f, 0x02, 0x5a, 0x2b, 0x28, 0xae, 0xbb, 0x0b, 0xae, 0x2c, 0xec, 0x5e, 0x46, 0xa1, 0x01, 0xd8,
	0x1a, 0x73, 0x92, 0x84, 0x2c, 0x9e, 0xca, 0xcd, 0x6c, 0x21, 0xe6, 0x24, 0x0e, 0x68, 0xce, 0xb1,
	0x7a, 0x33, 0xc7, 0x2d, 0x25, 0xe8, 0x6b, 0xf0, 0xe0, 0x6a, 0x96, 0x9a, 0x24, 0xa5, 0xb5, 0x22,
	0xd6, 0xea, 0xeb, 0xed, 0x35, 0x9d, 0x71, 0x5d, 0xae, 0xdd, 0x9b, 0x61, 0xe8, 0x19, 0xd8, 0x08,
	0x79, 0x30, 0x3f, 0xa3, 0x71, 0xea, 0xbd, 0x9f, 0xd1, 0x5a, 0x49, 0x49, 0x80, 0x0b, 0x09, 0x26,
	0x0f, 0x64, 0xdc, 0xbd, 0x84, 0x42, 0x3a, 0x40, 0xd7, 0xe8, 0x92, 0x5a, 0x19, 0x6b, 0xf5, 0xaa,
	0x7b, 0x43, 0x06, 0xfd, 0x55, 0x04, 0x70, 0x4c, 0xa6, 0x53, 0x32, 0xa5, 0x1e, 0x99, 0x3a, 0xf3,
	0xb3, 0x31, 0x4d, 0x6a, 0x15, 0xbc, 0x5a, 0xaf, 0x76, 0xfe, 0x2c, 0xfe, 0x62, 0xfc, 0x51, 0x6c,
	0x7f, 0x28, 0x7a, 0x27, 0x4c, 0xe0, 0x09, 0xa3, 0x51, 0x88, 0x49, 0x14, 0xf1, 0x77, 0x02, 0x07,
	0x24, 0x49, 0x18, 0x4d, 0x04, 0x4e, 0x39, 0x9e, 0xf1, 0xd9, 0x3c, 0x22, 0x29, 0xc5, 0x39, 0x0d,
	0x4e, 0xc9, 0x14, 0xc7, 0x8a, 0x48, 0x60, 0x12, 0x87, 0x38, 0x3d, 0xa1, 0xf9, 0x1a, 0xf3, 0x09,
	0x0e, 0x78, 0x2c, 0x68, 0x30, 0x4f, 0xd9, 0x0f, 0xaa, 0x42, 0xe8, 0xd8, 0x4e, 0x65, 0x30, 0x25,
	0x2c, 0x16, 0xb8, 0xb5, 0x87, 0x83, 0x13, 0x92, 0x90, 0x20, 0x95, 0xd5, 0x01, 0x4f, 0x12, 0x2a,
	0x66, 0xf2, 0xd2, 0xc4, 0x53, 0xb9, 0x97, 0xa4, 0x6a, 0xed, 0xe2, 0x90, 0x4d, 0x59, 0x2a, 0xab,
	0x2f, 0xec, 0xb5, 0x8d, 0x89, 0xc0, 0x33, 0x9a, 0x60, 0xdb, 0xf0, 0x0c, 0xdc, 0xe9, 0x1e, 0x61,
	0x31, 0xa3, 0x01, 0x9b, 0xb0, 0x20, 0xbb, 0x85, 0xdb, 0xd8, 0xa5, 0x82, 0x47, 0x73, 0xb9, 0xc0,
	0x5f, 0x3e, 0xdb, 0x55, 0xd2, 0xf6, 0x32, 0x2a, 0x81, 0x59, 0x48, 0xe3, 0x94, 0x4d, 0xde, 0xab,
	0x7d, 0x3e, 0xaa, 0x37, 0x95, 0x7a, 0x41, 0xeb, 0x25, 0x8e, 0x28, 0x51, 0x92, 0x32, 0x21, 0x7f,
	0xff, 0xf4, 0x2b, 0xde, 0xc5, 0x13, 0x9e, 0x60, 0x16, 0xa7, 0x34, 0x89, 0x58, 0xac, 0xa0, 0xdb,
	0xb8, 0xa5, 0x82, 0x13, 0x12, 0x45, 0x3b, 0x63, 0x12, 0x9c, 0x66, 0xc1, 0xf6, 0x15, 0x64, 0x32,
	0x17, 0x27, 0x32, 0xa3, 0x83, 0xf6, 0xce, 0xb3, 0x97, 0xe7, 0xad, 0x95, 0x0a, 0x68, 0xc2, 0x02,
	0xac, 0x06, 0x13, 0x3c, 0xdf, 0x69, 0xed, 0x2e, 0x73, 0x2c, 0x66, 0x29, 0x23, 0xd1, 0x85, 0xa7,
	0xc7, 0xf5, 0x73, 0x49, 0x3f, 0xd2, 0x84, 0x8b, 0xa7, 0x3a, 0x68, 0xb5, 0x76, 0x5a, 0x7b, 0x2f,
	0x6f, 0x79, 0x10, 0x5c, 0xcf, 0x0f, 0x53, 0xca, 0x99, 0xcf, 0x64, 0x7b, 0x5f, 0xbc, 0x78, 0xa1,
	0x52, 0x4f, 0x75, 0x30, 0x52, 0x81, 0x36, 0x26, 0xe1, 0xf9, 0xc0, 0x6e, 0xe3, 0x98, 0xc7, 0x3b,
	0xd7, 0x58, 0x02, 0x12, 0xe3, 0x31, 0x95, 0x38, 0x1a, 0xea, 0xee, 0xb5, 0xbb, 0x84, 0x0e, 0xc0,
	0x3d, 0x41, 0x23, 0x1a, 0xa4, 0x94, 0xda, 0x71, 0x28, 0x4f, 0x83, 0x27, 0xb5, 0x35, 0x75, 0x97,
	0xb7, 0x96, 0xe3, 0x74, 0x15, 0xe1, 0x5e, 0x2f, 0x42, 0x0c, 0x3c, 0x56, 0x5d, 0x8b, 0x49, 0x26,
	0xc8, 0xe4, 0xc1, 0x6b, 0x9a, 0x2c, 0x4e, 0xb8, 0x56, 0x55, 0xa4, 0x5f, 0x2c, 0x48, 0xed, 0x5b,
	0xc0, 0xee, 0xad, 0x54, 0xe8, 0x2b, 0xb0, 0xc9, 0x42, 0x12, 0x2e, 0x05, 0x03, 0xc5, 0xfd, 0xdf,
	0x25, 0xb7, 0x69, 0x98, 0x4b, 0xb1, 0x97, 0xc1, 0x68, 0x17, 0x54, 0x27, 0x44, 0xa4, 0x5e, 0x42,
	0x82, 0xd3, 0xda, 0xba, 0xaa, 0x44, 0x8b, 0xca, 0xfd, 0xf3, 0x8c, 0xbb, 0x04, 0x35, 0xfe, 0x29,
	0x80, 0xbb, 0x8b, 0x57, 0xdc, 0x30, 0x25, 0xe9, 0x5c, 0xa0, 0x47, 0xe0, 0x81, 0x3d, 0x1c, 0x8e,
	0x2c, 0xd3, 0x77, 0x06, 0x9e, 0xdf, 0x3d, 0xb0, 0xba, 0x87, 0x96, 0xe9, 0xdb, 0x0e, 0x5c, 0x41,
	0x0f, 0xc0, 0xbd, 0x3c, 0x75, 0x21, 0xac, 0xa1, 0x6d, 0x50, 0xef, 0x18, 0xbd, 0x9e, 0xd1, 0xb3,
	0x16, 0xf1, 0x63, 0x63, 0x38, 0xb4, 0x9c, 0x9e, 0xe5, 0x5e, 0x25, 0x29, 0xa0, 0x3a, 0xf8, 0xfc,
	0xe3, 0xe8, 0x0b, 0xc8, 0x55, 0xf4, 0x19, 0x78, 0xb2, 0xcc, 0x28, 0xcf, 0xf4, 0x87, 0x56, 0x77,
	0xe4, 0xda, 0xde, 0xdb, 0x0c, 0x09, 0x8b, 0x52, 0xee, 0x35, 0x50, 0xcf, 0xf0, 0x2c, 0x58, 0x42,
	0xeb, 0xa0, 0xe2, 0xb9, 0x86, 0x33, 0xb4, 0x3d, 0x58, 0x96, 0x8b, 0xa1, 0x67, 0x38, 0x66, 0xe7,
	0x2d, 0xac, 0x48, 0xe6, 0xce, 0xc0, 0x70, 0x4d, 0xdb, 0xe9, 0xf9, 0xa6, 0xe1, 0x19, 0xbe, 0x6b,
	0xbd, 0x36, 0xfa, 0xb6, 0x69, 0x78, 0xf6, 0xc0, 0xf1, 0xcd, 0x81, 0x63, 0xc1, 0x35, 0xd4, 0x06,
	0xfa, 0xc0, 0xb5, 0x7b, 0xb6, 0x63, 0xf4, 0xfd, 0x05, 0xba, 0x6f, 0x3b, 0x96, 0x3f, 0x92, 0x9b,
	0x18, 0x9e, 0xef, 0xd9, 0xdd, 0x43, 0xcb, 0xf3, 0x65, 0x4f, 0x0c, 0xa7, 0x6b, 0xc1, 0x2a, 0x7a,
	0x02, 0xfe, 0x37, 0x3a, 0xf6, 0x07, 0xae, 0x6f, 0x0e, 0xde, 0x38, 0x7e, 0xcf, 0x35, 0x54, 0x91,
	0x6b, 0xbd, 0x1a, 0xd9, 0xae, 0x65, 0x42, 0xd0, 0xf8, 0x59, 0x03, 0xf7, 0x6f, 0xfa, 0xa8, 0xa0,
	0x2a, 0x28, 0x19, 0xe6, 0xa8, 0xef, 0xc1, 0x15, 0xb4, 0x06, 0x8a, 0x47, 0x46, 0xdf, 0x82, 0x1a,
	0x02, 0xa0, 0xbc, 0x6f, 0x29, 0xbf, 0x20, 0x01, 0xdd, 0x03, 0xbb, 0x6f, 0xc2, 0x55, 0x19, 0xb6,
	0x9d, 0x7d, 0xc3, 0xf1, 0x60, 0x11, 0x41, 0xb0, 0xe1, 0x0c, 0x96, 0x1d, 0x84, 0x25, 0x79, 0x4a,
	0x8a, 0xc9, 0x7f, 0x63, 0x7b, 0x07, 0x7e, 0x0e, 0x2c, 0xa3, 0x87, 0xe0, 0x3f, 0x23, 0xc7, 0xe8,
	0x76, 0x07, 0x47, 0xc7, 0x86, 0x63, 0x5b, 0xa6, 0x7f, 0x64, 0x3b, 0x03, 0x17, 0x56, 0x1a, 0xdf,
	0x80, 0x72, 0xfe, 0xe9, 0x28, 0x01, 0xed, 0x0d, 0x5c, 0x91, 0xe6, 0x10, 0x6a, 0xd2, 0x7c, 0x0b,
	0x0b, 0xd2, 0xb8, 0x70, 0x55, 0x9a, 0x23, 0x58, 0x94, 0x66, 0x00, 0x4b, 0xd2, 0x48, 0xce, 0x12,
	0xd0, 0x5e, 0xc3, 0x4a, 0xe3, 0x21, 0xa8, 0xe4, 0x5f, 0x05, 0x19, 0xe9, 0x64, 0x14, 0x36, 0xd4,
	0x1a, 0x47, 0xe0, 0xf1, 0x6d, 0xd3, 0x90, 0x89, 0xf7, 0x96, 0xfd, 0x59, 0x41, 0x1b, 0x60, 0x6d,
	0xb1, 0xd2, 0xd0, 0x26, 0xa8, 0x4a, 0xbd, 0x7d, 0xcb, 0xb3, 0x4c, 0x58, 0x68, 0xfc, 0xae, 0x81,
	0xcd, 0x4b, 0x13, 0x20, 0x5b, 0xe5, 0xc8, 0xd3, 0x52, 0x4d, 0xb3, 0x4d, 0xa7, 0x05, 0xb5, 0xdc,
	0x6b, 0xc3, 0x42, 0xe6, 0x75, 0x5a, 0x70, 0x35, 0xf7, 0xda, 0xb0, 0x88, 0xca, 0xa0, 0x60, 0x98,
	0xb0, 0x24, 0xad, 0xd9, 0x83, 0x65, 0x65, 0x8f, 0x60, 0x45, 0xda, 0x9e, 0x3c, 0xf5, 0x32, 0x28,
	0xd8, 0x3d, 0x58, 0x95, 0xd6, 0xed, 0x41, 0x20, 0xed, 0xc8, 0x84, 0xeb, 0x2a, 0x6e, 0xc2, 0x0d,
	0x79, 0x0c, 0xb6, 0xb9, 0x3f, 0x6c, 0xc1, 0xcd, 0x73, 0xb7, 0x0d, 0xef, 0x64, 0xfc, 0x6e, 0x0b,
	0xde, 0xcd, 0xbd, 0x36, 0x84, 0x0d, 0x0b, 0xdc, 0xbb, 0xf6, 0x7e, 0x39, 0x7f, 0xe2, 0xa1, 0xd5,
	0xb7, 0xba, 0x9e, 0x65, 0x65, 0x4f, 0xbc, 0x58, 0x69, 0x32, 0xef, 0x0d, 0x0d, 0xff, 0xd8, 0xb5,
	0xb2, 0x0b, 0xbe, 0xda, 0x78, 0x04, 0xaa, 0x8b, 0xd9, 0x95, 0x7d, 0x75, 0xb2, 0xf6, 0xbe, 0x85,
	0x5a, 0x67, 0x08, 0x1e, 0x32, 0xae, 0xcb, 0xbf, 0x91, 0x53, 0x96, 0xea, 0xc7, 0x87, 0x8c, 0xeb,
	0xfb, 0xd9, 0xb8, 0x7f, 0xb7, 0x2d, 0x52, 0x22, 0x4e, 0x16, 0xb9, 0x80, 0x9f, 0x35, 0x19, 0x6f,
	0x9e, 0xf1, 0x90, 0x46, 0x4d, 0x11, 0x9e, 0x36, 0xa7, 0xbc, 0xb9, 0xfc, 0x11, 0xfb, 0x50, 0xa8,
	0x1e, 0x1f, 0xe6, 0x95, 0xe3, 0xb2, 0xfa, 0x6d, 0xda, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xca,
	0xc0, 0x3a, 0x42, 0xa9, 0x09, 0x00, 0x00,
}
