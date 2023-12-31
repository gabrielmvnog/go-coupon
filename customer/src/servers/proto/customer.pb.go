// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/servers/proto/customer.proto

package customer

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

type CreateCustomerRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerRequest) Reset()         { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()    {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{0}
}

func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRequest.Unmarshal(m, b)
}
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRequest.Merge(m, src)
}
func (m *CreateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRequest.Size(m)
}
func (m *CreateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRequest proto.InternalMessageInfo

func (m *CreateCustomerRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateCustomerRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateCustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateCustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type CreateCustomerResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerResponse) Reset()         { *m = CreateCustomerResponse{} }
func (m *CreateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerResponse) ProtoMessage()    {}
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{1}
}

func (m *CreateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerResponse.Unmarshal(m, b)
}
func (m *CreateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerResponse.Merge(m, src)
}
func (m *CreateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerResponse.Size(m)
}
func (m *CreateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerResponse proto.InternalMessageInfo

func (m *CreateCustomerResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateCustomerRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCustomerRequest) Reset()         { *m = UpdateCustomerRequest{} }
func (m *UpdateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerRequest) ProtoMessage()    {}
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{2}
}

func (m *UpdateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerRequest.Unmarshal(m, b)
}
func (m *UpdateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerRequest.Merge(m, src)
}
func (m *UpdateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerRequest.Size(m)
}
func (m *UpdateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerRequest proto.InternalMessageInfo

func (m *UpdateCustomerRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateCustomerRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateCustomerRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateCustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateCustomerRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type UpdateCustomerResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCustomerResponse) Reset()         { *m = UpdateCustomerResponse{} }
func (m *UpdateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerResponse) ProtoMessage()    {}
func (*UpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{3}
}

func (m *UpdateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerResponse.Unmarshal(m, b)
}
func (m *UpdateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerResponse.Merge(m, src)
}
func (m *UpdateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerResponse.Size(m)
}
func (m *UpdateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerResponse proto.InternalMessageInfo

func (m *UpdateCustomerResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateCustomerResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateCustomerResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateCustomerResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateCustomerResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type GetCustomerRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerRequest) Reset()         { *m = GetCustomerRequest{} }
func (m *GetCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerRequest) ProtoMessage()    {}
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{4}
}

func (m *GetCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerRequest.Unmarshal(m, b)
}
func (m *GetCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerRequest.Merge(m, src)
}
func (m *GetCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerRequest.Size(m)
}
func (m *GetCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerRequest proto.InternalMessageInfo

func (m *GetCustomerRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetCustomerResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerResponse) Reset()         { *m = GetCustomerResponse{} }
func (m *GetCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*GetCustomerResponse) ProtoMessage()    {}
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{5}
}

func (m *GetCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerResponse.Unmarshal(m, b)
}
func (m *GetCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerResponse.Marshal(b, m, deterministic)
}
func (m *GetCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerResponse.Merge(m, src)
}
func (m *GetCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_GetCustomerResponse.Size(m)
}
func (m *GetCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerResponse proto.InternalMessageInfo

func (m *GetCustomerResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetCustomerResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GetCustomerResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *GetCustomerResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetCustomerResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type DeleteCustomerRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerRequest) Reset()         { *m = DeleteCustomerRequest{} }
func (m *DeleteCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerRequest) ProtoMessage()    {}
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{6}
}

func (m *DeleteCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerRequest.Unmarshal(m, b)
}
func (m *DeleteCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerRequest.Merge(m, src)
}
func (m *DeleteCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerRequest.Size(m)
}
func (m *DeleteCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerRequest proto.InternalMessageInfo

func (m *DeleteCustomerRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCustomerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerResponse) Reset()         { *m = DeleteCustomerResponse{} }
func (m *DeleteCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerResponse) ProtoMessage()    {}
func (*DeleteCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8b7abf91856d53, []int{7}
}

func (m *DeleteCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerResponse.Unmarshal(m, b)
}
func (m *DeleteCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerResponse.Merge(m, src)
}
func (m *DeleteCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerResponse.Size(m)
}
func (m *DeleteCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateCustomerRequest)(nil), "CreateCustomerRequest")
	proto.RegisterType((*CreateCustomerResponse)(nil), "CreateCustomerResponse")
	proto.RegisterType((*UpdateCustomerRequest)(nil), "UpdateCustomerRequest")
	proto.RegisterType((*UpdateCustomerResponse)(nil), "UpdateCustomerResponse")
	proto.RegisterType((*GetCustomerRequest)(nil), "GetCustomerRequest")
	proto.RegisterType((*GetCustomerResponse)(nil), "GetCustomerResponse")
	proto.RegisterType((*DeleteCustomerRequest)(nil), "DeleteCustomerRequest")
	proto.RegisterType((*DeleteCustomerResponse)(nil), "DeleteCustomerResponse")
}

