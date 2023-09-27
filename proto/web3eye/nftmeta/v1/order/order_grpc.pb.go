// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: web3eye/nftmeta/v1/order/order.proto

package order

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
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrderOnly(ctx context.Context, in *GetOrderOnlyRequest, opts ...grpc.CallOption) (*GetOrderOnlyResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	ExistOrder(ctx context.Context, in *ExistOrderRequest, opts ...grpc.CallOption) (*ExistOrderResponse, error)
	ExistOrderConds(ctx context.Context, in *ExistOrderCondsRequest, opts ...grpc.CallOption) (*ExistOrderCondsResponse, error)
	CountOrders(ctx context.Context, in *CountOrdersRequest, opts ...grpc.CallOption) (*CountOrdersResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error) {
	out := new(CreateOrdersResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/CreateOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrderOnly(ctx context.Context, in *GetOrderOnlyRequest, opts ...grpc.CallOption) (*GetOrderOnlyResponse, error) {
	out := new(GetOrderOnlyResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/GetOrderOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistOrder(ctx context.Context, in *ExistOrderRequest, opts ...grpc.CallOption) (*ExistOrderResponse, error) {
	out := new(ExistOrderResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/ExistOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistOrderConds(ctx context.Context, in *ExistOrderCondsRequest, opts ...grpc.CallOption) (*ExistOrderCondsResponse, error) {
	out := new(ExistOrderCondsResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/ExistOrderConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountOrders(ctx context.Context, in *CountOrdersRequest, opts ...grpc.CallOption) (*CountOrdersResponse, error) {
	out := new(CountOrdersResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/CountOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, "/nftmeta.v1.order.Manager/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrderOnly(context.Context, *GetOrderOnlyRequest) (*GetOrderOnlyResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error)
	ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error)
	CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedManagerServer) CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedManagerServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedManagerServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedManagerServer) GetOrderOnly(context.Context, *GetOrderOnlyRequest) (*GetOrderOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderOnly not implemented")
}
func (UnimplementedManagerServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedManagerServer) ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrder not implemented")
}
func (UnimplementedManagerServer) ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderConds not implemented")
}
func (UnimplementedManagerServer) CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountOrders not implemented")
}
func (UnimplementedManagerServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
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

func _Manager_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/CreateOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateOrders(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrderOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrderOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/GetOrderOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrderOnly(ctx, req.(*GetOrderOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/ExistOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistOrder(ctx, req.(*ExistOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistOrderConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistOrderConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/ExistOrderConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistOrderConds(ctx, req.(*ExistOrderCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/CountOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountOrders(ctx, req.(*CountOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nftmeta.v1.order.Manager/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nftmeta.v1.order.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Manager_CreateOrder_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _Manager_CreateOrders_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Manager_UpdateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Manager_GetOrder_Handler,
		},
		{
			MethodName: "GetOrderOnly",
			Handler:    _Manager_GetOrderOnly_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Manager_GetOrders_Handler,
		},
		{
			MethodName: "ExistOrder",
			Handler:    _Manager_ExistOrder_Handler,
		},
		{
			MethodName: "ExistOrderConds",
			Handler:    _Manager_ExistOrderConds_Handler,
		},
		{
			MethodName: "CountOrders",
			Handler:    _Manager_CountOrders_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Manager_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/nftmeta/v1/order/order.proto",
}
