// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: services/temporal_service.proto

package services

import (
	context "context"
	resources "github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TemporalServiceClient is the client API for TemporalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemporalServiceClient interface {
	// Basic example of the Temporal key components: workflow, activity and worker
	CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error)
	GetWithdrawal(ctx context.Context, in *GetWithdrawalRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error)
	UpdateWithdrawal(ctx context.Context, in *UpdateWithdrawRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error)
}

type temporalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemporalServiceClient(cc grpc.ClientConnInterface) TemporalServiceClient {
	return &temporalServiceClient{cc}
}

func (c *temporalServiceClient) CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error) {
	out := new(resources.Withdrawal)
	err := c.cc.Invoke(ctx, "/pepeunlimited.grpckit.services.TemporalService/CreateWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temporalServiceClient) GetWithdrawal(ctx context.Context, in *GetWithdrawalRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error) {
	out := new(resources.Withdrawal)
	err := c.cc.Invoke(ctx, "/pepeunlimited.grpckit.services.TemporalService/GetWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temporalServiceClient) UpdateWithdrawal(ctx context.Context, in *UpdateWithdrawRequest, opts ...grpc.CallOption) (*resources.Withdrawal, error) {
	out := new(resources.Withdrawal)
	err := c.cc.Invoke(ctx, "/pepeunlimited.grpckit.services.TemporalService/UpdateWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemporalServiceServer is the server API for TemporalService service.
// All implementations must embed UnimplementedTemporalServiceServer
// for forward compatibility
type TemporalServiceServer interface {
	// Basic example of the Temporal key components: workflow, activity and worker
	CreateWithdraw(context.Context, *CreateWithdrawRequest) (*resources.Withdrawal, error)
	GetWithdrawal(context.Context, *GetWithdrawalRequest) (*resources.Withdrawal, error)
	UpdateWithdrawal(context.Context, *UpdateWithdrawRequest) (*resources.Withdrawal, error)
	mustEmbedUnimplementedTemporalServiceServer()
}

// UnimplementedTemporalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemporalServiceServer struct {
}

func (UnimplementedTemporalServiceServer) CreateWithdraw(context.Context, *CreateWithdrawRequest) (*resources.Withdrawal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWithdraw not implemented")
}
func (UnimplementedTemporalServiceServer) GetWithdrawal(context.Context, *GetWithdrawalRequest) (*resources.Withdrawal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawal not implemented")
}
func (UnimplementedTemporalServiceServer) UpdateWithdrawal(context.Context, *UpdateWithdrawRequest) (*resources.Withdrawal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithdrawal not implemented")
}
func (UnimplementedTemporalServiceServer) mustEmbedUnimplementedTemporalServiceServer() {}

// UnsafeTemporalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemporalServiceServer will
// result in compilation errors.
type UnsafeTemporalServiceServer interface {
	mustEmbedUnimplementedTemporalServiceServer()
}

func RegisterTemporalServiceServer(s grpc.ServiceRegistrar, srv TemporalServiceServer) {
	s.RegisterService(&TemporalService_ServiceDesc, srv)
}

func _TemporalService_CreateWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalServiceServer).CreateWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pepeunlimited.grpckit.services.TemporalService/CreateWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalServiceServer).CreateWithdraw(ctx, req.(*CreateWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemporalService_GetWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalServiceServer).GetWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pepeunlimited.grpckit.services.TemporalService/GetWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalServiceServer).GetWithdrawal(ctx, req.(*GetWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemporalService_UpdateWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalServiceServer).UpdateWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pepeunlimited.grpckit.services.TemporalService/UpdateWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalServiceServer).UpdateWithdrawal(ctx, req.(*UpdateWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemporalService_ServiceDesc is the grpc.ServiceDesc for TemporalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemporalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pepeunlimited.grpckit.services.TemporalService",
	HandlerType: (*TemporalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWithdraw",
			Handler:    _TemporalService_CreateWithdraw_Handler,
		},
		{
			MethodName: "GetWithdrawal",
			Handler:    _TemporalService_GetWithdrawal_Handler,
		},
		{
			MethodName: "UpdateWithdrawal",
			Handler:    _TemporalService_UpdateWithdrawal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/temporal_service.proto",
}