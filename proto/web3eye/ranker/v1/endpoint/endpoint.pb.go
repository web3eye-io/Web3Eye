// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/endpoint/endpoint.proto

package endpoint

import (
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	endpoint "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
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

type CreateEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainType v1.ChainType `protobuf:"varint,10,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID   string       `protobuf:"bytes,20,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Address   string       `protobuf:"bytes,30,opt,name=Address,proto3" json:"Address,omitempty"`
	RPS       uint32       `protobuf:"varint,40,opt,name=RPS,proto3" json:"RPS,omitempty"`
}

func (x *CreateEndpointRequest) Reset() {
	*x = CreateEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEndpointRequest) ProtoMessage() {}

func (x *CreateEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEndpointRequest.ProtoReflect.Descriptor instead.
func (*CreateEndpointRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEndpointRequest) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType(0)
}

func (x *CreateEndpointRequest) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *CreateEndpointRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateEndpointRequest) GetRPS() uint32 {
	if x != nil {
		return x.RPS
	}
	return 0
}

type CreateEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *endpoint.Endpoint `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateEndpointResponse) Reset() {
	*x = CreateEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEndpointResponse) ProtoMessage() {}

func (x *CreateEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEndpointResponse.ProtoReflect.Descriptor instead.
func (*CreateEndpointResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEndpointResponse) GetInfo() *endpoint.Endpoint {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32            `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Address *string           `protobuf:"bytes,20,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	State   *v1.EndpointState `protobuf:"varint,30,opt,name=State,proto3,enum=chain.EndpointState,oneof" json:"State,omitempty"`
	RPS     *uint32           `protobuf:"varint,40,opt,name=RPS,proto3,oneof" json:"RPS,omitempty"`
	Remark  *string           `protobuf:"bytes,50,opt,name=Remark,proto3,oneof" json:"Remark,omitempty"`
}

func (x *UpdateEndpointRequest) Reset() {
	*x = UpdateEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEndpointRequest) ProtoMessage() {}

func (x *UpdateEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEndpointRequest.ProtoReflect.Descriptor instead.
func (*UpdateEndpointRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEndpointRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateEndpointRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateEndpointRequest) GetState() v1.EndpointState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return v1.EndpointState(0)
}

func (x *UpdateEndpointRequest) GetRPS() uint32 {
	if x != nil && x.RPS != nil {
		return *x.RPS
	}
	return 0
}

func (x *UpdateEndpointRequest) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

type UpdateEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *endpoint.Endpoint `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateEndpointResponse) Reset() {
	*x = UpdateEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEndpointResponse) ProtoMessage() {}

func (x *UpdateEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEndpointResponse.ProtoReflect.Descriptor instead.
func (*UpdateEndpointResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEndpointResponse) GetInfo() *endpoint.Endpoint {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetEndpointRequest) Reset() {
	*x = GetEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointRequest) ProtoMessage() {}

func (x *GetEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetEndpointRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{4}
}

func (x *GetEndpointRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *endpoint.Endpoint `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetEndpointResponse) Reset() {
	*x = GetEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointResponse) ProtoMessage() {}

func (x *GetEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetEndpointResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{5}
}

func (x *GetEndpointResponse) GetInfo() *endpoint.Endpoint {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetEndpointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        *uint32           `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	ChainType *v1.ChainType     `protobuf:"varint,20,opt,name=ChainType,proto3,enum=chain.ChainType,oneof" json:"ChainType,omitempty"`
	ChainID   *string           `protobuf:"bytes,30,opt,name=ChainID,proto3,oneof" json:"ChainID,omitempty"`
	Address   *string           `protobuf:"bytes,40,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	State     *v1.EndpointState `protobuf:"varint,50,opt,name=State,proto3,enum=chain.EndpointState,oneof" json:"State,omitempty"`
	RPS       *uint32           `protobuf:"varint,60,opt,name=RPS,proto3,oneof" json:"RPS,omitempty"`
	Remark    *string           `protobuf:"bytes,70,opt,name=Remark,proto3,oneof" json:"Remark,omitempty"`
	Offset    int32             `protobuf:"varint,80,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32             `protobuf:"varint,90,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetEndpointsRequest) Reset() {
	*x = GetEndpointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointsRequest) ProtoMessage() {}

func (x *GetEndpointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointsRequest.ProtoReflect.Descriptor instead.
func (*GetEndpointsRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{6}
}

func (x *GetEndpointsRequest) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *GetEndpointsRequest) GetChainType() v1.ChainType {
	if x != nil && x.ChainType != nil {
		return *x.ChainType
	}
	return v1.ChainType(0)
}

func (x *GetEndpointsRequest) GetChainID() string {
	if x != nil && x.ChainID != nil {
		return *x.ChainID
	}
	return ""
}

func (x *GetEndpointsRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *GetEndpointsRequest) GetState() v1.EndpointState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return v1.EndpointState(0)
}

