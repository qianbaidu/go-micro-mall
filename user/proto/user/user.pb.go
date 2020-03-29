// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_srv_user

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

// 用户信息
type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Salt                 string   `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	CreatedUnix          int64    `protobuf:"varint,5,opt,name=created_unix,json=createdUnix,proto3" json:"created_unix,omitempty"`
	UpdatedUnix          int64    `protobuf:"varint,6,opt,name=updated_unix,json=updatedUnix,proto3" json:"updated_unix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *User) GetCreatedUnix() int64 {
	if m != nil {
		return m.CreatedUnix
	}
	return 0
}

func (m *User) GetUpdatedUnix() int64 {
	if m != nil {
		return m.UpdatedUnix
	}
	return 0
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type ToRevokeToken struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToRevokeToken) Reset()         { *m = ToRevokeToken{} }
func (m *ToRevokeToken) String() string { return proto.CompactTextString(m) }
func (*ToRevokeToken) ProtoMessage()    {}
func (*ToRevokeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *ToRevokeToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToRevokeToken.Unmarshal(m, b)
}
func (m *ToRevokeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToRevokeToken.Marshal(b, m, deterministic)
}
func (m *ToRevokeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToRevokeToken.Merge(m, src)
}
func (m *ToRevokeToken) XXX_Size() int {
	return xxx_messageInfo_ToRevokeToken.Size(m)
}
func (m *ToRevokeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ToRevokeToken.DiscardUnknown(m)
}

var xxx_messageInfo_ToRevokeToken proto.InternalMessageInfo

func (m *ToRevokeToken) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToRevokeToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*ToRevokeToken)(nil), "go.micro.srv.user.ToRevokeToken")
	proto.RegisterType((*Users)(nil), "go.micro.srv.user.Users")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x19, 0xdb, 0x10, 0x1e, 0x6a, 0x62, 0x63, 0xe2, 0x32, 0x63, 0x82, 0x3d, 0x71, 0x71,
	0x07, 0x8c, 0x26, 0xc6, 0x8b, 0x44, 0x12, 0xe2, 0xcd, 0x4c, 0x38, 0x1b, 0x64, 0x4f, 0x68, 0x98,
	0xed, 0x6c, 0x3b, 0xe4, 0xf3, 0xf8, 0x99, 0xfc, 0x40, 0xa6, 0x1d, 0x20, 0x89, 0x8c, 0x83, 0xde,
	0xf6, 0xde, 0xff, 0xdf, 0xff, 0xdb, 0xaf, 0x2f, 0x05, 0xc8, 0x15, 0xca, 0x28, 0x93, 0x42, 0x0b,
	0x72, 0x34, 0x11, 0xd1, 0x1b, 0x1b, 0x4b, 0x11, 0x29, 0x39, 0x8f, 0x8c, 0x40, 0x3f, 0x1d, 0xf0,
	0x86, 0x0a, 0x25, 0x39, 0x84, 0x2a, 0x4b, 0x02, 0xa7, 0xe5, 0xb4, 0xdd, 0xb8, 0xca, 0x12, 0x72,
	0x0c, 0x7e, 0x36, 0x15, 0x1c, 0x83, 0x6a, 0xcb, 0x69, 0x37, 0xe2, 0xa2, 0x20, 0x21, 0xd4, 0xb3,
	0x91, 0x52, 0x1f, 0x42, 0x26, 0x81, 0x6b, 0x85, 0x75, 0x4d, 0x08, 0x78, 0x6a, 0x94, 0xea, 0xc0,
	0xb3, 0x7d, 0xfb, 0x4d, 0xce, 0x61, 0x7f, 0x2c, 0x71, 0xa4, 0x31, 0x79, 0xce, 0x39, 0x5b, 0x04,
	0xbe, 0xcd, 0x6f, 0x2e, 0x7b, 0x43, 0xce, 0x16, 0xc6, 0x92, 0x67, 0xc9, 0x8f, 0xa5, 0x56, 0x58,
	0x96, 0x3d, 0x63, 0xa1, 0x0d, 0xd8, 0x8b, 0xf1, 0x3d, 0x47, 0xa5, 0x29, 0x40, 0x3d, 0x46, 0x95,
	0x09, 0xae, 0x90, 0x5e, 0xc1, 0xc1, 0x40, 0xc4, 0x38, 0x17, 0x33, 0x1c, 0x88, 0x19, 0xf2, 0x0d,
	0x86, 0xc6, 0x8a, 0x41, 0x1b, 0x61, 0xc5, 0x60, 0x0b, 0x7a, 0x0d, 0xbe, 0x21, 0x56, 0xe4, 0x02,
	0x7c, 0x73, 0x07, 0x2a, 0x70, 0x5a, 0x6e, 0xbb, 0xd9, 0x39, 0x89, 0x7e, 0x5d, 0x4f, 0x64, 0x8c,
	0x71, 0xe1, 0xa2, 0x67, 0xe0, 0x17, 0x63, 0xd6, 0xb1, 0xce, 0x46, 0x6c, 0xe7, 0xcb, 0x85, 0xa6,
	0xb1, 0x3f, 0xa1, 0x9c, 0xb3, 0x31, 0x92, 0x3b, 0xa8, 0xdd, 0x5b, 0x4c, 0x52, 0x16, 0x1c, 0x9e,
	0x6e, 0x11, 0xd6, 0x74, 0x15, 0x93, 0xd0, 0xc3, 0x14, 0xff, 0x91, 0x70, 0x03, 0x6e, 0x1f, 0x75,
	0xf9, 0xf1, 0x32, 0xa1, 0x18, 0xde, 0x47, 0xdd, 0x4d, 0x53, 0x12, 0x6e, 0x9d, 0x61, 0xd7, 0x11,
	0x06, 0x25, 0x01, 0x8a, 0x56, 0x48, 0x0f, 0x60, 0x68, 0x97, 0xf8, 0xc0, 0x5f, 0xc5, 0x9f, 0x11,
	0x6e, 0xc1, 0xeb, 0xe6, 0x7a, 0x5a, 0x7e, 0x7e, 0xdb, 0x2f, 0xd8, 0x3d, 0xd1, 0x0a, 0xe9, 0x82,
	0xf7, 0xc8, 0xf8, 0x64, 0x27, 0xc2, 0xee, 0xf9, 0x2f, 0x35, 0xfb, 0x74, 0x2e, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0x0f, 0x4e, 0x8d, 0x48, 0x03, 0x00, 0x00,
}