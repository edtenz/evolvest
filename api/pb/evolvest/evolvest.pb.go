// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: evolvest.proto

package evolvest

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetResponse) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{2}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ExistVal bool   `protobuf:"varint,2,opt,name=existVal,proto3" json:"existVal,omitempty"`
	OldVal   []byte `protobuf:"bytes,3,opt,name=oldVal,proto3" json:"oldVal,omitempty"`
	NewVal   []byte `protobuf:"bytes,4,opt,name=newVal,proto3" json:"newVal,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{3}
}

func (x *SetResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetResponse) GetExistVal() bool {
	if x != nil {
		return x.ExistVal
	}
	return false
}

func (x *SetResponse) GetOldVal() []byte {
	if x != nil {
		return x.OldVal
	}
	return nil
}

func (x *SetResponse) GetNewVal() []byte {
	if x != nil {
		return x.NewVal
	}
	return nil
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{4}
}

func (x *DelRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DelResponse) Reset() {
	*x = DelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResponse) ProtoMessage() {}

func (x *DelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResponse.ProtoReflect.Descriptor instead.
func (*DelResponse) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{5}
}

func (x *DelResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DelResponse) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{6}
}

type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []byte `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{7}
}

func (x *SyncResponse) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxCmds []string `protobuf:"bytes,1,rep,name=txCmds,proto3" json:"txCmds,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{8}
}

func (x *PushRequest) GetTxCmds() []string {
	if x != nil {
		return x.TxCmds
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evolvest_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evolvest_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_evolvest_proto_rawDescGZIP(), []int{9}
}

func (x *PushResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_evolvest_proto protoreflect.FileDescriptor

var file_evolvest_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x30, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x6b, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x6c,
	0x64, 0x56, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x22, 0x1e, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22,
	0x0d, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26,
	0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x43, 0x6d, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x43, 0x6d, 0x64, 0x73, 0x22, 0x1e, 0x0a,
	0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xa5, 0x02,
	0x0a, 0x0f, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14,
	0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x03, 0x44, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x6f,
	0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x15, 0x2e, 0x65, 0x76,
	0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04,
	0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x2e, 0x65, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x76,
	0x6f, 0x6c, 0x76, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x65, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evolvest_proto_rawDescOnce sync.Once
	file_evolvest_proto_rawDescData = file_evolvest_proto_rawDesc
)

func file_evolvest_proto_rawDescGZIP() []byte {
	file_evolvest_proto_rawDescOnce.Do(func() {
		file_evolvest_proto_rawDescData = protoimpl.X.CompressGZIP(file_evolvest_proto_rawDescData)
	})
	return file_evolvest_proto_rawDescData
}

var file_evolvest_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_evolvest_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: evolvest.GetRequest
	(*GetResponse)(nil),  // 1: evolvest.GetResponse
	(*SetRequest)(nil),   // 2: evolvest.SetRequest
	(*SetResponse)(nil),  // 3: evolvest.SetResponse
	(*DelRequest)(nil),   // 4: evolvest.DelRequest
	(*DelResponse)(nil),  // 5: evolvest.DelResponse
	(*SyncRequest)(nil),  // 6: evolvest.SyncRequest
	(*SyncResponse)(nil), // 7: evolvest.SyncResponse
	(*PushRequest)(nil),  // 8: evolvest.PushRequest
	(*PushResponse)(nil), // 9: evolvest.PushResponse
}
var file_evolvest_proto_depIdxs = []int32{
	0, // 0: evolvest.EvolvestService.Get:input_type -> evolvest.GetRequest
	2, // 1: evolvest.EvolvestService.Set:input_type -> evolvest.SetRequest
	4, // 2: evolvest.EvolvestService.Del:input_type -> evolvest.DelRequest
	6, // 3: evolvest.EvolvestService.Sync:input_type -> evolvest.SyncRequest
	8, // 4: evolvest.EvolvestService.Push:input_type -> evolvest.PushRequest
	1, // 5: evolvest.EvolvestService.Get:output_type -> evolvest.GetResponse
	3, // 6: evolvest.EvolvestService.Set:output_type -> evolvest.SetResponse
	5, // 7: evolvest.EvolvestService.Del:output_type -> evolvest.DelResponse
	7, // 8: evolvest.EvolvestService.Sync:output_type -> evolvest.SyncResponse
	9, // 9: evolvest.EvolvestService.Push:output_type -> evolvest.PushResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_evolvest_proto_init() }
func file_evolvest_proto_init() {
	if File_evolvest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evolvest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_evolvest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_evolvest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_evolvest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_evolvest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_evolvest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelResponse); i {
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
		file_evolvest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_evolvest_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
		file_evolvest_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_evolvest_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
			RawDescriptor: file_evolvest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evolvest_proto_goTypes,
		DependencyIndexes: file_evolvest_proto_depIdxs,
		MessageInfos:      file_evolvest_proto_msgTypes,
	}.Build()
	File_evolvest_proto = out.File
	file_evolvest_proto_rawDesc = nil
	file_evolvest_proto_goTypes = nil
	file_evolvest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EvolvestServiceClient is the client API for EvolvestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EvolvestServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type evolvestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEvolvestServiceClient(cc grpc.ClientConnInterface) EvolvestServiceClient {
	return &evolvestServiceClient{cc}
}

func (c *evolvestServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/evolvest.EvolvestService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evolvestServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/evolvest.EvolvestService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evolvestServiceClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := c.cc.Invoke(ctx, "/evolvest.EvolvestService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evolvestServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/evolvest.EvolvestService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evolvestServiceClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/evolvest.EvolvestService/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvolvestServiceServer is the server API for EvolvestService service.
type EvolvestServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Del(context.Context, *DelRequest) (*DelResponse, error)
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

// UnimplementedEvolvestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEvolvestServiceServer struct {
}

func (*UnimplementedEvolvestServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEvolvestServiceServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedEvolvestServiceServer) Del(context.Context, *DelRequest) (*DelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedEvolvestServiceServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedEvolvestServiceServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}

func RegisterEvolvestServiceServer(s *grpc.Server, srv EvolvestServiceServer) {
	s.RegisterService(&_EvolvestService_serviceDesc, srv)
}

func _EvolvestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvolvestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evolvest.EvolvestService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvolvestServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvolvestService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvolvestServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evolvest.EvolvestService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvolvestServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvolvestService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvolvestServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evolvest.EvolvestService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvolvestServiceServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvolvestService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvolvestServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evolvest.EvolvestService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvolvestServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvolvestService_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvolvestServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evolvest.EvolvestService/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvolvestServiceServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EvolvestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evolvest.EvolvestService",
	HandlerType: (*EvolvestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EvolvestService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _EvolvestService_Set_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _EvolvestService_Del_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _EvolvestService_Sync_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _EvolvestService_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evolvest.proto",
}
