// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: web3eye/basetype/v1/chain.proto

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

type ChainType int32

const (
	ChainType_ChainUnkonwn ChainType = 0
	ChainType_Ethereum     ChainType = 10
	ChainType_Solana       ChainType = 20
)

// Enum value maps for ChainType.
var (
	ChainType_name = map[int32]string{
		0:  "ChainUnkonwn",
		10: "Ethereum",
		20: "Solana",
	}
	ChainType_value = map[string]int32{
		"ChainUnkonwn": 0,
		"Ethereum":     10,
		"Solana":       20,
	}
)

func (x ChainType) Enum() *ChainType {
	p := new(ChainType)
	*p = x
	return p
}

func (x ChainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChainType) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[0].Descriptor()
}

func (ChainType) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[0]
}

func (x ChainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChainType.Descriptor instead.
func (ChainType) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{0}
}

type TokenType int32

const (
	TokenType_TokenUnkonwn          TokenType = 0
	TokenType_Native                TokenType = 1
	TokenType_ERC20                 TokenType = 2
	TokenType_ERC721                TokenType = 10
	TokenType_ERC721_WITH_CRITERIA  TokenType = 11
	TokenType_ERC1155               TokenType = 20
	TokenType_ERC1155_WITH_CRITERIA TokenType = 21
	TokenType_Metaplex              TokenType = 30
	TokenType_NoURI                 TokenType = 40
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0:  "TokenUnkonwn",
		1:  "Native",
		2:  "ERC20",
		10: "ERC721",
		11: "ERC721_WITH_CRITERIA",
		20: "ERC1155",
		21: "ERC1155_WITH_CRITERIA",
		30: "Metaplex",
		40: "NoURI",
	}
	TokenType_value = map[string]int32{
		"TokenUnkonwn":          0,
		"Native":                1,
		"ERC20":                 2,
		"ERC721":                10,
		"ERC721_WITH_CRITERIA":  11,
		"ERC1155":               20,
		"ERC1155_WITH_CRITERIA": 21,
		"Metaplex":              30,
		"NoURI":                 40,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[1].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[1]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{1}
}

type UrlType int32

const (
	UrlType_UrlTypeUnkonwn UrlType = 0
	UrlType_ImageUrl       UrlType = 10
	UrlType_VedieoUrl      UrlType = 20
)

// Enum value maps for UrlType.
var (
	UrlType_name = map[int32]string{
		0:  "UrlTypeUnkonwn",
		10: "ImageUrl",
		20: "VedieoUrl",
	}
	UrlType_value = map[string]int32{
		"UrlTypeUnkonwn": 0,
		"ImageUrl":       10,
		"VedieoUrl":      20,
	}
)

func (x UrlType) Enum() *UrlType {
	p := new(UrlType)
	*p = x
	return p
}

func (x UrlType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UrlType) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[2].Descriptor()
}

func (UrlType) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[2]
}

func (x UrlType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UrlType.Descriptor instead.
func (UrlType) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{2}
}

type BlockParseState int32

const (
	BlockParseState_BlockTypeUnkonwn BlockParseState = 0
	BlockParseState_BlockTypeStart   BlockParseState = 10
	BlockParseState_BlockTypeFinish  BlockParseState = 20
	BlockParseState_BlockTypeFailed  BlockParseState = 30
)

// Enum value maps for BlockParseState.
var (
	BlockParseState_name = map[int32]string{
		0:  "BlockTypeUnkonwn",
		10: "BlockTypeStart",
		20: "BlockTypeFinish",
		30: "BlockTypeFailed",
	}
	BlockParseState_value = map[string]int32{
		"BlockTypeUnkonwn": 0,
		"BlockTypeStart":   10,
		"BlockTypeFinish":  20,
		"BlockTypeFailed":  30,
	}
)

func (x BlockParseState) Enum() *BlockParseState {
	p := new(BlockParseState)
	*p = x
	return p
}

func (x BlockParseState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockParseState) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[3].Descriptor()
}

func (BlockParseState) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[3]
}

func (x BlockParseState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockParseState.Descriptor instead.
func (BlockParseState) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{3}
}

type OrderItemType int32

const (
	OrderItemType_OrderItemTypeUnkonwn OrderItemType = 0
	// Target: That`s mean want to gain the goods.
	OrderItemType_OrderItemTarget OrderItemType = 10
	// Target: That`s mean willing to offer goods for gaining something.
	OrderItemType_OrderItemOffer OrderItemType = 20
)

// Enum value maps for OrderItemType.
var (
	OrderItemType_name = map[int32]string{
		0:  "OrderItemTypeUnkonwn",
		10: "OrderItemTarget",
		20: "OrderItemOffer",
	}
	OrderItemType_value = map[string]int32{
		"OrderItemTypeUnkonwn": 0,
		"OrderItemTarget":      10,
		"OrderItemOffer":       20,
	}
)

func (x OrderItemType) Enum() *OrderItemType {
	p := new(OrderItemType)
	*p = x
	return p
}

func (x OrderItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[4].Descriptor()
}

func (OrderItemType) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[4]
}

func (x OrderItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderItemType.Descriptor instead.
func (OrderItemType) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{4}
}

type SyncState int32

const (
	SyncState_Default SyncState = 0
	SyncState_Start   SyncState = 10
	SyncState_Pause   SyncState = 20
	SyncState_Finish  SyncState = 40
	SyncState_Failed  SyncState = 50
)

