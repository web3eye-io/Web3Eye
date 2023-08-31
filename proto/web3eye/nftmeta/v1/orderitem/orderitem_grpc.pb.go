// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/nftmeta/v1/orderitem/orderitem.proto

package orderitem

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*CreateOrderItemResponse, error)
	CreateOrderItems(ctx context.Context, in *CreateOrderItemsRequest, opts ...grpc.CallOption) (*CreateOrderItemsResponse, error)
	UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*UpdateOrderItemResponse, error)
	GetOrderItem(ctx context.Context, in *GetOrderItemRequest, opts ...grpc.CallOption) (*GetOrderItemResponse, error)
	GetOrderItemOnly(ctx context.Context, in *GetOrderItemOnlyRequest, opts ...grpc.CallOption) (*GetOrderItemOnlyResponse, error)
	GetOrderItems(ctx context.Context, in *GetOrderItemsRequest, opts ...grpc.CallOption) (*GetOrderItemsResponse, error)
	ExistOrderItem(ctx context.Context, in *ExistOrderItemRequest, opts ...grpc.CallOption) (*ExistOrderItemResponse, error)
	ExistOrderItemConds(ctx context.Context, in *ExistOrderItemCondsRequest, opts ...grpc.CallOption) (*ExistOrderItemCondsResponse, error)
	CountOrderItems(ctx context.Context, in *CountOrderItemsRequest, opts ...grpc.CallOption) (*CountOrderItemsResponse, error)
	DeleteOrderItem(ctx context.Context, in *DeleteOrderItemRequest, opts ...grpc.CallOption) (*DeleteOrderItemResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*CreateOrderItemResponse, error) {
	out := new(CreateOrderItemResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/CreateOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateOrderItems(ctx context.Context, in *CreateOrderItemsRequest, opts ...grpc.CallOption) (*CreateOrderItemsResponse, error) {
	out := new(CreateOrderItemsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/CreateOrderItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*UpdateOrderItemResponse, error) {
	out := new(UpdateOrderItemResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/UpdateOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrderItem(ctx context.Context, in *GetOrderItemRequest, opts ...grpc.CallOption) (*GetOrderItemResponse, error) {
	out := new(GetOrderItemResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/GetOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrderItemOnly(ctx context.Context, in *GetOrderItemOnlyRequest, opts ...grpc.CallOption) (*GetOrderItemOnlyResponse, error) {
	out := new(GetOrderItemOnlyResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/GetOrderItemOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrderItems(ctx context.Context, in *GetOrderItemsRequest, opts ...grpc.CallOption) (*GetOrderItemsResponse, error) {
	out := new(GetOrderItemsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/GetOrderItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistOrderItem(ctx context.Context, in *ExistOrderItemRequest, opts ...grpc.CallOption) (*ExistOrderItemResponse, error) {
	out := new(ExistOrderItemResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/ExistOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistOrderItemConds(ctx context.Context, in *ExistOrderItemCondsRequest, opts ...grpc.CallOption) (*ExistOrderItemCondsResponse, error) {
	out := new(ExistOrderItemCondsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/ExistOrderItemConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountOrderItems(ctx context.Context, in *CountOrderItemsRequest, opts ...grpc.CallOption) (*CountOrderItemsResponse, error) {
	out := new(CountOrderItemsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/CountOrderItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteOrderItem(ctx context.Context, in *DeleteOrderItemRequest, opts ...grpc.CallOption) (*DeleteOrderItemResponse, error) {
	out := new(DeleteOrderItemResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.orderitem.Manager/DeleteOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateOrderItem(context.Context, *CreateOrderItemRequest) (*CreateOrderItemResponse, error)
	CreateOrderItems(context.Context, *CreateOrderItemsRequest) (*CreateOrderItemsResponse, error)
	UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*UpdateOrderItemResponse, error)
	GetOrderItem(context.Context, *GetOrderItemRequest) (*GetOrderItemResponse, error)
	GetOrderItemOnly(context.Context, *GetOrderItemOnlyRequest) (*GetOrderItemOnlyResponse, error)
	GetOrderItems(context.Context, *GetOrderItemsRequest) (*GetOrderItemsResponse, error)
	ExistOrderItem(context.Context, *ExistOrderItemRequest) (*ExistOrderItemResponse, error)
	ExistOrderItemConds(context.Context, *ExistOrderItemCondsRequest) (*ExistOrderItemCondsResponse, error)
	CountOrderItems(context.Context, *CountOrderItemsRequest) (*CountOrderItemsResponse, error)
	DeleteOrderItem(context.Context, *DeleteOrderItemRequest) (*DeleteOrderItemResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateOrderItem(context.Context, *CreateOrderItemRequest) (*CreateOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderItem not implemented")
}
func (UnimplementedManagerServer) CreateOrderItems(context.Context, *CreateOrderItemsRequest) (*CreateOrderItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderItems not implemented")
}
func (UnimplementedManagerServer) UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*UpdateOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderItem not implemented")
}
func (UnimplementedManagerServer) GetOrderItem(context.Context, *GetOrderItemRequest) (*GetOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderItem not implemented")
}
func (UnimplementedManagerServer) GetOrderItemOnly(context.Context, *GetOrderItemOnlyRequest) (*GetOrderItemOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderItemOnly not implemented")
}
func (UnimplementedManagerServer) GetOrderItems(context.Context, *GetOrderItemsRequest) (*GetOrderItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderItems not implemented")
}
func (UnimplementedManagerServer) ExistOrderItem(context.Context, *ExistOrderItemRequest) (*ExistOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderItem not implemented")
}
func (UnimplementedManagerServer) ExistOrderItemConds(context.Context, *ExistOrderItemCondsRequest) (*ExistOrderItemCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderItemConds not implemented")
}
func (UnimplementedManagerServer) CountOrderItems(context.Context, *CountOrderItemsRequest) (*CountOrderItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountOrderItems not implemented")
}
func (UnimplementedManagerServer) DeleteOrderItem(context.Context, *DeleteOrderItemRequest) (*DeleteOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderItem not implemented")
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

func _Manager_CreateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/CreateOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateOrderItem(ctx, req.(*CreateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateOrderItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateOrderItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/CreateOrderItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateOrderItems(ctx, req.(*CreateOrderItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/UpdateOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateOrderItem(ctx, req.(*UpdateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/GetOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrderItem(ctx, req.(*GetOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrderItemOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderItemOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrderItemOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/GetOrderItemOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrderItemOnly(ctx, req.(*GetOrderItemOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrderItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrderItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/GetOrderItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrderItems(ctx, req.(*GetOrderItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/ExistOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistOrderItem(ctx, req.(*ExistOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistOrderItemConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderItemCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistOrderItemConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/ExistOrderItemConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistOrderItemConds(ctx, req.(*ExistOrderItemCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountOrderItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountOrderItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountOrderItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/CountOrderItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountOrderItems(ctx, req.(*CountOrderItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.orderitem.Manager/DeleteOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteOrderItem(ctx, req.(*DeleteOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nftmeta.v1.orderitem.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderItem",
			Handler:    _Manager_CreateOrderItem_Handler,
		},
		{
			MethodName: "CreateOrderItems",
			Handler:    _Manager_CreateOrderItems_Handler,
		},
		{
			MethodName: "UpdateOrderItem",
			Handler:    _Manager_UpdateOrderItem_Handler,
		},
		{
			MethodName: "GetOrderItem",
			Handler:    _Manager_GetOrderItem_Handler,
		},
		{
			MethodName: "GetOrderItemOnly",
			Handler:    _Manager_GetOrderItemOnly_Handler,
		},
		{
			MethodName: "GetOrderItems",
			Handler:    _Manager_GetOrderItems_Handler,
		},
		{
			MethodName: "ExistOrderItem",
			Handler:    _Manager_ExistOrderItem_Handler,
		},
		{
			MethodName: "ExistOrderItemConds",
			Handler:    _Manager_ExistOrderItemConds_Handler,
		},
		{
			MethodName: "CountOrderItems",
			Handler:    _Manager_CountOrderItems_Handler,
		},
		{
			MethodName: "DeleteOrderItem",
			Handler:    _Manager_DeleteOrderItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/nftmeta/v1/orderitem/orderitem.proto",
}
