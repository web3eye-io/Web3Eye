// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/transfer/transfer.proto

package transfer

import (
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	_ "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/order"
	_ "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
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

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract      string       `protobuf:"bytes,10,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenType     v1.TokenType `protobuf:"varint,20,opt,name=TokenType,proto3,enum=chain.TokenType" json:"TokenType,omitempty"`
	TokenID       string       `protobuf:"bytes,30,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Amount        uint64       `protobuf:"varint,40,opt,name=Amount,proto3" json:"Amount,omitempty"`
	AmountStr     string       `protobuf:"bytes,41,opt,name=AmountStr,proto3" json:"AmountStr,omitempty"`
	Remark        string       `protobuf:"bytes,50,opt,name=Remark,proto3" json:"Remark,omitempty"`
	Name          string       `protobuf:"bytes,60,opt,name=Name,proto3" json:"Name,omitempty"`
	Symbol        string       `protobuf:"bytes,70,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Decimals      uint32       `protobuf:"varint,80,opt,name=Decimals,proto3" json:"Decimals,omitempty"`
	OrderItemType string       `protobuf:"bytes,90,opt,name=OrderItemType,proto3" json:"OrderItemType,omitempty"`
	ImageURL      string       `protobuf:"bytes,100,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItem) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *OrderItem) GetTokenType() v1.TokenType {
	if x != nil {
		return x.TokenType
	}
	return v1.TokenType(0)
}

func (x *OrderItem) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *OrderItem) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderItem) GetAmountStr() string {
	if x != nil {
		return x.AmountStr
	}
	return ""
}

func (x *OrderItem) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *OrderItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderItem) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *OrderItem) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *OrderItem) GetOrderItemType() string {
	if x != nil {
		return x.OrderItemType
	}
	return ""
}

func (x *OrderItem) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32       `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID       string       `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty"`
	ChainType   v1.ChainType `protobuf:"varint,20,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID     string       `protobuf:"bytes,30,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract    string       `protobuf:"bytes,40,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenType   v1.TokenType `protobuf:"varint,50,opt,name=TokenType,proto3,enum=chain.TokenType" json:"TokenType,omitempty"`
	TokenID     string       `protobuf:"bytes,60,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	From        string       `protobuf:"bytes,70,opt,name=From,proto3" json:"From,omitempty"`
	To          string       `protobuf:"bytes,80,opt,name=To,proto3" json:"To,omitempty"`
	Amount      uint64       `protobuf:"varint,90,opt,name=Amount,proto3" json:"Amount,omitempty"`
	BlockNumber uint64       `protobuf:"varint,100,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	TxHash      string       `protobuf:"bytes,110,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
	BlockHash   string       `protobuf:"bytes,120,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	TxTime      uint64       `protobuf:"varint,130,opt,name=TxTime,proto3" json:"TxTime,omitempty"`
	Remark      string       `protobuf:"bytes,140,opt,name=Remark,proto3" json:"Remark,omitempty"`
	LogIndex    uint32       `protobuf:"varint,150,opt,name=LogIndex,proto3" json:"LogIndex,omitempty"`
	TargetItems []*OrderItem `protobuf:"bytes,250,rep,name=TargetItems,proto3" json:"TargetItems,omitempty"`
	OfferItems  []*OrderItem `protobuf:"bytes,160,rep,name=OfferItems,proto3" json:"OfferItems,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *Transfer) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Transfer) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Transfer) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType(0)
}

func (x *Transfer) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *Transfer) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *Transfer) GetTokenType() v1.TokenType {
	if x != nil {
		return x.TokenType
	}
	return v1.TokenType(0)
}

func (x *Transfer) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *Transfer) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transfer) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transfer) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transfer) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Transfer) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *Transfer) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transfer) GetTxTime() uint64 {
	if x != nil {
		return x.TxTime
	}
	return 0
}

func (x *Transfer) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Transfer) GetLogIndex() uint32 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *Transfer) GetTargetItems() []*OrderItem {
	if x != nil {
		return x.TargetItems
	}
	return nil
}

func (x *Transfer) GetOfferItems() []*OrderItem {
	if x != nil {
		return x.OfferItems
	}
	return nil
}

type GetTransfersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainType v1.ChainType `protobuf:"varint,10,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID   string       `protobuf:"bytes,20,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract  string       `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenID   *string      `protobuf:"bytes,40,opt,name=TokenID,proto3,oneof" json:"TokenID,omitempty"`
	Offset    int32        `protobuf:"varint,50,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32        `protobuf:"varint,60,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTransfersRequest) Reset() {
	*x = GetTransfersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransfersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransfersRequest) ProtoMessage() {}