func (x *GetEndpointsRequest) GetRPS() uint32 {
	if x != nil && x.RPS != nil {
		return *x.RPS
	}
	return 0
}

func (x *GetEndpointsRequest) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *GetEndpointsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEndpointsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEndpointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*endpoint.Endpoint `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetEndpointsResponse) Reset() {
	*x = GetEndpointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointsResponse) ProtoMessage() {}

func (x *GetEndpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointsResponse.ProtoReflect.Descriptor instead.
func (*GetEndpointsResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{7}
}

func (x *GetEndpointsResponse) GetInfos() []*endpoint.Endpoint {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetEndpointsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteEndpointRequest) Reset() {
	*x = DeleteEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointRequest) ProtoMessage() {}

func (x *DeleteEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointRequest.ProtoReflect.Descriptor instead.
func (*DeleteEndpointRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEndpointRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *endpoint.Endpoint `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteEndpointResponse) Reset() {
	*x = DeleteEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointResponse) ProtoMessage() {}

func (x *DeleteEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointResponse.ProtoReflect.Descriptor instead.
func (*DeleteEndpointResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEndpointResponse) GetInfo() *endpoint.Endpoint {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_web3eye_ranker_v1_endpoint_endpoint_proto protoreflect.FileDescriptor

var file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a,
	0x2a, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x77, 0x65, 0x62,
	0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x50,
	0x53, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x52, 0x50, 0x53, 0x22, 0x4b, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x01, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x52, 0x50, 0x53, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x02, 0x52, 0x03, 0x52, 0x50, 0x53, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x52, 0x50, 0x53, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x22, 0x4b, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xfa, 0x02,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x48,
	0x01, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x48, 0x04, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15,
	0x0a, 0x03, 0x52, 0x50, 0x53, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x03, 0x52,
	0x50, 0x53, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x88,
	0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x52, 0x50, 0x53, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x27, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x32, 0x91, 0x04, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x69, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x29, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72, 0x61, 0x6e, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f,
	0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77,
	0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescOnce sync.Once
	file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescData = file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDesc
)

func file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescGZIP() []byte {
	file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescOnce.Do(func() {
		file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescData)
	})
	return file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDescData
}

