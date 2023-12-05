// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/contract/contract.proto

package contract

import (
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	contract "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
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

type GetContractAndTokensReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract string `protobuf:"bytes,10,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Offset   uint32 `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit    uint32 `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetContractAndTokensReq) Reset() {
	*x = GetContractAndTokensReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractAndTokensReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractAndTokensReq) ProtoMessage() {}

func (x *GetContractAndTokensReq) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractAndTokensReq.ProtoReflect.Descriptor instead.
func (*GetContractAndTokensReq) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_contract_contract_proto_rawDescGZIP(), []int{0}
}

func (x *GetContractAndTokensReq) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *GetContractAndTokensReq) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetContractAndTokensReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ShotToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint32       `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	ChainType       v1.ChainType `protobuf:"varint,20,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	TokenType       v1.TokenType `protobuf:"varint,30,opt,name=TokenType,proto3,enum=chain.TokenType" json:"TokenType,omitempty"`
	TokenID         string       `protobuf:"bytes,40,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Owner           string       `protobuf:"bytes,50,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ImageURL        string       `protobuf:"bytes,60,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Name            string       `protobuf:"bytes,70,opt,name=Name,proto3" json:"Name,omitempty"`
	IPFSImageURL    string       `protobuf:"bytes,80,opt,name=IPFSImageURL,proto3" json:"IPFSImageURL,omitempty"`
	ImageSnapshotID string       `protobuf:"bytes,90,opt,name=ImageSnapshotID,proto3" json:"ImageSnapshotID,omitempty"`
	TransfersNum    uint32       `protobuf:"varint,100,opt,name=TransfersNum,proto3" json:"TransfersNum,omitempty"`
}

func (x *ShotToken) Reset() {
	*x = ShotToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShotToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShotToken) ProtoMessage() {}

func (x *ShotToken) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShotToken.ProtoReflect.Descriptor instead.
func (*ShotToken) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_contract_contract_proto_rawDescGZIP(), []int{1}
}

func (x *ShotToken) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShotToken) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType(0)
}

func (x *ShotToken) GetTokenType() v1.TokenType {
	if x != nil {
		return x.TokenType
	}
	return v1.TokenType(0)
}

func (x *ShotToken) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *ShotToken) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ShotToken) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *ShotToken) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShotToken) GetIPFSImageURL() string {
	if x != nil {
		return x.IPFSImageURL
	}
	return ""
}

func (x *ShotToken) GetImageSnapshotID() string {
	if x != nil {
		return x.ImageSnapshotID
	}
	return ""
}

func (x *ShotToken) GetTransfersNum() uint32 {
	if x != nil {
		return x.TransfersNum
	}
	return 0
}

type GetContractAndTokensResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract    *contract.Contract `protobuf:"bytes,10,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Tokens      []*ShotToken       `protobuf:"bytes,20,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
	TotalTokens uint32             `protobuf:"varint,30,opt,name=TotalTokens,proto3" json:"TotalTokens,omitempty"`
}

func (x *GetContractAndTokensResp) Reset() {
	*x = GetContractAndTokensResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractAndTokensResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractAndTokensResp) ProtoMessage() {}

func (x *GetContractAndTokensResp) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_contract_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractAndTokensResp.ProtoReflect.Descriptor instead.
func (*GetContractAndTokensResp) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_contract_contract_proto_rawDescGZIP(), []int{2}
}

func (x *GetContractAndTokensResp) GetContract() *contract.Contract {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *GetContractAndTokensResp) GetTokens() []*ShotToken {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *GetContractAndTokensResp) GetTotalTokens() uint32 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

var File_web3eye_ranker_v1_contract_contract_proto protoreflect.FileDescriptor

var file_web3eye_ranker_v1_contract_contract_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a,
	0x2a, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x77, 0x65, 0x62,
	0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0xcd, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2e, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x50, 0x46, 0x53, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x4e, 0x75,
	0x6d, 0x22, 0xae, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61, 0x6e, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53,
	0x68, 0x6f, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x32, 0xa6, 0x04, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x62,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x27, 0x2e,
	0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2b, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x12, 0x28, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e,
	0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2b,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41,
	0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79,
	0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_ranker_v1_contract_contract_proto_rawDescOnce sync.Once
	file_web3eye_ranker_v1_contract_contract_proto_rawDescData = file_web3eye_ranker_v1_contract_contract_proto_rawDesc
)

func file_web3eye_ranker_v1_contract_contract_proto_rawDescGZIP() []byte {
	file_web3eye_ranker_v1_contract_contract_proto_rawDescOnce.Do(func() {
		file_web3eye_ranker_v1_contract_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_ranker_v1_contract_contract_proto_rawDescData)
	})
	return file_web3eye_ranker_v1_contract_contract_proto_rawDescData
}

