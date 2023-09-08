// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/token/token.proto

package token

import (
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	token "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
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

type SiblingToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	TokenID      string `protobuf:"bytes,20,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	ImageURL     string `protobuf:"bytes,30,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	IPFSImageURL string `protobuf:"bytes,40,opt,name=IPFSImageURL,proto3" json:"IPFSImageURL,omitempty"`
}

func (x *SiblingToken) Reset() {
	*x = SiblingToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiblingToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiblingToken) ProtoMessage() {}

func (x *SiblingToken) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiblingToken.ProtoReflect.Descriptor instead.
func (*SiblingToken) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_token_token_proto_rawDescGZIP(), []int{0}
}

func (x *SiblingToken) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SiblingToken) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *SiblingToken) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *SiblingToken) GetIPFSImageURL() string {
	if x != nil {
		return x.IPFSImageURL
	}
	return ""
}

type SearchToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nftmeta.v1.token.Token Token = 10;
	ID              string             `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	ChainType       v1.ChainType       `protobuf:"varint,20,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID         string             `protobuf:"bytes,30,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract        string             `protobuf:"bytes,40,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenType       string             `protobuf:"bytes,50,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	TokenID         string             `protobuf:"bytes,60,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Owner           string             `protobuf:"bytes,70,opt,name=Owner,proto3" json:"Owner,omitempty"`
	URI             string             `protobuf:"bytes,80,opt,name=URI,proto3" json:"URI,omitempty"`
	URIType         string             `protobuf:"bytes,90,opt,name=URIType,proto3" json:"URIType,omitempty"`
	ImageURL        string             `protobuf:"bytes,100,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	VideoURL        string             `protobuf:"bytes,110,opt,name=VideoURL,proto3" json:"VideoURL,omitempty"`
	Description     string             `protobuf:"bytes,120,opt,name=Description,proto3" json:"Description,omitempty"`
	Name            string             `protobuf:"bytes,130,opt,name=Name,proto3" json:"Name,omitempty"`
	VectorState     token.ConvertState `protobuf:"varint,140,opt,name=VectorState,proto3,enum=nftmeta.v1.token.ConvertState" json:"VectorState,omitempty"`
	VectorID        int64              `protobuf:"varint,150,opt,name=VectorID,proto3" json:"VectorID,omitempty"`
	Remark          string             `protobuf:"bytes,160,opt,name=Remark,proto3" json:"Remark,omitempty"`
	IPFSImageURL    string             `protobuf:"bytes,170,opt,name=IPFSImageURL,proto3" json:"IPFSImageURL,omitempty"`
	ImageSnapshotID string             `protobuf:"bytes,180,opt,name=ImageSnapshotID,proto3" json:"ImageSnapshotID,omitempty"`
	TransfersNum    int32              `protobuf:"varint,190,opt,name=TransfersNum,proto3" json:"TransfersNum,omitempty"`
	SiblingTokens   []*SiblingToken    `protobuf:"bytes,400,rep,name=SiblingTokens,proto3" json:"SiblingTokens,omitempty"`
	SiblingsNum     uint32             `protobuf:"varint,410,opt,name=SiblingsNum,proto3" json:"SiblingsNum,omitempty"`
	Distance        float32            `protobuf:"fixed32,500,opt,name=Distance,proto3" json:"Distance,omitempty"`
}

func (x *SearchToken) Reset() {
	*x = SearchToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchToken) ProtoMessage() {}

func (x *SearchToken) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchToken.ProtoReflect.Descriptor instead.
func (*SearchToken) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_token_token_proto_rawDescGZIP(), []int{1}
}

func (x *SearchToken) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SearchToken) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType_ChainUnkonwn
}

func (x *SearchToken) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *SearchToken) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *SearchToken) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *SearchToken) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *SearchToken) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SearchToken) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *SearchToken) GetURIType() string {
	if x != nil {
		return x.URIType
	}
	return ""
}

func (x *SearchToken) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *SearchToken) GetVideoURL() string {
	if x != nil {
		return x.VideoURL
	}
	return ""
}

func (x *SearchToken) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SearchToken) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchToken) GetVectorState() token.ConvertState {
	if x != nil {
		return x.VectorState
	}
	return token.ConvertState_Default
}

func (x *SearchToken) GetVectorID() int64 {
	if x != nil {
		return x.VectorID
	}
	return 0
}

func (x *SearchToken) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SearchToken) GetIPFSImageURL() string {
	if x != nil {
		return x.IPFSImageURL
	}
	return ""
}

func (x *SearchToken) GetImageSnapshotID() string {
	if x != nil {
		return x.ImageSnapshotID
	}
	return ""
}

func (x *SearchToken) GetTransfersNum() int32 {
	if x != nil {
		return x.TransfersNum
	}
	return 0
}

func (x *SearchToken) GetSiblingTokens() []*SiblingToken {
	if x != nil {
		return x.SiblingTokens
	}
	return nil
}

func (x *SearchToken) GetSiblingsNum() uint32 {
	if x != nil {
		return x.SiblingsNum
	}
	return 0
}

func (x *SearchToken) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type SearchTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector []float32 `protobuf:"fixed32,10,rep,packed,name=Vector,proto3" json:"Vector,omitempty"`
	Limit  uint32    `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *SearchTokenRequest) Reset() {
	*x = SearchTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTokenRequest) ProtoMessage() {}