var file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_web3eye_ranker_v1_endpoint_endpoint_proto_goTypes = []interface{}{
	(*CreateEndpointRequest)(nil),  // 0: ranker.v1.endpoint.CreateEndpointRequest
	(*CreateEndpointResponse)(nil), // 1: ranker.v1.endpoint.CreateEndpointResponse
	(*UpdateEndpointRequest)(nil),  // 2: ranker.v1.endpoint.UpdateEndpointRequest
	(*UpdateEndpointResponse)(nil), // 3: ranker.v1.endpoint.UpdateEndpointResponse
	(*GetEndpointRequest)(nil),     // 4: ranker.v1.endpoint.GetEndpointRequest
	(*GetEndpointResponse)(nil),    // 5: ranker.v1.endpoint.GetEndpointResponse
	(*GetEndpointsRequest)(nil),    // 6: ranker.v1.endpoint.GetEndpointsRequest
	(*GetEndpointsResponse)(nil),   // 7: ranker.v1.endpoint.GetEndpointsResponse
	(*DeleteEndpointRequest)(nil),  // 8: ranker.v1.endpoint.DeleteEndpointRequest
	(*DeleteEndpointResponse)(nil), // 9: ranker.v1.endpoint.DeleteEndpointResponse
	(v1.ChainType)(0),              // 10: chain.ChainType
	(*endpoint.Endpoint)(nil),      // 11: nftmeta.v1.endpoint.Endpoint
	(v1.EndpointState)(0),          // 12: chain.EndpointState
}
var file_web3eye_ranker_v1_endpoint_endpoint_proto_depIdxs = []int32{
	10, // 0: ranker.v1.endpoint.CreateEndpointRequest.ChainType:type_name -> chain.ChainType
	11, // 1: ranker.v1.endpoint.CreateEndpointResponse.Info:type_name -> nftmeta.v1.endpoint.Endpoint
	12, // 2: ranker.v1.endpoint.UpdateEndpointRequest.State:type_name -> chain.EndpointState
	11, // 3: ranker.v1.endpoint.UpdateEndpointResponse.Info:type_name -> nftmeta.v1.endpoint.Endpoint
	11, // 4: ranker.v1.endpoint.GetEndpointResponse.Info:type_name -> nftmeta.v1.endpoint.Endpoint
	10, // 5: ranker.v1.endpoint.GetEndpointsRequest.ChainType:type_name -> chain.ChainType
	12, // 6: ranker.v1.endpoint.GetEndpointsRequest.State:type_name -> chain.EndpointState
	11, // 7: ranker.v1.endpoint.GetEndpointsResponse.Infos:type_name -> nftmeta.v1.endpoint.Endpoint
	11, // 8: ranker.v1.endpoint.DeleteEndpointResponse.Info:type_name -> nftmeta.v1.endpoint.Endpoint
	0,  // 9: ranker.v1.endpoint.Manager.CreateEndpoint:input_type -> ranker.v1.endpoint.CreateEndpointRequest
	2,  // 10: ranker.v1.endpoint.Manager.UpdateEndpoint:input_type -> ranker.v1.endpoint.UpdateEndpointRequest
	4,  // 11: ranker.v1.endpoint.Manager.GetEndpoint:input_type -> ranker.v1.endpoint.GetEndpointRequest
	6,  // 12: ranker.v1.endpoint.Manager.GetEndpoints:input_type -> ranker.v1.endpoint.GetEndpointsRequest
	8,  // 13: ranker.v1.endpoint.Manager.DeleteEndpoint:input_type -> ranker.v1.endpoint.DeleteEndpointRequest
	1,  // 14: ranker.v1.endpoint.Manager.CreateEndpoint:output_type -> ranker.v1.endpoint.CreateEndpointResponse
	3,  // 15: ranker.v1.endpoint.Manager.UpdateEndpoint:output_type -> ranker.v1.endpoint.UpdateEndpointResponse
	5,  // 16: ranker.v1.endpoint.Manager.GetEndpoint:output_type -> ranker.v1.endpoint.GetEndpointResponse
	7,  // 17: ranker.v1.endpoint.Manager.GetEndpoints:output_type -> ranker.v1.endpoint.GetEndpointsResponse
	9,  // 18: ranker.v1.endpoint.Manager.DeleteEndpoint:output_type -> ranker.v1.endpoint.DeleteEndpointResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_endpoint_endpoint_proto_init() }
func file_web3eye_ranker_v1_endpoint_endpoint_proto_init() {
	if File_web3eye_ranker_v1_endpoint_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEndpointRequest); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEndpointResponse); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEndpointRequest); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEndpointResponse); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointRequest); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointResponse); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointsRequest); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndpointsResponse); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEndpointRequest); i {
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
		file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEndpointResponse); i {
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
	file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_ranker_v1_endpoint_endpoint_proto_goTypes,
		DependencyIndexes: file_web3eye_ranker_v1_endpoint_endpoint_proto_depIdxs,
		MessageInfos:      file_web3eye_ranker_v1_endpoint_endpoint_proto_msgTypes,
	}.Build()
	File_web3eye_ranker_v1_endpoint_endpoint_proto = out.File
	file_web3eye_ranker_v1_endpoint_endpoint_proto_rawDesc = nil
	file_web3eye_ranker_v1_endpoint_endpoint_proto_goTypes = nil
	file_web3eye_ranker_v1_endpoint_endpoint_proto_depIdxs = nil
}