func init() {
	proto.RegisterFile("src/servers/proto/customer.proto", fileDescriptor_cc8b7abf91856d53)
}

var fileDescriptor_cc8b7abf91856d53 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x9b, 0xb4, 0x7d, 0xbc, 0xde, 0xc7, 0xab, 0x30, 0x6d, 0xd3, 0x10, 0x11, 0xca, 0x20,
	0x58, 0x44, 0x27, 0xa0, 0x3b, 0x97, 0x56, 0x70, 0xe7, 0xa2, 0xe2, 0xc6, 0x8d, 0xa4, 0xe9, 0x35,
	0x0d, 0x24, 0x99, 0x38, 0x33, 0xe9, 0xd6, 0x5f, 0x20, 0xfa, 0x03, 0xfc, 0xb1, 0xd2, 0x89, 0x95,
	0x26, 0x8e, 0x76, 0xa7, 0xcb, 0x7b, 0xce, 0x25, 0xf9, 0x72, 0x72, 0xb8, 0x30, 0x92, 0x22, 0xf4,
	0x25, 0x8a, 0x25, 0x0a, 0xe9, 0xe7, 0x82, 0x2b, 0xee, 0x87, 0x85, 0x54, 0x3c, 0x45, 0xc1, 0xf4,
	0x48, 0x1f, 0x61, 0x30, 0x11, 0x18, 0x28, 0x9c, 0xbc, 0xeb, 0x53, 0x7c, 0x28, 0x50, 0x2a, 0xb2,
	0x07, 0x70, 0x1f, 0x0b, 0xa9, 0xee, 0xb2, 0x20, 0x45, 0xd7, 0x1a, 0x59, 0xe3, 0xce, 0xb4, 0xa3,
	0x95, 0xab, 0x20, 0x45, 0xb2, 0x0b, 0x9d, 0x24, 0x58, 0xbb, 0xb6, 0x76, 0xff, 0xae, 0x04, 0x6d,
	0xf6, 0xa1, 0x8d, 0x69, 0x10, 0x27, 0x6e, 0x53, 0x1b, 0xe5, 0xb0, 0x52, 0xf3, 0x05, 0xcf, 0xd0,
	0x6d, 0x95, 0xaa, 0x1e, 0xe8, 0x18, 0x9c, 0x3a, 0x80, 0xcc, 0x79, 0x26, 0x91, 0x74, 0xc1, 0x8e,
	0xe7, 0xfa, 0xcd, 0xff, 0xa7, 0x76, 0x3c, 0xa7, 0xcf, 0x16, 0x0c, 0x6e, 0xf2, 0xb9, 0x81, 0xb5,
	0xb6, 0x59, 0x63, 0xb7, 0xbf, 0x65, 0x6f, 0x7e, 0xc5, 0xde, 0x32, 0xb2, 0xb7, 0x37, 0xd9, 0x5f,
	0x2c, 0x70, 0xea, 0x44, 0x66, 0xf8, 0x1f, 0x43, 0xda, 0x07, 0x72, 0x89, 0x6a, 0x4b, 0x40, 0xf4,
	0xc9, 0x82, 0x5e, 0x65, 0xed, 0x97, 0xa9, 0x0f, 0x60, 0x70, 0x81, 0x09, 0x6e, 0xfd, 0xb3, 0xd4,
	0x05, 0xa7, 0xbe, 0x58, 0xa2, 0x9f, 0xbc, 0xda, 0xb0, 0xb3, 0x16, 0xaf, 0x51, 0x2c, 0xe3, 0x10,
	0xc9, 0x04, 0xba, 0xd5, 0x6e, 0x11, 0x87, 0x19, 0xdb, 0xee, 0x0d, 0x99, 0xb9, 0x84, 0xb4, 0x41,
	0xce, 0xe0, 0xdf, 0x46, 0x54, 0xa4, 0xc7, 0x3e, 0xe7, 0xeb, 0xf5, 0x99, 0x21, 0x4d, 0xda, 0x58,
	0x01, 0x54, 0xfb, 0x41, 0x1c, 0x66, 0xac, 0xb0, 0x37, 0x64, 0xe6, 0x22, 0x95, 0x0f, 0xa9, 0x7e,
	0x33, 0x71, 0x98, 0x31, 0x2d, 0x6f, 0xc8, 0xcc, 0xe1, 0xd0, 0xc6, 0xf9, 0xd1, 0xed, 0x61, 0x14,
	0xab, 0x45, 0x31, 0x63, 0x21, 0x4f, 0xfd, 0x28, 0x98, 0x89, 0x18, 0x93, 0x74, 0x99, 0xf1, 0xc8,
	0x8f, 0xf8, 0x71, 0xc8, 0x8b, 0x9c, 0x67, 0x1f, 0xb7, 0x61, 0xf6, 0x47, 0x1f, 0x87, 0xd3, 0xb7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xec, 0x33, 0x92, 0x40, 0x04, 0x00, 0x00,
}
