// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/entrance/v1/retriever/retriever.proto

package retriever

import (
	context "context"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	StartRetrieve(ctx context.Context, in *v1.StartRetrieveRequest, opts ...grpc.CallOption) (*v1.StartRetrieveResponse, error)
	StatRetrieve(ctx context.Context, in *v1.StatRetrieveRequest, opts ...grpc.CallOption) (*v1.StatRetrieveResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) StartRetrieve(ctx context.Context, in *v1.StartRetrieveRequest, opts ...grpc.CallOption) (*v1.StartRetrieveResponse, error) {
	out := new(v1.StartRetrieveResponse)
	err := c.cc.Invoke(ctx, "/entrance.v1.retriever1.Manager/StartRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) StatRetrieve(ctx context.Context, in *v1.StatRetrieveRequest, opts ...grpc.CallOption) (*v1.StatRetrieveResponse, error) {
	out := new(v1.StatRetrieveResponse)
	err := c.cc.Invoke(ctx, "/entrance.v1.retriever1.Manager/StatRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	StartRetrieve(context.Context, *v1.StartRetrieveRequest) (*v1.StartRetrieveResponse, error)
	StatRetrieve(context.Context, *v1.StatRetrieveRequest) (*v1.StatRetrieveResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) StartRetrieve(context.Context, *v1.StartRetrieveRequest) (*v1.StartRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRetrieve not implemented")
}
func (UnimplementedManagerServer) StatRetrieve(context.Context, *v1.StatRetrieveRequest) (*v1.StatRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatRetrieve not implemented")
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

func _Manager_StartRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.StartRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StartRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entrance.v1.retriever1.Manager/StartRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StartRetrieve(ctx, req.(*v1.StartRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_StatRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.StatRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StatRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entrance.v1.retriever1.Manager/StatRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StatRetrieve(ctx, req.(*v1.StatRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entrance.v1.retriever1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartRetrieve",
			Handler:    _Manager_StartRetrieve_Handler,
		},
		{
			MethodName: "StatRetrieve",
			Handler:    _Manager_StatRetrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/entrance/v1/retriever/retriever.proto",
}
