// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/dealer/v1/dealer.proto

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

type BackupState int32

const (
	BackupState_DefaultBackupState  BackupState = 0
	BackupState_BackupStateNone     BackupState = 10
	BackupState_BackupStateCreated  BackupState = 20
	BackupState_BackupStateProposed BackupState = 30
	BackupState_BackupStateAccepted BackupState = 40
	BackupState_BackupStateSuccess  BackupState = 50
	BackupState_BackupStateFail     BackupState = 60
)

// Enum value maps for BackupState.
var (
	BackupState_name = map[int32]string{
		0:  "DefaultBackupState",
		10: "BackupStateNone",
		20: "BackupStateCreated",
		30: "BackupStateProposed",
		40: "BackupStateAccepted",
		50: "BackupStateSuccess",
		60: "BackupStateFail",
	}
	BackupState_value = map[string]int32{
		"DefaultBackupState":  0,
		"BackupStateNone":     10,
		"BackupStateCreated":  20,
		"BackupStateProposed": 30,
		"BackupStateAccepted": 40,
		"BackupStateSuccess":  50,
		"BackupStateFail":     60,
	}
)

func (x BackupState) Enum() *BackupState {
	p := new(BackupState)
	*p = x
	return p
}

func (x BackupState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupState) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_dealer_v1_dealer_proto_enumTypes[0].Descriptor()
}

func (BackupState) Type() protoreflect.EnumType {
	return &file_web3eye_dealer_v1_dealer_proto_enumTypes[0]
}

func (x BackupState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupState.Descriptor instead.
func (BackupState) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{0}
}

type ContentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	URI       string `protobuf:"bytes,20,opt,name=URI,proto3" json:"URI,omitempty"`
	ChainType string `protobuf:"bytes,30,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	ChainID   string `protobuf:"bytes,40,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract  string `protobuf:"bytes,50,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenID   string `protobuf:"bytes,60,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	FileName  string `protobuf:"bytes,70,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *ContentItem) Reset() {
	*x = ContentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentItem) ProtoMessage() {}

func (x *ContentItem) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentItem.ProtoReflect.Descriptor instead.
func (*ContentItem) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{0}
}

func (x *ContentItem) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ContentItem) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *ContentItem) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *ContentItem) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *ContentItem) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *ContentItem) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *ContentItem) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            uint32         `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Index         uint64         `protobuf:"varint,20,opt,name=Index,proto3" json:"Index,omitempty"`
	SnapshotCommP string         `protobuf:"bytes,30,opt,name=SnapshotCommP,proto3" json:"SnapshotCommP,omitempty"`
	SnapshotRoot  string         `protobuf:"bytes,40,opt,name=SnapshotRoot,proto3" json:"SnapshotRoot,omitempty"`
	SnapshotURI   string         `protobuf:"bytes,50,opt,name=SnapshotURI,proto3" json:"SnapshotURI,omitempty"`
	Items         []*ContentItem `protobuf:"bytes,60,rep,name=Items,proto3" json:"Items,omitempty"`
	BackupState   BackupState    `protobuf:"varint,70,opt,name=BackupState,proto3,enum=dealer.v1.BackupState" json:"BackupState,omitempty"`
	ProposalCID   string         `protobuf:"bytes,80,opt,name=ProposalCID,proto3" json:"ProposalCID,omitempty"`
	DealID        uint64         `protobuf:"varint,90,opt,name=DealID,proto3" json:"DealID,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{1}
}

func (x *Snapshot) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Snapshot) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Snapshot) GetSnapshotCommP() string {
	if x != nil {
		return x.SnapshotCommP
	}
	return ""
}

func (x *Snapshot) GetSnapshotRoot() string {
	if x != nil {
		return x.SnapshotRoot
	}
	return ""
}

func (x *Snapshot) GetSnapshotURI() string {
	if x != nil {
		return x.SnapshotURI
	}
	return ""
}

func (x *Snapshot) GetItems() []*ContentItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Snapshot) GetBackupState() BackupState {
	if x != nil {
		return x.BackupState
	}
	return BackupState_DefaultBackupState
}

func (x *Snapshot) GetProposalCID() string {
	if x != nil {
		return x.ProposalCID
	}
	return ""
}