func (x *GetTransfersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransfersRequest.ProtoReflect.Descriptor instead.
func (*GetTransfersRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransfersRequest) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType(0)
}

func (x *GetTransfersRequest) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *GetTransfersRequest) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *GetTransfersRequest) GetTokenID() string {
	if x != nil && x.TokenID != nil {
		return *x.TokenID
	}
	return ""
}

func (x *GetTransfersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTransfersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetTransfersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Transfer `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetTransfersResponse) Reset() {
	*x = GetTransfersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransfersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransfersResponse) ProtoMessage() {}

func (x *GetTransfersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransfersResponse.ProtoReflect.Descriptor instead.
func (*GetTransfersResponse) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransfersResponse) GetInfos() []*Transfer {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetTransfersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_web3eye_ranker_v1_transfer_transfer_proto protoreflect.FileDescriptor

var file_web3eye_ranker_v1_transfer_transfer_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a,
	0x2a, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x65, 0x62,
	0x33, 0x65, 0x79, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x09,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x22, 0xc5,
	0x04, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x5a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x82, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x54, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x40, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0xfa, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0xa0, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x22, 0x60, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32,
	0x6e, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_ranker_v1_transfer_transfer_proto_rawDescOnce sync.Once
	file_web3eye_ranker_v1_transfer_transfer_proto_rawDescData = file_web3eye_ranker_v1_transfer_transfer_proto_rawDesc
)

func file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP() []byte {
	file_web3eye_ranker_v1_transfer_transfer_proto_rawDescOnce.Do(func() {
		file_web3eye_ranker_v1_transfer_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_ranker_v1_transfer_transfer_proto_rawDescData)
	})
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescData
}

var file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_web3eye_ranker_v1_transfer_transfer_proto_goTypes = []interface{}{
	(*OrderItem)(nil),            // 0: ranker.v1.transfer.OrderItem
	(*Transfer)(nil),             // 1: ranker.v1.transfer.Transfer
	(*GetTransfersRequest)(nil),  // 2: ranker.v1.transfer.GetTransfersRequest
	(*GetTransfersResponse)(nil), // 3: ranker.v1.transfer.GetTransfersResponse
	(v1.TokenType)(0),            // 4: chain.TokenType
	(v1.ChainType)(0),            // 5: chain.ChainType
}
var file_web3eye_ranker_v1_transfer_transfer_proto_depIdxs = []int32{
	4, // 0: ranker.v1.transfer.OrderItem.TokenType:type_name -> chain.TokenType
	5, // 1: ranker.v1.transfer.Transfer.ChainType:type_name -> chain.ChainType
	4, // 2: ranker.v1.transfer.Transfer.TokenType:type_name -> chain.TokenType
	0, // 3: ranker.v1.transfer.Transfer.TargetItems:type_name -> ranker.v1.transfer.OrderItem
	0, // 4: ranker.v1.transfer.Transfer.OfferItems:type_name -> ranker.v1.transfer.OrderItem
	5, // 5: ranker.v1.transfer.GetTransfersRequest.ChainType:type_name -> chain.ChainType
	1, // 6: ranker.v1.transfer.GetTransfersResponse.Infos:type_name -> ranker.v1.transfer.Transfer
	2, // 7: ranker.v1.transfer.Manager.GetTransfers:input_type -> ranker.v1.transfer.GetTransfersRequest
	3, // 8: ranker.v1.transfer.Manager.GetTransfers:output_type -> ranker.v1.transfer.GetTransfersResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_transfer_transfer_proto_init() }
func file_web3eye_ranker_v1_transfer_transfer_proto_init() {
	if File_web3eye_ranker_v1_transfer_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransfersRequest); i {
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
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransfersResponse); i {
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
	file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web3eye_ranker_v1_transfer_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_ranker_v1_transfer_transfer_proto_goTypes,
		DependencyIndexes: file_web3eye_ranker_v1_transfer_transfer_proto_depIdxs,
		MessageInfos:      file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes,
	}.Build()
	File_web3eye_ranker_v1_transfer_transfer_proto = out.File
	file_web3eye_ranker_v1_transfer_transfer_proto_rawDesc = nil
	file_web3eye_ranker_v1_transfer_transfer_proto_goTypes = nil
	file_web3eye_ranker_v1_transfer_transfer_proto_depIdxs = nil
}
