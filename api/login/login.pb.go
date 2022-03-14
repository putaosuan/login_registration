// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: api/login/login.proto

package login

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLoginRequest) Reset() {
	*x = CreateLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginRequest) ProtoMessage() {}

func (x *CreateLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginRequest.ProtoReflect.Descriptor instead.
func (*CreateLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{0}
}

type CreateLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLoginReply) Reset() {
	*x = CreateLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginReply) ProtoMessage() {}

func (x *CreateLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginReply.ProtoReflect.Descriptor instead.
func (*CreateLoginReply) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{1}
}

type UpdateLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateLoginRequest) Reset() {
	*x = UpdateLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLoginRequest) ProtoMessage() {}

func (x *UpdateLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLoginRequest.ProtoReflect.Descriptor instead.
func (*UpdateLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{2}
}

type UpdateLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateLoginReply) Reset() {
	*x = UpdateLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLoginReply) ProtoMessage() {}

func (x *UpdateLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLoginReply.ProtoReflect.Descriptor instead.
func (*UpdateLoginReply) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{3}
}

type DeleteLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLoginRequest) Reset() {
	*x = DeleteLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLoginRequest) ProtoMessage() {}

func (x *DeleteLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLoginRequest.ProtoReflect.Descriptor instead.
func (*DeleteLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{4}
}

type DeleteLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLoginReply) Reset() {
	*x = DeleteLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLoginReply) ProtoMessage() {}

func (x *DeleteLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLoginReply.ProtoReflect.Descriptor instead.
func (*DeleteLoginReply) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{5}
}

type GetLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoginRequest) Reset() {
	*x = GetLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginRequest) ProtoMessage() {}

func (x *GetLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginRequest.ProtoReflect.Descriptor instead.
func (*GetLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{6}
}

type GetLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoginReply) Reset() {
	*x = GetLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginReply) ProtoMessage() {}

func (x *GetLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginReply.ProtoReflect.Descriptor instead.
func (*GetLoginReply) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{7}
}

type ListLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLoginRequest) Reset() {
	*x = ListLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoginRequest) ProtoMessage() {}

func (x *ListLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoginRequest.ProtoReflect.Descriptor instead.
func (*ListLoginRequest) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{8}
}

type ListLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLoginReply) Reset() {
	*x = ListLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_login_login_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLoginReply) ProtoMessage() {}

func (x *ListLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_login_login_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLoginReply.ProtoReflect.Descriptor instead.
func (*ListLoginReply) Descriptor() ([]byte, []int) {
	return file_api_login_login_proto_rawDescGZIP(), []int{9}
}

var File_api_login_login_proto protoreflect.FileDescriptor

var file_api_login_login_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xe8, 0x03, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x5d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x12, 0x64, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x1a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x64, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5b, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x57, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x42, 0x31, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x01, 0x5a, 0x22, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_login_login_proto_rawDescOnce sync.Once
	file_api_login_login_proto_rawDescData = file_api_login_login_proto_rawDesc
)

func file_api_login_login_proto_rawDescGZIP() []byte {
	file_api_login_login_proto_rawDescOnce.Do(func() {
		file_api_login_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_login_login_proto_rawDescData)
	})
	return file_api_login_login_proto_rawDescData
}

var file_api_login_login_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_login_login_proto_goTypes = []interface{}{
	(*CreateLoginRequest)(nil), // 0: api.login.CreateLoginRequest
	(*CreateLoginReply)(nil),   // 1: api.login.CreateLoginReply
	(*UpdateLoginRequest)(nil), // 2: api.login.UpdateLoginRequest
	(*UpdateLoginReply)(nil),   // 3: api.login.UpdateLoginReply
	(*DeleteLoginRequest)(nil), // 4: api.login.DeleteLoginRequest
	(*DeleteLoginReply)(nil),   // 5: api.login.DeleteLoginReply
	(*GetLoginRequest)(nil),    // 6: api.login.GetLoginRequest
	(*GetLoginReply)(nil),      // 7: api.login.GetLoginReply
	(*ListLoginRequest)(nil),   // 8: api.login.ListLoginRequest
	(*ListLoginReply)(nil),     // 9: api.login.ListLoginReply
}
var file_api_login_login_proto_depIdxs = []int32{
	0, // 0: api.login.Login.CreateLogin:input_type -> api.login.CreateLoginRequest
	2, // 1: api.login.Login.UpdateLogin:input_type -> api.login.UpdateLoginRequest
	4, // 2: api.login.Login.DeleteLogin:input_type -> api.login.DeleteLoginRequest
	6, // 3: api.login.Login.GetLogin:input_type -> api.login.GetLoginRequest
	8, // 4: api.login.Login.ListLogin:input_type -> api.login.ListLoginRequest
	1, // 5: api.login.Login.CreateLogin:output_type -> api.login.CreateLoginReply
	3, // 6: api.login.Login.UpdateLogin:output_type -> api.login.UpdateLoginReply
	5, // 7: api.login.Login.DeleteLogin:output_type -> api.login.DeleteLoginReply
	7, // 8: api.login.Login.GetLogin:output_type -> api.login.GetLoginReply
	9, // 9: api.login.Login.ListLogin:output_type -> api.login.ListLoginReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_login_login_proto_init() }
func file_api_login_login_proto_init() {
	if File_api_login_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_login_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLoginReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLoginReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_login_login_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLoginReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_login_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_login_login_proto_goTypes,
		DependencyIndexes: file_api_login_login_proto_depIdxs,
		MessageInfos:      file_api_login_login_proto_msgTypes,
	}.Build()
	File_api_login_login_proto = out.File
	file_api_login_login_proto_rawDesc = nil
	file_api_login_login_proto_goTypes = nil
	file_api_login_login_proto_depIdxs = nil
}
