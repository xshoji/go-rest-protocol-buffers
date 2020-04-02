// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package usersapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Country int32

const (
	Country_UNKNOWN Country = 0
	Country_US      Country = 1
	Country_JP      Country = 2
)

var Country_name = map[int32]string{
	0: "UNKNOWN",
	1: "US",
	2: "JP",
}

var Country_value = map[string]int32{
	"UNKNOWN": 0,
	"US":      1,
	"JP":      2,
}

func (x Country) String() string {
	return proto.EnumName(Country_name, int32(x))
}

func (Country) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

type UserServiceResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Users                []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceResponse) Reset()         { *m = UserServiceResponse{} }
func (m *UserServiceResponse) String() string { return proto.CompactTextString(m) }
func (*UserServiceResponse) ProtoMessage()    {}
func (*UserServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *UserServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceResponse.Unmarshal(m, b)
}
func (m *UserServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceResponse.Marshal(b, m, deterministic)
}
func (m *UserServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceResponse.Merge(m, src)
}
func (m *UserServiceResponse) XXX_Size() int {
	return xxx_messageInfo_UserServiceResponse.Size(m)
}
func (m *UserServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceResponse proto.InternalMessageInfo

func (m *UserServiceResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserServiceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserServiceResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UserServiceResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to NicknameOptional:
	//	*User_Nickname
	NicknameOptional isUser_NicknameOptional `protobuf_oneof:"nickname_optional"`
	// Types that are valid to be assigned to AgeOptional:
	//	*User_Age
	AgeOptional          isUser_AgeOptional `protobuf_oneof:"age_optional"`
	Birth                *Birth             `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	Addresss             *Address           `protobuf:"bytes,6,opt,name=addresss,proto3" json:"addresss,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isUser_NicknameOptional interface {
	isUser_NicknameOptional()
}

type User_Nickname struct {
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3,oneof"`
}

func (*User_Nickname) isUser_NicknameOptional() {}

func (m *User) GetNicknameOptional() isUser_NicknameOptional {
	if m != nil {
		return m.NicknameOptional
	}
	return nil
}

func (m *User) GetNickname() string {
	if x, ok := m.GetNicknameOptional().(*User_Nickname); ok {
		return x.Nickname
	}
	return ""
}

type isUser_AgeOptional interface {
	isUser_AgeOptional()
}

type User_Age struct {
	Age int32 `protobuf:"varint,4,opt,name=age,proto3,oneof"`
}

func (*User_Age) isUser_AgeOptional() {}

func (m *User) GetAgeOptional() isUser_AgeOptional {
	if m != nil {
		return m.AgeOptional
	}
	return nil
}

func (m *User) GetAge() int32 {
	if x, ok := m.GetAgeOptional().(*User_Age); ok {
		return x.Age
	}
	return 0
}

func (m *User) GetBirth() *Birth {
	if m != nil {
		return m.Birth
	}
	return nil
}

func (m *User) GetAddresss() *Address {
	if m != nil {
		return m.Addresss
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*User) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*User_Nickname)(nil),
		(*User_Age)(nil),
	}
}