// Enum value maps for SyncState.
var (
	SyncState_name = map[int32]string{
		0:  "Default",
		10: "Start",
		20: "Pause",
		40: "Finish",
		50: "Failed",
	}
	SyncState_value = map[string]int32{
		"Default": 0,
		"Start":   10,
		"Pause":   20,
		"Finish":  40,
		"Failed":  50,
	}
)

func (x SyncState) Enum() *SyncState {
	p := new(SyncState)
	*p = x
	return p
}

func (x SyncState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncState) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[5].Descriptor()
}

func (SyncState) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[5]
}

func (x SyncState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SyncState.Descriptor instead.
func (SyncState) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{5}
}

type EndpointState int32

const (
	EndpointState_EndpointDefault   EndpointState = 0
	EndpointState_EndpointAvaliable EndpointState = 10
	EndpointState_EndpointUnstable  EndpointState = 20
	EndpointState_EndpointError     EndpointState = 40
)

// Enum value maps for EndpointState.
var (
	EndpointState_name = map[int32]string{
		0:  "EndpointDefault",
		10: "EndpointAvaliable",
		20: "EndpointUnstable",
		40: "EndpointError",
	}
	EndpointState_value = map[string]int32{
		"EndpointDefault":   0,
		"EndpointAvaliable": 10,
		"EndpointUnstable":  20,
		"EndpointError":     40,
	}
)

func (x EndpointState) Enum() *EndpointState {
	p := new(EndpointState)
	*p = x
	return p
}

func (x EndpointState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointState) Descriptor() protoreflect.EnumDescriptor {
	return file_web3eye_basetype_v1_chain_proto_enumTypes[6].Descriptor()
}

func (EndpointState) Type() protoreflect.EnumType {
	return &file_web3eye_basetype_v1_chain_proto_enumTypes[6]
}

func (x EndpointState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointState.Descriptor instead.
func (EndpointState) EnumDescriptor() ([]byte, []int) {
	return file_web3eye_basetype_v1_chain_proto_rawDescGZIP(), []int{6}
}

var File_web3eye_basetype_v1_chain_proto protoreflect.FileDescriptor

var file_web3eye_basetype_v1_chain_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2a, 0x37, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x55, 0x6e,
	0x6b, 0x6f, 0x6e, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x10,
	0x14, 0x2a, 0x9b, 0x01, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x6e, 0x6b, 0x6f, 0x6e, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x43, 0x32, 0x30, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x52, 0x43, 0x37,
	0x32, 0x31, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x5f, 0x57,
	0x49, 0x54, 0x48, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41, 0x10, 0x0b, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x10, 0x14, 0x12, 0x19, 0x0a, 0x15, 0x45,
	0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x52, 0x49, 0x54,
	0x45, 0x52, 0x49, 0x41, 0x10, 0x15, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x78, 0x10, 0x1e, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x55, 0x52, 0x49, 0x10, 0x28, 0x2a,
	0x3a, 0x0a, 0x07, 0x55, 0x72, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x72,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6f, 0x6e, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x65, 0x64, 0x69, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x10, 0x14, 0x2a, 0x65, 0x0a, 0x0f, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6f, 0x6e,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x10, 0x14, 0x12, 0x13, 0x0a,
	0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x1e, 0x2a, 0x52, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6f, 0x6e, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x10, 0x14, 0x2a, 0x46, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x10, 0x28, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x32, 0x2a, 0x64,
	0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x14, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x28, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65,
	0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33,
	0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web3eye_basetype_v1_chain_proto_rawDescOnce sync.Once
	file_web3eye_basetype_v1_chain_proto_rawDescData = file_web3eye_basetype_v1_chain_proto_rawDesc
)

func file_web3eye_basetype_v1_chain_proto_rawDescGZIP() []byte {
	file_web3eye_basetype_v1_chain_proto_rawDescOnce.Do(func() {
		file_web3eye_basetype_v1_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_web3eye_basetype_v1_chain_proto_rawDescData)
	})
	return file_web3eye_basetype_v1_chain_proto_rawDescData
}

var file_web3eye_basetype_v1_chain_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_web3eye_basetype_v1_chain_proto_goTypes = []interface{}{
	(ChainType)(0),       // 0: chain.ChainType
	(TokenType)(0),       // 1: chain.TokenType
	(UrlType)(0),         // 2: chain.UrlType
	(BlockParseState)(0), // 3: chain.BlockParseState
	(OrderItemType)(0),   // 4: chain.OrderItemType
	(SyncState)(0),       // 5: chain.SyncState
	(EndpointState)(0),   // 6: chain.EndpointState
}
var file_web3eye_basetype_v1_chain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web3eye_basetype_v1_chain_proto_init() }
func file_web3eye_basetype_v1_chain_proto_init() {
	if File_web3eye_basetype_v1_chain_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web3eye_basetype_v1_chain_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web3eye_basetype_v1_chain_proto_goTypes,
		DependencyIndexes: file_web3eye_basetype_v1_chain_proto_depIdxs,
		EnumInfos:         file_web3eye_basetype_v1_chain_proto_enumTypes,
	}.Build()
	File_web3eye_basetype_v1_chain_proto = out.File
	file_web3eye_basetype_v1_chain_proto_rawDesc = nil
	file_web3eye_basetype_v1_chain_proto_goTypes = nil
	file_web3eye_basetype_v1_chain_proto_depIdxs = nil
}
