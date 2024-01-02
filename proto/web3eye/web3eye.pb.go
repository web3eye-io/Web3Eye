// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/web3eye.proto

package web3eye

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type FilterCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op  string          `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Val *structpb.Value `protobuf:"bytes,20,opt,name=Val,proto3" json:"Val,omitempty"`
}

func (x *FilterCond) Reset() {
	*x = FilterCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterCond) ProtoMessage() {}

func (x *FilterCond) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterCond.ProtoReflect.Descriptor instead.
func (*FilterCond) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{1}
}

func (x *FilterCond) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *FilterCond) GetVal() *structpb.Value {
	if x != nil {
		return x.Val
	}
	return nil
}

type Int32Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value int32  `protobuf:"varint,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Int32Val) Reset() {
	*x = Int32Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Val) ProtoMessage() {}

func (x *Int32Val) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Val.ProtoReflect.Descriptor instead.
func (*Int32Val) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{2}
}

func (x *Int32Val) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Int32Val) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Int64Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value int64  `protobuf:"varint,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Int64Val) Reset() {
	*x = Int64Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Val) ProtoMessage() {}

func (x *Int64Val) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Val.ProtoReflect.Descriptor instead.
func (*Int64Val) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{3}
}

func (x *Int64Val) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Int64Val) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Uint32Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value uint32 `protobuf:"varint,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Uint32Val) Reset() {
	*x = Uint32Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint32Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32Val) ProtoMessage() {}

func (x *Uint32Val) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint32Val.ProtoReflect.Descriptor instead.
func (*Uint32Val) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{4}
}

func (x *Uint32Val) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Uint32Val) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Uint64Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value uint64 `protobuf:"varint,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Uint64Val) Reset() {
	*x = Uint64Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint64Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint64Val) ProtoMessage() {}

func (x *Uint64Val) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint64Val.ProtoReflect.Descriptor instead.
func (*Uint64Val) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{5}
}

func (x *Uint64Val) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Uint64Val) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DoubleVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string  `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value float64 `protobuf:"fixed64,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *DoubleVal) Reset() {
	*x = DoubleVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleVal) ProtoMessage() {}

func (x *DoubleVal) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleVal.ProtoReflect.Descriptor instead.
func (*DoubleVal) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{6}
}

func (x *DoubleVal) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *DoubleVal) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type StringVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value string `protobuf:"bytes,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *StringVal) Reset() {
	*x = StringVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringVal) ProtoMessage() {}

func (x *StringVal) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringVal.ProtoReflect.Descriptor instead.
func (*StringVal) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{7}
}

func (x *StringVal) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *StringVal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type BoolVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value bool   `protobuf:"varint,20,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *BoolVal) Reset() {
	*x = BoolVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolVal) ProtoMessage() {}

func (x *BoolVal) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolVal.ProtoReflect.Descriptor instead.
func (*BoolVal) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{8}
}

func (x *BoolVal) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *BoolVal) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type StringSliceVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string   `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value []string `protobuf:"bytes,20,rep,name=Value,proto3" json:"Value,omitempty"`
}

func (x *StringSliceVal) Reset() {
	*x = StringSliceVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringSliceVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringSliceVal) ProtoMessage() {}

func (x *StringSliceVal) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringSliceVal.ProtoReflect.Descriptor instead.
func (*StringSliceVal) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{9}
}

func (x *StringSliceVal) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *StringSliceVal) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int64SliceVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string  `protobuf:"bytes,10,opt,name=Op,proto3" json:"Op,omitempty"`
	Value []int64 `protobuf:"varint,20,rep,packed,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Int64SliceVal) Reset() {
	*x = Int64SliceVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_web3eye_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64SliceVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64SliceVal) ProtoMessage() {}

func (x *Int64SliceVal) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_web3eye_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64SliceVal.ProtoReflect.Descriptor instead.
func (*Int64SliceVal) Descriptor() ([]byte, []int) {
	return file_web3eye_web3eye_proto_rawDescGZIP(), []int{10}
}

func (x *Int64SliceVal) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Int64SliceVal) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_web3eye_web3eye_proto protoreflect.FileDescriptor

var file_web3eye_web3eye_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x4f, 0x70, 0x12, 0x28, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x56, 0x61, 0x6c, 0x22, 0x30, 0x0a,
	0x08, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x30, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f,
	0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x31, 0x0a, 0x09, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2f, 0x0a,
	0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33,
	0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_web3eye_proto_rawDescOnce sync.Once
	file_web3eye_web3eye_proto_rawDescData = file_web3eye_web3eye_proto_rawDesc
)

func file_web3eye_web3eye_proto_rawDescGZIP() []byte {
	file_web3eye_web3eye_proto_rawDescOnce.Do(func() {
		file_web3eye_web3eye_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_web3eye_proto_rawDescData)
	})
	return file_web3eye_web3eye_proto_rawDescData
}

var file_web3eye_web3eye_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_web3eye_web3eye_proto_goTypes = []interface{}{
	(*VersionResponse)(nil), // 0: web3eye.VersionResponse
	(*FilterCond)(nil),      // 1: web3eye.FilterCond
	(*Int32Val)(nil),        // 2: web3eye.Int32Val
	(*Int64Val)(nil),        // 3: web3eye.Int64Val
	(*Uint32Val)(nil),       // 4: web3eye.Uint32Val
	(*Uint64Val)(nil),       // 5: web3eye.Uint64Val
	(*DoubleVal)(nil),       // 6: web3eye.DoubleVal
	(*StringVal)(nil),       // 7: web3eye.StringVal
	(*BoolVal)(nil),         // 8: web3eye.BoolVal
	(*StringSliceVal)(nil),  // 9: web3eye.StringSliceVal
	(*Int64SliceVal)(nil),   // 10: web3eye.Int64SliceVal
	(*structpb.Value)(nil),  // 11: google.protobuf.Value
}
var file_web3eye_web3eye_proto_depIdxs = []int32{
	11, // 0: web3eye.FilterCond.Val:type_name -> google.protobuf.Value
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_web3eye_web3eye_proto_init() }
func file_web3eye_web3eye_proto_init() {
	if File_web3eye_web3eye_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_web3eye_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_web3eye_web3eye_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterCond); i {
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
		file_web3eye_web3eye_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Val); i {
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
		file_web3eye_web3eye_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Val); i {
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
		file_web3eye_web3eye_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint32Val); i {
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
		file_web3eye_web3eye_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint64Val); i {
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
		file_web3eye_web3eye_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleVal); i {
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
		file_web3eye_web3eye_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringVal); i {
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
		file_web3eye_web3eye_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolVal); i {
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
		file_web3eye_web3eye_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringSliceVal); i {
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
		file_web3eye_web3eye_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64SliceVal); i {
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
			RawDescriptor: file_web3eye_web3eye_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web3eye_web3eye_proto_goTypes,
		DependencyIndexes: file_web3eye_web3eye_proto_depIdxs,
		MessageInfos:      file_web3eye_web3eye_proto_msgTypes,
	}.Build()
	File_web3eye_web3eye_proto = out.File
	file_web3eye_web3eye_proto_rawDesc = nil
	file_web3eye_web3eye_proto_goTypes = nil
	file_web3eye_web3eye_proto_depIdxs = nil
}