func (x *Snapshot) GetDealID() uint64 {
	if x != nil {
		return x.DealID
	}
	return 0
}

type CreateSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapshotCommP string         `protobuf:"bytes,10,opt,name=SnapshotCommP,proto3" json:"SnapshotCommP,omitempty"`
	SnapshotRoot  string         `protobuf:"bytes,20,opt,name=SnapshotRoot,proto3" json:"SnapshotRoot,omitempty"`
	SnapshotURI   string         `protobuf:"bytes,30,opt,name=SnapshotURI,proto3" json:"SnapshotURI,omitempty"`
	Items         []*ContentItem `protobuf:"bytes,40,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *CreateSnapshotRequest) Reset() {
	*x = CreateSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnapshotRequest) ProtoMessage() {}

func (x *CreateSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnapshotRequest.ProtoReflect.Descriptor instead.
func (*CreateSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSnapshotRequest) GetSnapshotCommP() string {
	if x != nil {
		return x.SnapshotCommP
	}
	return ""
}

func (x *CreateSnapshotRequest) GetSnapshotRoot() string {
	if x != nil {
		return x.SnapshotRoot
	}
	return ""
}

func (x *CreateSnapshotRequest) GetSnapshotURI() string {
	if x != nil {
		return x.SnapshotURI
	}
	return ""
}

func (x *CreateSnapshotRequest) GetItems() []*ContentItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Snapshot `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateSnapshotResponse) Reset() {
	*x = CreateSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnapshotResponse) ProtoMessage() {}

func (x *CreateSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnapshotResponse.ProtoReflect.Descriptor instead.
func (*CreateSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSnapshotResponse) GetInfo() *Snapshot {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetSnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexes []uint64 `protobuf:"varint,10,rep,packed,name=Indexes,proto3" json:"Indexes,omitempty"`
}

func (x *GetSnapshotsRequest) Reset() {
	*x = GetSnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotsRequest) ProtoMessage() {}

func (x *GetSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*GetSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{4}
}

func (x *GetSnapshotsRequest) GetIndexes() []uint64 {
	if x != nil {
		return x.Indexes
	}
	return nil
}

type GetSnapshotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Snapshot `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint64      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetSnapshotsResponse) Reset() {
	*x = GetSnapshotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotsResponse) ProtoMessage() {}

func (x *GetSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*GetSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{5}
}

func (x *GetSnapshotsResponse) GetInfos() []*Snapshot {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetSnapshotsResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateBackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64 `protobuf:"varint,10,opt,name=Index,proto3" json:"Index,omitempty"`
}

func (x *CreateBackupRequest) Reset() {
	*x = CreateBackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBackupRequest) ProtoMessage() {}

func (x *CreateBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBackupRequest.ProtoReflect.Descriptor instead.
func (*CreateBackupRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBackupRequest) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type CreateBackupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Snapshot `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateBackupResponse) Reset() {
	*x = CreateBackupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBackupResponse) ProtoMessage() {}

func (x *CreateBackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_dealer_v1_dealer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBackupResponse.ProtoReflect.Descriptor instead.
func (*CreateBackupResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_dealer_v1_dealer_proto_rawDescGZIP(), []int{7}
}

func (x *CreateBackupResponse) GetInfo() *Snapshot {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_web3eye_dealer_v1_dealer_proto protoreflect.FileDescriptor

var file_web3eye_dealer_v1_dealer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xb9, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x52, 0x49, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x50, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x50,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x6f, 0x6f, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x55, 0x52, 0x49, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x55, 0x52, 0x49, 0x12, 0x2c, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x3c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x43, 0x49, 0x44, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x43, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x44, 0x65, 0x61, 0x6c, 0x49, 0x44, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x50, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x50, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x55, 0x52, 0x49, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x55, 0x52, 0x49, 0x12, 0x2c,
	0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x41, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x22, 0x57, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0xb1, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x14, 0x12, 0x17, 0x0a, 0x13,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x10, 0x1e, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x28, 0x12, 0x16,
	0x0a, 0x12, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x3c, 0x32, 0x88, 0x02, 0x0a, 0x07,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x12, 0x1e, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f, 0x2f,
	0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65,
	0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_dealer_v1_dealer_proto_rawDescOnce sync.Once
	file_web3eye_dealer_v1_dealer_proto_rawDescData = file_web3eye_dealer_v1_dealer_proto_rawDesc
)

func file_web3eye_dealer_v1_dealer_proto_rawDescGZIP() []byte {
	file_web3eye_dealer_v1_dealer_proto_rawDescOnce.Do(func() {
		file_web3eye_dealer_v1_dealer_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_dealer_v1_dealer_proto_rawDescData)
	})
	return file_web3eye_dealer_v1_dealer_proto_rawDescData
}