var file_web3eye_ranker_v1_contract_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web3eye_ranker_v1_contract_contract_proto_goTypes = []interface{}{
	(*GetContractAndTokensReq)(nil),          // 0: ranker.v1.contract.GetContractAndTokensReq
	(*ShotToken)(nil),                        // 1: ranker.v1.contract.ShotToken
	(*GetContractAndTokensResp)(nil),         // 2: ranker.v1.contract.GetContractAndTokensResp
	(v1.ChainType)(0),                        // 3: chain.ChainType
	(v1.TokenType)(0),                        // 4: chain.TokenType
	(*contract.Contract)(nil),                // 5: nftmeta.v1.contract.Contract
	(*contract.GetContractRequest)(nil),      // 6: nftmeta.v1.contract.GetContractRequest
	(*contract.GetContractOnlyRequest)(nil),  // 7: nftmeta.v1.contract.GetContractOnlyRequest
	(*contract.GetContractsRequest)(nil),     // 8: nftmeta.v1.contract.GetContractsRequest
	(*contract.CountContractsRequest)(nil),   // 9: nftmeta.v1.contract.CountContractsRequest
	(*contract.GetContractResponse)(nil),     // 10: nftmeta.v1.contract.GetContractResponse
	(*contract.GetContractOnlyResponse)(nil), // 11: nftmeta.v1.contract.GetContractOnlyResponse
	(*contract.GetContractsResponse)(nil),    // 12: nftmeta.v1.contract.GetContractsResponse
	(*contract.CountContractsResponse)(nil),  // 13: nftmeta.v1.contract.CountContractsResponse
}
var file_web3eye_ranker_v1_contract_contract_proto_depIdxs = []int32{
	3,  // 0: ranker.v1.contract.ShotToken.ChainType:type_name -> chain.ChainType
	4,  // 1: ranker.v1.contract.ShotToken.TokenType:type_name -> chain.TokenType
	5,  // 2: ranker.v1.contract.GetContractAndTokensResp.Contract:type_name -> nftmeta.v1.contract.Contract
	1,  // 3: ranker.v1.contract.GetContractAndTokensResp.Tokens:type_name -> ranker.v1.contract.ShotToken
	6,  // 4: ranker.v1.contract.Manager.GetContract:input_type -> nftmeta.v1.contract.GetContractRequest
	7,  // 5: ranker.v1.contract.Manager.GetContractOnly:input_type -> nftmeta.v1.contract.GetContractOnlyRequest
	8,  // 6: ranker.v1.contract.Manager.GetContracts:input_type -> nftmeta.v1.contract.GetContractsRequest
	9,  // 7: ranker.v1.contract.Manager.CountContracts:input_type -> nftmeta.v1.contract.CountContractsRequest
	0,  // 8: ranker.v1.contract.Manager.GetContractAndTokens:input_type -> ranker.v1.contract.GetContractAndTokensReq
	10, // 9: ranker.v1.contract.Manager.GetContract:output_type -> nftmeta.v1.contract.GetContractResponse
	11, // 10: ranker.v1.contract.Manager.GetContractOnly:output_type -> nftmeta.v1.contract.GetContractOnlyResponse
	12, // 11: ranker.v1.contract.Manager.GetContracts:output_type -> nftmeta.v1.contract.GetContractsResponse
	13, // 12: ranker.v1.contract.Manager.CountContracts:output_type -> nftmeta.v1.contract.CountContractsResponse
	2,  // 13: ranker.v1.contract.Manager.GetContractAndTokens:output_type -> ranker.v1.contract.GetContractAndTokensResp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_contract_contract_proto_init() }
func file_web3eye_ranker_v1_contract_contract_proto_init() {
	if File_web3eye_ranker_v1_contract_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_ranker_v1_contract_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractAndTokensReq); i {
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
		file_web3eye_ranker_v1_contract_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShotToken); i {
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
		file_web3eye_ranker_v1_contract_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractAndTokensResp); i {
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
			RawDescriptor: file_web3eye_ranker_v1_contract_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_ranker_v1_contract_contract_proto_goTypes,
		DependencyIndexes: file_web3eye_ranker_v1_contract_contract_proto_depIdxs,
		MessageInfos:      file_web3eye_ranker_v1_contract_contract_proto_msgTypes,
	}.Build()
	File_web3eye_ranker_v1_contract_contract_proto = out.File
	file_web3eye_ranker_v1_contract_contract_proto_rawDesc = nil
	file_web3eye_ranker_v1_contract_contract_proto_goTypes = nil
	file_web3eye_ranker_v1_contract_contract_proto_depIdxs = nil
}
