// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: web3eye/entrance/v1/token/token.proto

package token

import (
	token "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	token1 "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_web3eye_entrance_v1_token_token_proto protoreflect.FileDescriptor

var file_web3eye_entrance_v1_token_token_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79,
	0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x6b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x0a,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79,
	0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_web3eye_entrance_v1_token_token_proto_goTypes = []interface{}{
	(*token.GetTokenRequest)(nil),    // 0: nftmeta.v1.token.GetTokenRequest
	(*token1.SearchPageRequest)(nil), // 1: ranker.v1.token.SearchPageRequest
	(*token.GetTokenResponse)(nil),   // 2: nftmeta.v1.token.GetTokenResponse
	(*token1.SearchResponse)(nil),    // 3: ranker.v1.token.SearchResponse
}
var file_web3eye_entrance_v1_token_token_proto_depIdxs = []int32{
	0, // 0: entrance.v1.token.Manager.GetToken:input_type -> nftmeta.v1.token.GetTokenRequest
	1, // 1: entrance.v1.token.Manager.SearchPage:input_type -> ranker.v1.token.SearchPageRequest
	2, // 2: entrance.v1.token.Manager.GetToken:output_type -> nftmeta.v1.token.GetTokenResponse
	3, // 3: entrance.v1.token.Manager.SearchPage:output_type -> ranker.v1.token.SearchResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web3eye_entrance_v1_token_token_proto_init() }
func file_web3eye_entrance_v1_token_token_proto_init() {
	if File_web3eye_entrance_v1_token_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web3eye_entrance_v1_token_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_entrance_v1_token_token_proto_goTypes,
		DependencyIndexes: file_web3eye_entrance_v1_token_token_proto_depIdxs,
	}.Build()
	File_web3eye_entrance_v1_token_token_proto = out.File
	file_web3eye_entrance_v1_token_token_proto_rawDesc = nil
	file_web3eye_entrance_v1_token_token_proto_goTypes = nil
	file_web3eye_entrance_v1_token_token_proto_depIdxs = nil
}