type Birth struct {
	Datetime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	// Types that are valid to be assigned to WeightOptional:
	//	*Birth_Weight
	WeightOptional isBirth_WeightOptional `protobuf_oneof:"weight_optional"`
	// Types that are valid to be assigned to HospitalOptional:
	//	*Birth_Hospital
	HospitalOptional     isBirth_HospitalOptional `protobuf_oneof:"hospital_optional"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Birth) Reset()         { *m = Birth{} }
func (m *Birth) String() string { return proto.CompactTextString(m) }
func (*Birth) ProtoMessage()    {}
func (*Birth) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *Birth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Birth.Unmarshal(m, b)
}
func (m *Birth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Birth.Marshal(b, m, deterministic)
}
func (m *Birth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Birth.Merge(m, src)
}
func (m *Birth) XXX_Size() int {
	return xxx_messageInfo_Birth.Size(m)
}
func (m *Birth) XXX_DiscardUnknown() {
	xxx_messageInfo_Birth.DiscardUnknown(m)
}

var xxx_messageInfo_Birth proto.InternalMessageInfo

func (m *Birth) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

type isBirth_WeightOptional interface {
	isBirth_WeightOptional()
}

type Birth_Weight struct {
	Weight int32 `protobuf:"varint,2,opt,name=weight,proto3,oneof"`
}

func (*Birth_Weight) isBirth_WeightOptional() {}

func (m *Birth) GetWeightOptional() isBirth_WeightOptional {
	if m != nil {
		return m.WeightOptional
	}
	return nil
}

func (m *Birth) GetWeight() int32 {
	if x, ok := m.GetWeightOptional().(*Birth_Weight); ok {
		return x.Weight
	}
	return 0
}

type isBirth_HospitalOptional interface {
	isBirth_HospitalOptional()
}

type Birth_Hospital struct {
	Hospital string `protobuf:"bytes,3,opt,name=hospital,proto3,oneof"`
}

func (*Birth_Hospital) isBirth_HospitalOptional() {}

func (m *Birth) GetHospitalOptional() isBirth_HospitalOptional {
	if m != nil {
		return m.HospitalOptional
	}
	return nil
}

func (m *Birth) GetHospital() string {
	if x, ok := m.GetHospitalOptional().(*Birth_Hospital); ok {
		return x.Hospital
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Birth) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Birth_Weight)(nil),
		(*Birth_Hospital)(nil),
	}
}

type Address struct {
	Country Country `protobuf:"varint,1,opt,name=country,proto3,enum=usersapi.Country" json:"country,omitempty"`
	// Types that are valid to be assigned to ZipCodeOptional:
	//	*Address_ZipCode
	ZipCodeOptional isAddress_ZipCodeOptional `protobuf_oneof:"zipCode_optional"`
	// Types that are valid to be assigned to StateOptional:
	//	*Address_State
	StateOptional isAddress_StateOptional `protobuf_oneof:"state_optional"`
	// Types that are valid to be assigned to CityOptional:
	//	*Address_City
	CityOptional         isAddress_CityOptional `protobuf_oneof:"city_optional"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{3}
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

func (m *Address) GetCountry() Country {
	if m != nil {
		return m.Country
	}
	return Country_UNKNOWN
}

type isAddress_ZipCodeOptional interface {
	isAddress_ZipCodeOptional()
}

type Address_ZipCode struct {
	ZipCode int32 `protobuf:"varint,2,opt,name=zipCode,proto3,oneof"`
}

func (*Address_ZipCode) isAddress_ZipCodeOptional() {}

func (m *Address) GetZipCodeOptional() isAddress_ZipCodeOptional {
	if m != nil {
		return m.ZipCodeOptional
	}
	return nil
}

func (m *Address) GetZipCode() int32 {
	if x, ok := m.GetZipCodeOptional().(*Address_ZipCode); ok {
		return x.ZipCode
	}
	return 0
}

type isAddress_StateOptional interface {
	isAddress_StateOptional()
}

type Address_State struct {
	State string `protobuf:"bytes,3,opt,name=state,proto3,oneof"`
}

func (*Address_State) isAddress_StateOptional() {}

func (m *Address) GetStateOptional() isAddress_StateOptional {
	if m != nil {
		return m.StateOptional
	}
	return nil
}

func (m *Address) GetState() string {
	if x, ok := m.GetStateOptional().(*Address_State); ok {
		return x.State
	}
	return ""
}

type isAddress_CityOptional interface {
	isAddress_CityOptional()
}

type Address_City struct {
	City string `protobuf:"bytes,4,opt,name=city,proto3,oneof"`
}

func (*Address_City) isAddress_CityOptional() {}

func (m *Address) GetCityOptional() isAddress_CityOptional {
	if m != nil {
		return m.CityOptional
	}
	return nil
}

func (m *Address) GetCity() string {
	if x, ok := m.GetCityOptional().(*Address_City); ok {
		return x.City
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_ZipCode)(nil),
		(*Address_State)(nil),
		(*Address_City)(nil),
	}
}

type UserServiceSelector struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceSelector) Reset()         { *m = UserServiceSelector{} }
func (m *UserServiceSelector) String() string { return proto.CompactTextString(m) }
func (*UserServiceSelector) ProtoMessage()    {}
func (*UserServiceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{4}
}

func (m *UserServiceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceSelector.Unmarshal(m, b)
}
func (m *UserServiceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceSelector.Marshal(b, m, deterministic)
}
func (m *UserServiceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceSelector.Merge(m, src)
}
func (m *UserServiceSelector) XXX_Size() int {
	return xxx_messageInfo_UserServiceSelector.Size(m)
}
func (m *UserServiceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceSelector proto.InternalMessageInfo

func (m *UserServiceSelector) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterEnum("usersapi.Country", Country_name, Country_value)
	proto.RegisterType((*UserServiceResponse)(nil), "usersapi.UserServiceResponse")
	proto.RegisterType((*User)(nil), "usersapi.User")
	proto.RegisterType((*Birth)(nil), "usersapi.Birth")
	proto.RegisterType((*Address)(nil), "usersapi.Address")
	proto.RegisterType((*UserServiceSelector)(nil), "usersapi.UserServiceSelector")
}

