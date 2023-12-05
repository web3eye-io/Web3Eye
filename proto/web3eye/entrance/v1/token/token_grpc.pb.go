// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: web3eye/entrance/v1/token/token.proto

package token

import (
	context "context"
	token "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	token1 "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Manager_GetToken_FullMethodName   = "/entrance.v1.token.Manager/GetToken"
	Manager_SearchPage_FullMethodName = "/entrance.v1.token.Manager/SearchPage"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	GetToken(ctx context.Context, in *token.GetTokenRequest, opts ...grpc.CallOption) (*token.GetTokenResponse, error)
	SearchPage(ctx context.Context, in *token1.SearchPageRequest, opts ...grpc.CallOption) (*token1.SearchResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetToken(ctx context.Context, in *token.GetTokenRequest, opts ...grpc.CallOption) (*token.GetTokenResponse, error) {
	out := new(token.GetTokenResponse)
	err := c.cc.Invoke(ctx, Manager_GetToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) SearchPage(ctx context.Context, in *token1.SearchPageRequest, opts ...grpc.CallOption) (*token1.SearchResponse, error) {
	out := new(token1.SearchResponse)
	err := c.cc.Invoke(ctx, Manager_SearchPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	GetToken(context.Context, *token.GetTokenRequest) (*token.GetTokenResponse, error)
	SearchPage(context.Context, *token1.SearchPageRequest) (*token1.SearchResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetToken(context.Context, *token.GetTokenRequest) (*token.GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedManagerServer) SearchPage(context.Context, *token1.SearchPageRequest) (*token1.SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPage not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token.GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetToken(ctx, req.(*token.GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_SearchPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(token1.SearchPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).SearchPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_SearchPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).SearchPage(ctx, req.(*token1.SearchPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entrance.v1.token.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Manager_GetToken_Handler,
		},
		{
			MethodName: "SearchPage",
			Handler:    _Manager_SearchPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/entrance/v1/token/token.proto",
}
