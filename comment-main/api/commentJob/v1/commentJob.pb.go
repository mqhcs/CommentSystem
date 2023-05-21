// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: api/commentJob/v1/commentJob.proto

package v1

import (
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

type CreateCommentJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCommentJobRequest) Reset() {
	*x = CreateCommentJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentJobRequest) ProtoMessage() {}

func (x *CreateCommentJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentJobRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentJobRequest) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{0}
}

type CreateCommentJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCommentJobReply) Reset() {
	*x = CreateCommentJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentJobReply) ProtoMessage() {}

func (x *CreateCommentJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentJobReply.ProtoReflect.Descriptor instead.
func (*CreateCommentJobReply) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{1}
}

type UpdateCommentJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentJobRequest) Reset() {
	*x = UpdateCommentJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentJobRequest) ProtoMessage() {}

func (x *UpdateCommentJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentJobRequest) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{2}
}

type UpdateCommentJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentJobReply) Reset() {
	*x = UpdateCommentJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentJobReply) ProtoMessage() {}

func (x *UpdateCommentJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentJobReply.ProtoReflect.Descriptor instead.
func (*UpdateCommentJobReply) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{3}
}

type DeleteCommentJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentJobRequest) Reset() {
	*x = DeleteCommentJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentJobRequest) ProtoMessage() {}

func (x *DeleteCommentJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentJobRequest) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{4}
}

type DeleteCommentJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentJobReply) Reset() {
	*x = DeleteCommentJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentJobReply) ProtoMessage() {}

func (x *DeleteCommentJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentJobReply.ProtoReflect.Descriptor instead.
func (*DeleteCommentJobReply) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{5}
}

type GetCommentJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommentJobRequest) Reset() {
	*x = GetCommentJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentJobRequest) ProtoMessage() {}

func (x *GetCommentJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentJobRequest.ProtoReflect.Descriptor instead.
func (*GetCommentJobRequest) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{6}
}

type GetCommentJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommentJobReply) Reset() {
	*x = GetCommentJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentJobReply) ProtoMessage() {}

func (x *GetCommentJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentJobReply.ProtoReflect.Descriptor instead.
func (*GetCommentJobReply) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{7}
}

type ListCommentJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCommentJobRequest) Reset() {
	*x = ListCommentJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentJobRequest) ProtoMessage() {}

func (x *ListCommentJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentJobRequest.ProtoReflect.Descriptor instead.
func (*ListCommentJobRequest) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{8}
}

type ListCommentJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCommentJobReply) Reset() {
	*x = ListCommentJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentJobReply) ProtoMessage() {}

func (x *ListCommentJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_commentJob_v1_commentJob_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentJobReply.ProtoReflect.Descriptor instead.
func (*ListCommentJobReply) Descriptor() ([]byte, []int) {
	return file_api_commentJob_v1_commentJob_proto_rawDescGZIP(), []int{9}
}

var File_api_commentJob_v1_commentJob_proto protoreflect.FileDescriptor

var file_api_commentJob_v1_commentJob_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x8f, 0x04, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62,
	0x12, 0x68, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4a, 0x6f, 0x62, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x68, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x68, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x12,
	0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x62, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f,
	0x62, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a,
	0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x38, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x21, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_commentJob_v1_commentJob_proto_rawDescOnce sync.Once
	file_api_commentJob_v1_commentJob_proto_rawDescData = file_api_commentJob_v1_commentJob_proto_rawDesc
)

func file_api_commentJob_v1_commentJob_proto_rawDescGZIP() []byte {
	file_api_commentJob_v1_commentJob_proto_rawDescOnce.Do(func() {
		file_api_commentJob_v1_commentJob_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_commentJob_v1_commentJob_proto_rawDescData)
	})
	return file_api_commentJob_v1_commentJob_proto_rawDescData
}

var file_api_commentJob_v1_commentJob_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_commentJob_v1_commentJob_proto_goTypes = []interface{}{
	(*CreateCommentJobRequest)(nil), // 0: api.commentJob.v1.CreateCommentJobRequest
	(*CreateCommentJobReply)(nil),   // 1: api.commentJob.v1.CreateCommentJobReply
	(*UpdateCommentJobRequest)(nil), // 2: api.commentJob.v1.UpdateCommentJobRequest
	(*UpdateCommentJobReply)(nil),   // 3: api.commentJob.v1.UpdateCommentJobReply
	(*DeleteCommentJobRequest)(nil), // 4: api.commentJob.v1.DeleteCommentJobRequest
	(*DeleteCommentJobReply)(nil),   // 5: api.commentJob.v1.DeleteCommentJobReply
	(*GetCommentJobRequest)(nil),    // 6: api.commentJob.v1.GetCommentJobRequest
	(*GetCommentJobReply)(nil),      // 7: api.commentJob.v1.GetCommentJobReply
	(*ListCommentJobRequest)(nil),   // 8: api.commentJob.v1.ListCommentJobRequest
	(*ListCommentJobReply)(nil),     // 9: api.commentJob.v1.ListCommentJobReply
}
var file_api_commentJob_v1_commentJob_proto_depIdxs = []int32{
	0, // 0: api.commentJob.v1.CommentJob.CreateCommentJob:input_type -> api.commentJob.v1.CreateCommentJobRequest
	2, // 1: api.commentJob.v1.CommentJob.UpdateCommentJob:input_type -> api.commentJob.v1.UpdateCommentJobRequest
	4, // 2: api.commentJob.v1.CommentJob.DeleteCommentJob:input_type -> api.commentJob.v1.DeleteCommentJobRequest
	6, // 3: api.commentJob.v1.CommentJob.GetCommentJob:input_type -> api.commentJob.v1.GetCommentJobRequest
	8, // 4: api.commentJob.v1.CommentJob.ListCommentJob:input_type -> api.commentJob.v1.ListCommentJobRequest
	1, // 5: api.commentJob.v1.CommentJob.CreateCommentJob:output_type -> api.commentJob.v1.CreateCommentJobReply
	3, // 6: api.commentJob.v1.CommentJob.UpdateCommentJob:output_type -> api.commentJob.v1.UpdateCommentJobReply
	5, // 7: api.commentJob.v1.CommentJob.DeleteCommentJob:output_type -> api.commentJob.v1.DeleteCommentJobReply
	7, // 8: api.commentJob.v1.CommentJob.GetCommentJob:output_type -> api.commentJob.v1.GetCommentJobReply
	9, // 9: api.commentJob.v1.CommentJob.ListCommentJob:output_type -> api.commentJob.v1.ListCommentJobReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_commentJob_v1_commentJob_proto_init() }
func file_api_commentJob_v1_commentJob_proto_init() {
	if File_api_commentJob_v1_commentJob_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_commentJob_v1_commentJob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentJobRequest); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentJobReply); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentJobRequest); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentJobReply); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentJobRequest); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentJobReply); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentJobRequest); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentJobReply); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentJobRequest); i {
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
		file_api_commentJob_v1_commentJob_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentJobReply); i {
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
			RawDescriptor: file_api_commentJob_v1_commentJob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_commentJob_v1_commentJob_proto_goTypes,
		DependencyIndexes: file_api_commentJob_v1_commentJob_proto_depIdxs,
		MessageInfos:      file_api_commentJob_v1_commentJob_proto_msgTypes,
	}.Build()
	File_api_commentJob_v1_commentJob_proto = out.File
	file_api_commentJob_v1_commentJob_proto_rawDesc = nil
	file_api_commentJob_v1_commentJob_proto_goTypes = nil
	file_api_commentJob_v1_commentJob_proto_depIdxs = nil
}