func init() {
	proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5)
}

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0xae, 0x93, 0xd8, 0x4e, 0x26, 0xe7, 0xb8, 0xce, 0xa4, 0x8a, 0x7c, 0x72, 0x7a, 0x74, 0xa2,
	0x88, 0xa2, 0xa8, 0x88, 0x44, 0x04, 0x89, 0x0b, 0xee, 0x9a, 0x82, 0x14, 0x81, 0x28, 0x68, 0x43,
	0xc4, 0x0d, 0x02, 0x6d, 0xe3, 0x25, 0x5d, 0xe1, 0xd8, 0x96, 0x77, 0x53, 0x54, 0x10, 0x37, 0xf0,
	0x08, 0x3c, 0x00, 0x8f, 0xc1, 0x03, 0xf0, 0x08, 0x5c, 0xf0, 0x02, 0x3c, 0x08, 0xda, 0xf5, 0x5f,
	0x9b, 0x22, 0x55, 0x42, 0x5c, 0x65, 0xbf, 0x99, 0x6f, 0xbf, 0xf9, 0x76, 0x66, 0x1c, 0x70, 0xe3,
	0x24, 0x92, 0xd1, 0x68, 0x2d, 0x58, 0x32, 0xd4, 0x47, 0xac, 0xab, 0xb3, 0xa0, 0x31, 0xef, 0xee,
	0x2e, 0xa3, 0x68, 0x19, 0xb0, 0x11, 0x8d, 0xf9, 0x88, 0x86, 0x61, 0x24, 0xa9, 0xe4, 0x51, 0x28,
	0x52, 0x5e, 0xf7, 0xdf, 0x2c, 0xab, 0xd1, 0xf1, 0xfa, 0xd5, 0x88, 0xad, 0x62, 0x79, 0x96, 0x25,
	0xff, 0xdf, 0x4c, 0x4a, 0xbe, 0x62, 0x42, 0xd2, 0x55, 0x9c, 0x12, 0xfa, 0x1f, 0x0d, 0x68, 0xcf,
	0x05, 0x4b, 0x66, 0x2c, 0x39, 0xe5, 0x0b, 0x46, 0x98, 0x88, 0xa3, 0x50, 0x30, 0xec, 0x80, 0x25,
	0x24, 0x95, 0x6b, 0xe1, 0x19, 0x3d, 0x63, 0x50, 0x25, 0x19, 0x42, 0x0f, 0xec, 0x15, 0x13, 0x82,
	0x2e, 0x99, 0x57, 0xe9, 0x19, 0x83, 0x06, 0xc9, 0x21, 0xee, 0x80, 0xb9, 0x88, 0xd6, 0xa1, 0xf4,
	0xaa, 0xfa, 0x42, 0x0a, 0xf0, 0x1a, 0x98, 0xfa, 0x1d, 0x5e, 0xad, 0x57, 0x1d, 0x34, 0xc7, 0xce,
	0x30, 0x7f, 0xd5, 0x50, 0x55, 0x25, 0x69, 0xb2, 0xff, 0xdd, 0x80, 0x9a, 0xc2, 0xe8, 0x40, 0x85,
	0xfb, 0xba, 0xa4, 0x49, 0x2a, 0xdc, 0x47, 0x84, 0x5a, 0x48, 0x57, 0x79, 0x2d, 0x7d, 0xc6, 0x5d,
	0xa8, 0x87, 0x7c, 0xf1, 0x5a, 0xc7, 0x55, 0xad, 0xc6, 0x74, 0x8b, 0x14, 0x11, 0x44, 0xa8, 0x2a,
	0x73, 0x35, 0x25, 0x31, 0x35, 0x88, 0x02, 0xb8, 0x07, 0xe6, 0x31, 0x4f, 0xe4, 0x89, 0x67, 0xf6,
	0x8c, 0x41, 0x73, 0xbc, 0x5d, 0x9a, 0x98, 0xa8, 0x30, 0x49, 0xb3, 0x78, 0x13, 0xea, 0xd4, 0xf7,
	0x13, 0x26, 0x84, 0xf0, 0x2c, 0xcd, 0x6c, 0x95, 0xcc, 0x83, 0x34, 0x43, 0x0a, 0xca, 0xa4, 0x0d,
	0xad, 0xbc, 0xea, 0xcb, 0x28, 0x56, 0x23, 0xa1, 0xc1, 0xc4, 0x81, 0xbf, 0xe8, 0xb2, 0xc4, 0xfd,
	0xcf, 0x06, 0x98, 0xba, 0x08, 0xde, 0x81, 0xba, 0x4f, 0x25, 0x53, 0x03, 0xd0, 0x0f, 0x6c, 0x8e,
	0xbb, 0xc3, 0x74, 0x3a, 0xc3, 0x7c, 0x3a, 0xc3, 0xa7, 0xf9, 0x74, 0x48, 0xc1, 0x45, 0x0f, 0xac,
	0x37, 0x8c, 0x2f, 0x4f, 0xa4, 0x6e, 0x82, 0x39, 0xdd, 0x22, 0x19, 0x56, 0x8d, 0x38, 0x89, 0x44,
	0xcc, 0x25, 0x0d, 0xb2, 0x46, 0x18, 0xa4, 0x88, 0x4c, 0x5a, 0xb0, 0x9d, 0xf2, 0x4a, 0x73, 0x6d,
	0x68, 0xe5, 0xe9, 0xd2, 0xe1, 0x17, 0x03, 0xec, 0xec, 0x71, 0x78, 0x03, 0x6c, 0x3d, 0xb6, 0xe4,
	0x4c, 0x5b, 0x74, 0xce, 0x37, 0xe0, 0x30, 0x4d, 0x90, 0x9c, 0x81, 0x5d, 0xb0, 0xdf, 0xf2, 0xf8,
	0x30, 0xf2, 0x59, 0xe1, 0x2c, 0x0f, 0x60, 0x07, 0x4c, 0xb5, 0x30, 0xac, 0xf0, 0x95, 0x42, 0xdc,
	0x81, 0xda, 0x82, 0xcb, 0x33, 0x3d, 0x9e, 0xc6, 0xb4, 0x42, 0x34, 0x9a, 0x20, 0xb8, 0xd9, 0xc5,
	0xd2, 0xab, 0x0b, 0x8e, 0xbe, 0x52, 0x46, 0xb6, 0xe1, 0x6f, 0xc5, 0x2e, 0x9d, 0xef, 0x5d, 0x58,
	0xdd, 0x19, 0x0b, 0xd8, 0x42, 0x46, 0x97, 0x76, 0x68, 0xff, 0x3a, 0xd8, 0x99, 0x77, 0x6c, 0x82,
	0x3d, 0x3f, 0x7a, 0x78, 0xf4, 0xf8, 0xd9, 0x91, 0xbb, 0x85, 0x16, 0x54, 0xe6, 0x33, 0xd7, 0x50,
	0xbf, 0x0f, 0x9e, 0xb8, 0x95, 0xf1, 0xd7, 0x2a, 0x34, 0xcf, 0xe9, 0xe1, 0x23, 0xb0, 0x0e, 0x13,
	0xa6, 0x5c, 0x6f, 0x6c, 0x6d, 0xf7, 0xbf, 0x8b, 0x78, 0xe3, 0xdb, 0xe9, 0xef, 0x7c, 0xf8, 0xf6,
	0xe3, 0x53, 0xc5, 0xe9, 0x37, 0x46, 0xa7, 0xb7, 0xf4, 0x17, 0x2d, 0xee, 0x1a, 0xfb, 0xf8, 0x1c,
	0x6a, 0x84, 0x51, 0x1f, 0x7f, 0x7d, 0x39, 0x77, 0x7f, 0x95, 0x76, 0x47, 0x6b, 0xbb, 0xe8, 0x14,
	0xda, 0xa3, 0x77, 0xdc, 0x7f, 0x8f, 0x33, 0xb0, 0x95, 0xfa, 0x41, 0x10, 0x60, 0xe7, 0xd2, 0x5a,
	0xdd, 0x57, 0xff, 0x08, 0x57, 0x29, 0xb7, 0xb4, 0x72, 0x13, 0x4b, 0xd7, 0x48, 0xc0, 0x9a, 0xc7,
	0xfe, 0x6f, 0x74, 0xe0, 0x1f, 0xad, 0xd5, 0xee, 0x6f, 0xb8, 0x54, 0x6d, 0x78, 0x01, 0xd6, 0x3d,
	0x16, 0x30, 0xc9, 0xfe, 0x4c, 0x23, 0xf6, 0x37, 0x4a, 0x1c, 0x5b, 0xfa, 0xd5, 0xb7, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x5a, 0x63, 0x0e, 0x51, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Read(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error)
	ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Delete(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/usersapi.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/usersapi.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/usersapi.UserService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/usersapi.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/usersapi.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *User) (*UserServiceResponse, error)
	Read(context.Context, *UserServiceSelector) (*UserServiceResponse, error)
	ReadAll(context.Context, *empty.Empty) (*UserServiceResponse, error)
	Update(context.Context, *User) (*UserServiceResponse, error)
	Delete(context.Context, *UserServiceSelector) (*UserServiceResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *User) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Read(ctx context.Context, req *UserServiceSelector) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedUserServiceServer) ReadAll(ctx context.Context, req *empty.Empty) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *User) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *UserServiceSelector) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usersapi.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usersapi.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*UserServiceSelector))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usersapi.UserService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usersapi.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usersapi.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserServiceSelector))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usersapi.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _UserService_ReadAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}