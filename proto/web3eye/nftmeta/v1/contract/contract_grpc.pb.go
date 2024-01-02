// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: web3eye/nftmeta/v1/contract/contract.proto

package contract

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
	Manager_CreateContract_FullMethodName     = "/nftmeta.v1.contract.Manager/CreateContract"
	Manager_CreateContracts_FullMethodName    = "/nftmeta.v1.contract.Manager/CreateContracts"
	Manager_UpdateContract_FullMethodName     = "/nftmeta.v1.contract.Manager/UpdateContract"
	Manager_UpsertContract_FullMethodName     = "/nftmeta.v1.contract.Manager/UpsertContract"
	Manager_GetContract_FullMethodName        = "/nftmeta.v1.contract.Manager/GetContract"
	Manager_GetContractOnly_FullMethodName    = "/nftmeta.v1.contract.Manager/GetContractOnly"
	Manager_GetContracts_FullMethodName       = "/nftmeta.v1.contract.Manager/GetContracts"
	Manager_ExistContract_FullMethodName      = "/nftmeta.v1.contract.Manager/ExistContract"
	Manager_ExistContractConds_FullMethodName = "/nftmeta.v1.contract.Manager/ExistContractConds"
	Manager_DeleteContract_FullMethodName     = "/nftmeta.v1.contract.Manager/DeleteContract"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error)
	CreateContracts(ctx context.Context, in *CreateContractsRequest, opts ...grpc.CallOption) (*CreateContractsResponse, error)
	UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*UpdateContractResponse, error)
	UpsertContract(ctx context.Context, in *UpsertContractRequest, opts ...grpc.CallOption) (*UpsertContractResponse, error)
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
	GetContractOnly(ctx context.Context, in *GetContractOnlyRequest, opts ...grpc.CallOption) (*GetContractOnlyResponse, error)
	GetContracts(ctx context.Context, in *GetContractsRequest, opts ...grpc.CallOption) (*GetContractsResponse, error)
	ExistContract(ctx context.Context, in *ExistContractRequest, opts ...grpc.CallOption) (*ExistContractResponse, error)
	ExistContractConds(ctx context.Context, in *ExistContractCondsRequest, opts ...grpc.CallOption) (*ExistContractCondsResponse, error)
	DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateContract(ctx context.Context, in *CreateContractRequest, opts ...grpc.CallOption) (*CreateContractResponse, error) {
	out := new(CreateContractResponse)
	err := c.cc.Invoke(ctx, Manager_CreateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateContracts(ctx context.Context, in *CreateContractsRequest, opts ...grpc.CallOption) (*CreateContractsResponse, error) {
	out := new(CreateContractsResponse)
	err := c.cc.Invoke(ctx, Manager_CreateContracts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateContract(ctx context.Context, in *UpdateContractRequest, opts ...grpc.CallOption) (*UpdateContractResponse, error) {
	out := new(UpdateContractResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpsertContract(ctx context.Context, in *UpsertContractRequest, opts ...grpc.CallOption) (*UpsertContractResponse, error) {
	out := new(UpsertContractResponse)
	err := c.cc.Invoke(ctx, Manager_UpsertContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, Manager_GetContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetContractOnly(ctx context.Context, in *GetContractOnlyRequest, opts ...grpc.CallOption) (*GetContractOnlyResponse, error) {
	out := new(GetContractOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetContractOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetContracts(ctx context.Context, in *GetContractsRequest, opts ...grpc.CallOption) (*GetContractsResponse, error) {
	out := new(GetContractsResponse)
	err := c.cc.Invoke(ctx, Manager_GetContracts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistContract(ctx context.Context, in *ExistContractRequest, opts ...grpc.CallOption) (*ExistContractResponse, error) {
	out := new(ExistContractResponse)
	err := c.cc.Invoke(ctx, Manager_ExistContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistContractConds(ctx context.Context, in *ExistContractCondsRequest, opts ...grpc.CallOption) (*ExistContractCondsResponse, error) {
	out := new(ExistContractCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistContractConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteContract(ctx context.Context, in *DeleteContractRequest, opts ...grpc.CallOption) (*DeleteContractResponse, error) {
	out := new(DeleteContractResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error)
	CreateContracts(context.Context, *CreateContractsRequest) (*CreateContractsResponse, error)
	UpdateContract(context.Context, *UpdateContractRequest) (*UpdateContractResponse, error)
	UpsertContract(context.Context, *UpsertContractRequest) (*UpsertContractResponse, error)
	GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error)
	GetContractOnly(context.Context, *GetContractOnlyRequest) (*GetContractOnlyResponse, error)
	GetContracts(context.Context, *GetContractsRequest) (*GetContractsResponse, error)
	ExistContract(context.Context, *ExistContractRequest) (*ExistContractResponse, error)
	ExistContractConds(context.Context, *ExistContractCondsRequest) (*ExistContractCondsResponse, error)
	DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateContract(context.Context, *CreateContractRequest) (*CreateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContract not implemented")
}
func (UnimplementedManagerServer) CreateContracts(context.Context, *CreateContractsRequest) (*CreateContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContracts not implemented")
}
func (UnimplementedManagerServer) UpdateContract(context.Context, *UpdateContractRequest) (*UpdateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContract not implemented")
}
func (UnimplementedManagerServer) UpsertContract(context.Context, *UpsertContractRequest) (*UpsertContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertContract not implemented")
}
func (UnimplementedManagerServer) GetContract(context.Context, *GetContractRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedManagerServer) GetContractOnly(context.Context, *GetContractOnlyRequest) (*GetContractOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractOnly not implemented")
}
func (UnimplementedManagerServer) GetContracts(context.Context, *GetContractsRequest) (*GetContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContracts not implemented")
}
func (UnimplementedManagerServer) ExistContract(context.Context, *ExistContractRequest) (*ExistContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistContract not implemented")
}
func (UnimplementedManagerServer) ExistContractConds(context.Context, *ExistContractCondsRequest) (*ExistContractCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistContractConds not implemented")
}
func (UnimplementedManagerServer) DeleteContract(context.Context, *DeleteContractRequest) (*DeleteContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContract not implemented")
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

func _Manager_CreateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateContract(ctx, req.(*CreateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateContracts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateContracts(ctx, req.(*CreateContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateContract(ctx, req.(*UpdateContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpsertContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpsertContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpsertContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpsertContract(ctx, req.(*UpsertContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetContractOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContractOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetContractOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContractOnly(ctx, req.(*GetContractOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetContracts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetContracts(ctx, req.(*GetContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistContract(ctx, req.(*ExistContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistContractConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistContractCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistContractConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistContractConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistContractConds(ctx, req.(*ExistContractCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteContract(ctx, req.(*DeleteContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nftmeta.v1.contract.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContract",
			Handler:    _Manager_CreateContract_Handler,
		},
		{
			MethodName: "CreateContracts",
			Handler:    _Manager_CreateContracts_Handler,
		},
		{
			MethodName: "UpdateContract",
			Handler:    _Manager_UpdateContract_Handler,
		},
		{
			MethodName: "UpsertContract",
			Handler:    _Manager_UpsertContract_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _Manager_GetContract_Handler,
		},
		{
			MethodName: "GetContractOnly",
			Handler:    _Manager_GetContractOnly_Handler,
		},
		{
			MethodName: "GetContracts",
			Handler:    _Manager_GetContracts_Handler,
		},
		{
			MethodName: "ExistContract",
			Handler:    _Manager_ExistContract_Handler,
		},
		{
			MethodName: "ExistContractConds",
			Handler:    _Manager_ExistContractConds_Handler,
		},
		{
			MethodName: "DeleteContract",
			Handler:    _Manager_DeleteContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/nftmeta/v1/contract/contract.proto",
}