func (x *SearchTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTokenRequest.ProtoReflect.Descriptor instead.
func (*SearchTokenRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_token_token_proto_rawDescGZIP(), []int{2}
}

func (x *SearchTokenRequest) GetVector() []float32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

func (x *SearchTokenRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageKey string `protobuf:"bytes,10,opt,name=StorageKey,proto3" json:"StorageKey,omitempty"`
	Page       uint32 `protobuf:"varint,20,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (x *SearchPageRequest) Reset() {
	*x = SearchPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPageRequest) ProtoMessage() {}

func (x *SearchPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPageRequest.ProtoReflect.Descriptor instead.
func (*SearchPageRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_token_token_proto_rawDescGZIP(), []int{3}
}

func (x *SearchPageRequest) GetStorageKey() string {
	if x != nil {
		return x.StorageKey
	}
	return ""
}

func (x *SearchPageRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos       []*SearchToken `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	StorageKey  string         `protobuf:"bytes,20,opt,name=StorageKey,proto3" json:"StorageKey,omitempty"`
	Page        uint32         `protobuf:"varint,30,opt,name=Page,proto3" json:"Page,omitempty"`
	TotalPages  uint32         `protobuf:"varint,40,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	TotalTokens uint32         `protobuf:"varint,50,opt,name=TotalTokens,proto3" json:"TotalTokens,omitempty"`
	Limit       uint32         `protobuf:"varint,60,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_token_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_token_token_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResponse) GetInfos() []*SearchToken {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *SearchResponse) GetStorageKey() string {
	if x != nil {
		return x.StorageKey
	}
	return ""
}

func (x *SearchResponse) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchResponse) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *SearchResponse) GetTotalTokens() uint32 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

func (x *SearchResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_web3eye_ranker_v1_token_token_proto protoreflect.FileDescriptor

var file_web3eye_ranker_v1_token_token_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x24, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f,
	0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x77, 0x65,
	0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a,
	0x0c, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x22, 0xe0, 0x05, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x52, 0x49, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12, 0x18, 0x0a,
	0x07, 0x55, 0x52, 0x49, 0x54, 0x79, 0x70, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x55, 0x52, 0x49, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x52, 0x4c, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x52, 0x4c, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x78,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x13, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6e,
	0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x23, 0x0a, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0xaa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x52, 0x4c, 0x12, 0x29, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x44, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x44, 0x12,
	0x23, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x4e, 0x75, 0x6d, 0x18,
	0xbe, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x4e, 0x75, 0x6d, 0x12, 0x44, 0x0a, 0x0d, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x90, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53,
	0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0d, 0x53, 0x69, 0x62,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x53, 0x69,
	0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x9a, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x06, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x47,
	0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61, 0x6e, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0x9c, 0x04, 0x0a, 0x07, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x21, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x25, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x6e,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x6e, 0x66, 0x74, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x66, 0x74, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x23, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2d,
	0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_ranker_v1_token_token_proto_rawDescOnce sync.Once
	file_web3eye_ranker_v1_token_token_proto_rawDescData = file_web3eye_ranker_v1_token_token_proto_rawDesc
)

func file_web3eye_ranker_v1_token_token_proto_rawDescGZIP() []byte {
	file_web3eye_ranker_v1_token_token_proto_rawDescOnce.Do(func() {
		file_web3eye_ranker_v1_token_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_ranker_v1_token_token_proto_rawDescData)
	})
	return file_web3eye_ranker_v1_token_token_proto_rawDescData
}

