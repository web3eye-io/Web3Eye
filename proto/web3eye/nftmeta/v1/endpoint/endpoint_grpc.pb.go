// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: web3eye/nftmeta/v1/endpoint/endpoint.proto

package endpoint

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Manager_CreateEndpoint_FullMethodName     = "/nftmeta.v1.endpoint.Manager/CreateEndpoint"
	Manager_CreateEndpoints_FullMethodName    = "/nftmeta.v1.endpoint.Manager/CreateEndpoints"
	Manager_UpdateEndpoint_FullMethodName     = "/nftmeta.v1.endpoint.Manager/UpdateEndpoint"
	Manager_UpdateEndpoints_FullMethodName    = "/nftmeta.v1.endpoint.Manager/UpdateEndpoints"
	Manager_GetEndpoint_FullMethodName        = "/nftmeta.v1.endpoint.Manager/GetEndpoint"
	Manager_GetEndpointOnly_FullMethodName    = "/nftmeta.v1.endpoint.Manager/GetEndpointOnly"
	Manager_GetEndpoints_FullMethodName       = "/nftmeta.v1.endpoint.Manager/GetEndpoints"
	Manager_ExistEndpoint_FullMethodName      = "/nftmeta.v1.endpoint.Manager/ExistEndpoint"
	Manager_ExistEndpointConds_FullMethodName = "/nftmeta.v1.endpoint.Manager/ExistEndpointConds"
	Manager_DeleteEndpoint_FullMethodName     = "/nftmeta.v1.endpoint.Manager/DeleteEndpoint"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateEndpoint(ctx context.Context, in *CreateEndpointRequest, opts ...grpc.CallOption) (*CreateEndpointResponse, error)
	CreateEndpoints(ctx context.Context, in *CreateEndpointsRequest, opts ...grpc.CallOption) (*CreateEndpointsResponse, error)
	UpdateEndpoint(ctx context.Context, in *UpdateEndpointRequest, opts ...grpc.CallOption) (*UpdateEndpointResponse, error)
	UpdateEndpoints(ctx context.Context, in *UpdateEndpointsRequest, opts ...grpc.CallOption) (*UpdateEndpointsResponse, error)
	GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointResponse, error)
	GetEndpointOnly(ctx context.Context, in *GetEndpointOnlyRequest, opts ...grpc.CallOption) (*GetEndpointOnlyResponse, error)
	GetEndpoints(ctx context.Context, in *GetEndpointsRequest, opts ...grpc.CallOption) (*GetEndpointsResponse, error)
	ExistEndpoint(ctx context.Context, in *ExistEndpointRequest, opts ...grpc.CallOption) (*ExistEndpointResponse, error)
	ExistEndpointConds(ctx context.Context, in *ExistEndpointCondsRequest, opts ...grpc.CallOption) (*ExistEndpointCondsResponse, error)
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*DeleteEndpointResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateEndpoint(ctx context.Context, in *CreateEndpointRequest, opts ...grpc.CallOption) (*CreateEndpointResponse, error) {
	out := new(CreateEndpointResponse)
	err := c.cc.Invoke(ctx, Manager_CreateEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateEndpoints(ctx context.Context, in *CreateEndpointsRequest, opts ...grpc.CallOption) (*CreateEndpointsResponse, error) {
	out := new(CreateEndpointsResponse)
	err := c.cc.Invoke(ctx, Manager_CreateEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateEndpoint(ctx context.Context, in *UpdateEndpointRequest, opts ...grpc.CallOption) (*UpdateEndpointResponse, error) {
	out := new(UpdateEndpointResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateEndpoints(ctx context.Context, in *UpdateEndpointsRequest, opts ...grpc.CallOption) (*UpdateEndpointsResponse, error) {
	out := new(UpdateEndpointsResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...grpc.CallOption) (*GetEndpointResponse, error) {
	out := new(GetEndpointResponse)
	err := c.cc.Invoke(ctx, Manager_GetEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetEndpointOnly(ctx context.Context, in *GetEndpointOnlyRequest, opts ...grpc.CallOption) (*GetEndpointOnlyResponse, error) {
	out := new(GetEndpointOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetEndpointOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetEndpoints(ctx context.Context, in *GetEndpointsRequest, opts ...grpc.CallOption) (*GetEndpointsResponse, error) {
	out := new(GetEndpointsResponse)
	err := c.cc.Invoke(ctx, Manager_GetEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistEndpoint(ctx context.Context, in *ExistEndpointRequest, opts ...grpc.CallOption) (*ExistEndpointResponse, error) {
	out := new(ExistEndpointResponse)
	err := c.cc.Invoke(ctx, Manager_ExistEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistEndpointConds(ctx context.Context, in *ExistEndpointCondsRequest, opts ...grpc.CallOption) (*ExistEndpointCondsResponse, error) {
	out := new(ExistEndpointCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistEndpointConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*DeleteEndpointResponse, error) {
	out := new(DeleteEndpointResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateEndpoint(context.Context, *CreateEndpointRequest) (*CreateEndpointResponse, error)
	CreateEndpoints(context.Context, *CreateEndpointsRequest) (*CreateEndpointsResponse, error)
	UpdateEndpoint(context.Context, *UpdateEndpointRequest) (*UpdateEndpointResponse, error)
	UpdateEndpoints(context.Context, *UpdateEndpointsRequest) (*UpdateEndpointsResponse, error)
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointResponse, error)
	GetEndpointOnly(context.Context, *GetEndpointOnlyRequest) (*GetEndpointOnlyResponse, error)
	GetEndpoints(context.Context, *GetEndpointsRequest) (*GetEndpointsResponse, error)
	ExistEndpoint(context.Context, *ExistEndpointRequest) (*ExistEndpointResponse, error)
	ExistEndpointConds(context.Context, *ExistEndpointCondsRequest) (*ExistEndpointCondsResponse, error)
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateEndpoint(context.Context, *CreateEndpointRequest) (*CreateEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEndpoint not implemented")
}
func (UnimplementedManagerServer) CreateEndpoints(context.Context, *CreateEndpointsRequest) (*CreateEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEndpoints not implemented")
}
func (UnimplementedManagerServer) UpdateEndpoint(context.Context, *UpdateEndpointRequest) (*UpdateEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpoint not implemented")
}
func (UnimplementedManagerServer) UpdateEndpoints(context.Context, *UpdateEndpointsRequest) (*UpdateEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEndpoints not implemented")
}
func (UnimplementedManagerServer) GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoint not implemented")
}
func (UnimplementedManagerServer) GetEndpointOnly(context.Context, *GetEndpointOnlyRequest) (*GetEndpointOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpointOnly not implemented")
}
func (UnimplementedManagerServer) GetEndpoints(context.Context, *GetEndpointsRequest) (*GetEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndpoints not implemented")
}
func (UnimplementedManagerServer) ExistEndpoint(context.Context, *ExistEndpointRequest) (*ExistEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistEndpoint not implemented")
}
func (UnimplementedManagerServer) ExistEndpointConds(context.Context, *ExistEndpointCondsRequest) (*ExistEndpointCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistEndpointConds not implemented")
}
func (UnimplementedManagerServer) DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEndpoint not implemented")
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

func _Manager_CreateEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateEndpoint(ctx, req.(*CreateEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateEndpoints(ctx, req.(*CreateEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateEndpoint(ctx, req.(*UpdateEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateEndpoints(ctx, req.(*UpdateEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetEndpoint(ctx, req.(*GetEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetEndpointOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetEndpointOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetEndpointOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetEndpointOnly(ctx, req.(*GetEndpointOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetEndpoints(ctx, req.(*GetEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistEndpoint(ctx, req.(*ExistEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistEndpointConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistEndpointCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistEndpointConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistEndpointConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistEndpointConds(ctx, req.(*ExistEndpointCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nftmeta.v1.endpoint.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEndpoint",
			Handler:    _Manager_CreateEndpoint_Handler,
		},
		{
			MethodName: "CreateEndpoints",
			Handler:    _Manager_CreateEndpoints_Handler,
		},
		{
			MethodName: "UpdateEndpoint",
			Handler:    _Manager_UpdateEndpoint_Handler,
		},
		{
			MethodName: "UpdateEndpoints",
			Handler:    _Manager_UpdateEndpoints_Handler,
		},
		{
			MethodName: "GetEndpoint",
			Handler:    _Manager_GetEndpoint_Handler,
		},
		{
			MethodName: "GetEndpointOnly",
			Handler:    _Manager_GetEndpointOnly_Handler,
		},
		{
			MethodName: "GetEndpoints",
			Handler:    _Manager_GetEndpoints_Handler,
		},
		{
			MethodName: "ExistEndpoint",
			Handler:    _Manager_ExistEndpoint_Handler,
		},
		{
			MethodName: "ExistEndpointConds",
			Handler:    _Manager_ExistEndpointConds_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Manager_DeleteEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/nftmeta/v1/endpoint/endpoint.proto",
}