var file_web3eye_dealer_v1_dealer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_web3eye_dealer_v1_dealer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_web3eye_dealer_v1_dealer_proto_goTypes = []interface{}{
	(BackupState)(0),               // 0: dealer.v1.BackupState
	(*ContentItem)(nil),            // 1: dealer.v1.ContentItem
	(*Snapshot)(nil),               // 2: dealer.v1.Snapshot
	(*CreateSnapshotRequest)(nil),  // 3: dealer.v1.CreateSnapshotRequest
	(*CreateSnapshotResponse)(nil), // 4: dealer.v1.CreateSnapshotResponse
	(*GetSnapshotsRequest)(nil),    // 5: dealer.v1.GetSnapshotsRequest
	(*GetSnapshotsResponse)(nil),   // 6: dealer.v1.GetSnapshotsResponse
	(*CreateBackupRequest)(nil),    // 7: dealer.v1.CreateBackupRequest
	(*CreateBackupResponse)(nil),   // 8: dealer.v1.CreateBackupResponse
}
var file_web3eye_dealer_v1_dealer_proto_depIdxs = []int32{
	1, // 0: dealer.v1.Snapshot.Items:type_name -> dealer.v1.ContentItem
	0, // 1: dealer.v1.Snapshot.BackupState:type_name -> dealer.v1.BackupState
	1, // 2: dealer.v1.CreateSnapshotRequest.Items:type_name -> dealer.v1.ContentItem
	2, // 3: dealer.v1.CreateSnapshotResponse.Info:type_name -> dealer.v1.Snapshot
	2, // 4: dealer.v1.GetSnapshotsResponse.Infos:type_name -> dealer.v1.Snapshot
	2, // 5: dealer.v1.CreateBackupResponse.Info:type_name -> dealer.v1.Snapshot
	3, // 6: dealer.v1.Manager.CreateSnapshot:input_type -> dealer.v1.CreateSnapshotRequest
	5, // 7: dealer.v1.Manager.GetSnapshots:input_type -> dealer.v1.GetSnapshotsRequest
	7, // 8: dealer.v1.Manager.CreateBackup:input_type -> dealer.v1.CreateBackupRequest
	4, // 9: dealer.v1.Manager.CreateSnapshot:output_type -> dealer.v1.CreateSnapshotResponse
	6, // 10: dealer.v1.Manager.GetSnapshots:output_type -> dealer.v1.GetSnapshotsResponse
	8, // 11: dealer.v1.Manager.CreateBackup:output_type -> dealer.v1.CreateBackupResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_web3eye_dealer_v1_dealer_proto_init() }
func file_web3eye_dealer_v1_dealer_proto_init() {
	if File_web3eye_dealer_v1_dealer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_dealer_v1_dealer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentItem); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnapshotRequest); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnapshotResponse); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotsRequest); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotsResponse); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBackupRequest); i {
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
		file_web3eye_dealer_v1_dealer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBackupResponse); i {
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
			RawDescriptor: file_web3eye_dealer_v1_dealer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_dealer_v1_dealer_proto_goTypes,
		DependencyIndexes: file_web3eye_dealer_v1_dealer_proto_depIdxs,
		EnumInfos:         file_web3eye_dealer_v1_dealer_proto_enumTypes,
		MessageInfos:      file_web3eye_dealer_v1_dealer_proto_msgTypes,
	}.Build()
	File_web3eye_dealer_v1_dealer_proto = out.File
	file_web3eye_dealer_v1_dealer_proto_rawDesc = nil
	file_web3eye_dealer_v1_dealer_proto_goTypes = nil
	file_web3eye_dealer_v1_dealer_proto_depIdxs = nil
}