var file_web3eye_ranker_v1_token_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_web3eye_ranker_v1_token_token_proto_goTypes = []interface{}{
	(*SiblingToken)(nil),               // 0: ranker.v1.token.SiblingToken
	(*SearchToken)(nil),                // 1: ranker.v1.token.SearchToken
	(*SearchTokenRequest)(nil),         // 2: ranker.v1.token.SearchTokenRequest
	(*SearchPageRequest)(nil),          // 3: ranker.v1.token.SearchPageRequest
	(*SearchResponse)(nil),             // 4: ranker.v1.token.SearchResponse
	(v1.ChainType)(0),                  // 5: chain.ChainType
	(token.ConvertState)(0),            // 6: nftmeta.v1.token.ConvertState
	(*token.GetTokenRequest)(nil),      // 7: nftmeta.v1.token.GetTokenRequest
	(*token.GetTokenOnlyRequest)(nil),  // 8: nftmeta.v1.token.GetTokenOnlyRequest
	(*token.GetTokensRequest)(nil),     // 9: nftmeta.v1.token.GetTokensRequest
	(*token.CountTokensRequest)(nil),   // 10: nftmeta.v1.token.CountTokensRequest
	(*token.GetTokenResponse)(nil),     // 11: nftmeta.v1.token.GetTokenResponse
	(*token.GetTokenOnlyResponse)(nil), // 12: nftmeta.v1.token.GetTokenOnlyResponse
	(*token.GetTokensResponse)(nil),    // 13: nftmeta.v1.token.GetTokensResponse
	(*token.CountTokensResponse)(nil),  // 14: nftmeta.v1.token.CountTokensResponse
}
var file_web3eye_ranker_v1_token_token_proto_depIdxs = []int32{
	5,  // 0: ranker.v1.token.SearchToken.ChainType:type_name -> chain.ChainType
	6,  // 1: ranker.v1.token.SearchToken.VectorState:type_name -> nftmeta.v1.token.ConvertState
	0,  // 2: ranker.v1.token.SearchToken.SiblingTokens:type_name -> ranker.v1.token.SiblingToken
	1,  // 3: ranker.v1.token.SearchResponse.Infos:type_name -> ranker.v1.token.SearchToken
	7,  // 4: ranker.v1.token.Manager.GetToken:input_type -> nftmeta.v1.token.GetTokenRequest
	8,  // 5: ranker.v1.token.Manager.GetTokenOnly:input_type -> nftmeta.v1.token.GetTokenOnlyRequest
	9,  // 6: ranker.v1.token.Manager.GetTokens:input_type -> nftmeta.v1.token.GetTokensRequest
	10, // 7: ranker.v1.token.Manager.CountTokens:input_type -> nftmeta.v1.token.CountTokensRequest
	2,  // 8: ranker.v1.token.Manager.Search:input_type -> ranker.v1.token.SearchTokenRequest
	3,  // 9: ranker.v1.token.Manager.SearchPage:input_type -> ranker.v1.token.SearchPageRequest
	11, // 10: ranker.v1.token.Manager.GetToken:output_type -> nftmeta.v1.token.GetTokenResponse
	12, // 11: ranker.v1.token.Manager.GetTokenOnly:output_type -> nftmeta.v1.token.GetTokenOnlyResponse
	13, // 12: ranker.v1.token.Manager.GetTokens:output_type -> nftmeta.v1.token.GetTokensResponse
	14, // 13: ranker.v1.token.Manager.CountTokens:output_type -> nftmeta.v1.token.CountTokensResponse
	4,  // 14: ranker.v1.token.Manager.Search:output_type -> ranker.v1.token.SearchResponse
	4,  // 15: ranker.v1.token.Manager.SearchPage:output_type -> ranker.v1.token.SearchResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_token_token_proto_init() }
func file_web3eye_ranker_v1_token_token_proto_init() {
	if File_web3eye_ranker_v1_token_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_ranker_v1_token_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiblingToken); i {
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
		file_web3eye_ranker_v1_token_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchToken); i {
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
		file_web3eye_ranker_v1_token_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTokenRequest); i {
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
		file_web3eye_ranker_v1_token_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPageRequest); i {
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
		file_web3eye_ranker_v1_token_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
			RawDescriptor: file_web3eye_ranker_v1_token_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_ranker_v1_token_token_proto_goTypes,
		DependencyIndexes: file_web3eye_ranker_v1_token_token_proto_depIdxs,
		MessageInfos:      file_web3eye_ranker_v1_token_token_proto_msgTypes,
	}.Build()
	File_web3eye_ranker_v1_token_token_proto = out.File
	file_web3eye_ranker_v1_token_token_proto_rawDesc = nil
	file_web3eye_ranker_v1_token_token_proto_goTypes = nil
	file_web3eye_ranker_v1_token_token_proto_depIdxs = nil
}
