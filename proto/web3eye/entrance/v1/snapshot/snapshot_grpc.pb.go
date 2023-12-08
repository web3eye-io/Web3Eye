// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: web3eye/entrance/v1/snapshot/snapshot.proto

package snapshot

import (
	context "context"
	v1 "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"
	snapshot "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Manager_GetSnapshot_FullMethodName     = "/entrance.v1.snapshot.Manager/GetSnapshot"
	Manager_GetSnapshotOnly_FullMethodName = "/entrance.v1.snapshot.Manager/GetSnapshotOnly"
	Manager_GetSnapshots_FullMethodName    = "/entrance.v1.snapshot.Manager/GetSnapshots"
	Manager_CreateBackup_FullMethodName    = "/entrance.v1.snapshot.Manager/CreateBackup"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	GetSnapshot(ctx context.Context, in *snapshot.GetSnapshotRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotResponse, error)
	GetSnapshotOnly(ctx context.Context, in *snapshot.GetSnapshotOnlyRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotOnlyResponse, error)
	GetSnapshots(ctx context.Context, in *snapshot.GetSnapshotsRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotsResponse, error)
	CreateBackup(ctx context.Context, in *v1.CreateBackupRequest, opts ...grpc.CallOption) (*v1.CreateBackupResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetSnapshot(ctx context.Context, in *snapshot.GetSnapshotRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotResponse, error) {
	out := new(snapshot.GetSnapshotResponse)
	err := c.cc.Invoke(ctx, Manager_GetSnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshotOnly(ctx context.Context, in *snapshot.GetSnapshotOnlyRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotOnlyResponse, error) {
	out := new(snapshot.GetSnapshotOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetSnapshotOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSnapshots(ctx context.Context, in *snapshot.GetSnapshotsRequest, opts ...grpc.CallOption) (*snapshot.GetSnapshotsResponse, error) {
	out := new(snapshot.GetSnapshotsResponse)
	err := c.cc.Invoke(ctx, Manager_GetSnapshots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateBackup(ctx context.Context, in *v1.CreateBackupRequest, opts ...grpc.CallOption) (*v1.CreateBackupResponse, error) {
	out := new(v1.CreateBackupResponse)
	err := c.cc.Invoke(ctx, Manager_CreateBackup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	GetSnapshot(context.Context, *snapshot.GetSnapshotRequest) (*snapshot.GetSnapshotResponse, error)
	GetSnapshotOnly(context.Context, *snapshot.GetSnapshotOnlyRequest) (*snapshot.GetSnapshotOnlyResponse, error)
	GetSnapshots(context.Context, *snapshot.GetSnapshotsRequest) (*snapshot.GetSnapshotsResponse, error)
	CreateBackup(context.Context, *v1.CreateBackupRequest) (*v1.CreateBackupResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetSnapshot(context.Context, *snapshot.GetSnapshotRequest) (*snapshot.GetSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshot not implemented")
}
func (UnimplementedManagerServer) GetSnapshotOnly(context.Context, *snapshot.GetSnapshotOnlyRequest) (*snapshot.GetSnapshotOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshotOnly not implemented")
}
func (UnimplementedManagerServer) GetSnapshots(context.Context, *snapshot.GetSnapshotsRequest) (*snapshot.GetSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshots not implemented")
}
func (UnimplementedManagerServer) CreateBackup(context.Context, *v1.CreateBackupRequest) (*v1.CreateBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackup not implemented")
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

func _Manager_GetSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(snapshot.GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshot(ctx, req.(*snapshot.GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSnapshotOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(snapshot.GetSnapshotOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshotOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetSnapshotOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshotOnly(ctx, req.(*snapshot.GetSnapshotOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(snapshot.GetSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetSnapshots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSnapshots(ctx, req.(*snapshot.GetSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CreateBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateBackup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateBackup(ctx, req.(*v1.CreateBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entrance.v1.snapshot.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSnapshot",
			Handler:    _Manager_GetSnapshot_Handler,
		},
		{
			MethodName: "GetSnapshotOnly",
			Handler:    _Manager_GetSnapshotOnly_Handler,
		},
		{
			MethodName: "GetSnapshots",
			Handler:    _Manager_GetSnapshots_Handler,
		},
		{
			MethodName: "CreateBackup",
			Handler:    _Manager_CreateBackup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web3eye/entrance/v1/snapshot/snapshot.proto",
}
