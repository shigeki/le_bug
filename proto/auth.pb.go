// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type Authorization struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	RegistrationID       int64    `protobuf:"varint,3,opt,name=registrationID,proto3" json:"registrationID,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Expires              int64    `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authorization) Reset()         { *m = Authorization{} }
func (m *Authorization) String() string { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()    {}
func (*Authorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Authorization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorization.Unmarshal(m, b)
}
func (m *Authorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorization.Marshal(b, m, deterministic)
}
func (m *Authorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorization.Merge(m, src)
}
func (m *Authorization) XXX_Size() int {
	return xxx_messageInfo_Authorization.Size(m)
}
func (m *Authorization) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorization.DiscardUnknown(m)
}

var xxx_messageInfo_Authorization proto.InternalMessageInfo

func (m *Authorization) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Authorization) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Authorization) GetRegistrationID() int64 {
	if m != nil {
		return m.RegistrationID
	}
	return 0
}

func (m *Authorization) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Authorization) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type Authorizations struct {
	Authz                []*Authorizations_MapElement `protobuf:"bytes,1,rep,name=authz,proto3" json:"authz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Authorizations) Reset()         { *m = Authorizations{} }
func (m *Authorizations) String() string { return proto.CompactTextString(m) }
func (*Authorizations) ProtoMessage()    {}
func (*Authorizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Authorizations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorizations.Unmarshal(m, b)
}
func (m *Authorizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorizations.Marshal(b, m, deterministic)
}
func (m *Authorizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorizations.Merge(m, src)
}
func (m *Authorizations) XXX_Size() int {
	return xxx_messageInfo_Authorizations.Size(m)
}
func (m *Authorizations) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorizations.DiscardUnknown(m)
}

var xxx_messageInfo_Authorizations proto.InternalMessageInfo

func (m *Authorizations) GetAuthz() []*Authorizations_MapElement {
	if m != nil {
		return m.Authz
	}
	return nil
}

type Authorizations_MapElement struct {
	Domain               string         `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Authz                *Authorization `protobuf:"bytes,2,opt,name=authz,proto3" json:"authz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Authorizations_MapElement) Reset()         { *m = Authorizations_MapElement{} }
func (m *Authorizations_MapElement) String() string { return proto.CompactTextString(m) }
func (*Authorizations_MapElement) ProtoMessage()    {}
func (*Authorizations_MapElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1, 0}
}

func (m *Authorizations_MapElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorizations_MapElement.Unmarshal(m, b)
}
func (m *Authorizations_MapElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorizations_MapElement.Marshal(b, m, deterministic)
}
func (m *Authorizations_MapElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorizations_MapElement.Merge(m, src)
}
func (m *Authorizations_MapElement) XXX_Size() int {
	return xxx_messageInfo_Authorizations_MapElement.Size(m)
}
func (m *Authorizations_MapElement) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorizations_MapElement.DiscardUnknown(m)
}

var xxx_messageInfo_Authorizations_MapElement proto.InternalMessageInfo

func (m *Authorizations_MapElement) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Authorizations_MapElement) GetAuthz() *Authorization {
	if m != nil {
		return m.Authz
	}
	return nil
}

func init() {
	proto.RegisterType((*Authorization)(nil), "Authorization")
	proto.RegisterType((*Authorizations)(nil), "Authorizations")
	proto.RegisterType((*Authorizations_MapElement)(nil), "Authorizations.MapElement")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xdb, 0x6d, 0xeb, 0x2c, 0x96, 0x25, 0x07, 0x09, 0x7b, 0x90, 0xb2, 0x88, 0xf4,
	0x54, 0x64, 0x7d, 0x02, 0x41, 0x0f, 0x16, 0xbc, 0xf4, 0x0d, 0x22, 0x89, 0xee, 0x80, 0x9b, 0x94,
	0x64, 0x0a, 0xb2, 0x0f, 0xe1, 0x5b, 0xf8, 0x9e, 0xd2, 0x98, 0xa2, 0xdd, 0xe3, 0xff, 0x25, 0xf3,
	0x31, 0xff, 0x00, 0xc8, 0x91, 0x0e, 0xed, 0xe0, 0x2c, 0xd9, 0xdd, 0x37, 0x83, 0xcb, 0x87, 0x91,
	0x0e, 0xd6, 0xe1, 0x49, 0x12, 0x5a, 0xc3, 0x2b, 0x48, 0x50, 0x09, 0x56, 0xb3, 0xe6, 0xa2, 0x4f,
	0x50, 0xf1, 0x6b, 0x00, 0x54, 0xda, 0x10, 0xbe, 0xa1, 0x76, 0x22, 0x09, 0xfc, 0x1f, 0xe1, 0xb7,
	0x50, 0x39, 0xfd, 0x8e, 0x9e, 0x5c, 0x98, 0x7f, 0x7e, 0x14, 0x69, 0xcd, 0x9a, 0xb4, 0x3f, 0xa3,
	0xfc, 0x0a, 0x72, 0x4f, 0x92, 0x46, 0x2f, 0xb2, 0xe0, 0x88, 0x89, 0x0b, 0x28, 0xf4, 0xe7, 0x80,
	0x4e, 0x7b, 0xb1, 0x0a, 0x83, 0x73, 0xec, 0xb2, 0x32, 0xdf, 0x14, 0x5d, 0x56, 0x16, 0x9b, 0x72,
	0xf7, 0xc5, 0xa0, 0x5a, 0xec, 0xe9, 0xf9, 0x1d, 0xac, 0xa6, 0x22, 0x27, 0xc1, 0xea, 0xb4, 0x59,
	0xef, 0xb7, 0xed, 0xf2, 0xbd, 0x7d, 0x91, 0xc3, 0xd3, 0x87, 0x3e, 0x6a, 0x43, 0xfd, 0xef, 0xc7,
	0x6d, 0x07, 0xf0, 0x07, 0xa7, 0x85, 0x94, 0x3d, 0x4a, 0x34, 0xb1, 0x6c, 0x4c, 0xfc, 0x66, 0xf6,
	0x4e, 0x5d, 0xd7, 0xfb, 0x6a, 0xe9, 0x8d, 0xae, 0xd7, 0x3c, 0xdc, 0xef, 0xfe, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0x51, 0x5f, 0xc1, 0x4d, 0x01, 0x00, 0x00,
}
