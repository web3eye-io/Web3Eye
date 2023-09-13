// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/transfer/transfer.proto

package transfer

import (
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	transfer "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/transfer"
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

type GetTransfersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainType v1.ChainType `protobuf:"varint,10,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID   string       `protobuf:"bytes,20,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract  string       `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenID   string       `protobuf:"bytes,40,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Offset    int32        `protobuf:"varint,50,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32        `protobuf:"varint,60,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTransfersRequest) Reset() {
	*x = GetTransfersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransfersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransfersRequest) ProtoMessage() {}

func (x *GetTransfersRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTransfersRequest.ProtoReflect.Descriptor instead.
func (*GetTransfersRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *GetTransfersRequest) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType_ChainUnkonwn
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
	if x != nil {
		return x.TokenID
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

type CountTransfersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainType v1.ChainType `protobuf:"varint,10,opt,name=ChainType,proto3,enum=chain.ChainType" json:"ChainType,omitempty"`
	ChainID   string       `protobuf:"bytes,20,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Contract  string       `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
	TokenID   string       `protobuf:"bytes,40,opt,name=TokenID,proto3" json:"TokenID,omitempty"`
	Offset    int32        `protobuf:"varint,50,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32        `protobuf:"varint,60,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *CountTransfersRequest) Reset() {
	*x = CountTransfersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountTransfersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountTransfersRequest) ProtoMessage() {}

func (x *CountTransfersRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CountTransfersRequest.ProtoReflect.Descriptor instead.
func (*CountTransfersRequest) Descriptor() ([]byte, []int) {
	return file_web3eye_ranker_v1_transfer_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *CountTransfersRequest) GetChainType() v1.ChainType {
	if x != nil {
		return x.ChainType
	}
	return v1.ChainType_ChainUnkonwn
}

func (x *CountTransfersRequest) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

func (x *CountTransfersRequest) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *CountTransfersRequest) GetTokenID() string {
	if x != nil {
		return x.TokenID
	}
	return ""
}

func (x *CountTransfersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CountTransfersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
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
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x77, 0x65, 0x62,
	0x33, 0x65, 0x79, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0xc5, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0xdb, 0x01, 0x0a, 0x07, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x29,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x66, 0x74, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69,
	0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_web3eye_ranker_v1_transfer_transfer_proto_goTypes = []interface{}{
	(*GetTransfersRequest)(nil),             // 0: ranker.v1.transfer.GetTransfersRequest
	(*CountTransfersRequest)(nil),           // 1: ranker.v1.transfer.CountTransfersRequest
	(v1.ChainType)(0),                       // 2: chain.ChainType
	(*transfer.GetTransfersResponse)(nil),   // 3: nftmeta.v1.transfer.GetTransfersResponse
	(*transfer.CountTransfersResponse)(nil), // 4: nftmeta.v1.transfer.CountTransfersResponse
}
var file_web3eye_ranker_v1_transfer_transfer_proto_depIdxs = []int32{
	2, // 0: ranker.v1.transfer.GetTransfersRequest.ChainType:type_name -> chain.ChainType
	2, // 1: ranker.v1.transfer.CountTransfersRequest.ChainType:type_name -> chain.ChainType
	0, // 2: ranker.v1.transfer.Manager.GetTransfers:input_type -> ranker.v1.transfer.GetTransfersRequest
	1, // 3: ranker.v1.transfer.Manager.CountTransfers:input_type -> ranker.v1.transfer.CountTransfersRequest
	3, // 4: ranker.v1.transfer.Manager.GetTransfers:output_type -> nftmeta.v1.transfer.GetTransfersResponse
	4, // 5: ranker.v1.transfer.Manager.CountTransfers:output_type -> nftmeta.v1.transfer.CountTransfersResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_transfer_transfer_proto_init() }
func file_web3eye_ranker_v1_transfer_transfer_proto_init() {
	if File_web3eye_ranker_v1_transfer_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_web3eye_ranker_v1_transfer_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountTransfersRequest); i {
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
			RawDescriptor: file_web3eye_ranker_v1_transfer_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
