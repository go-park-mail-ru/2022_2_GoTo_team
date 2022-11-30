// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: profile.proto

package userProfileServiceGrpcProtos

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

// AuthSessionServiceClient is the client API for AuthSessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthSessionServiceClient interface {
	UpdateEmailBySession(ctx context.Context, in *UpdateEmailData, opts ...grpc.CallOption) (*Nothing, error)
}

type authSessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthSessionServiceClient(cc grpc.ClientConnInterface) AuthSessionServiceClient {
	return &authSessionServiceClient{cc}
}

func (c *authSessionServiceClient) UpdateEmailBySession(ctx context.Context, in *UpdateEmailData, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/profile.AuthSessionService/UpdateEmailBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSessionServiceServer is the server API for AuthSessionService service.
// All implementations must embed UnimplementedAuthSessionServiceServer
// for forward compatibility
type AuthSessionServiceServer interface {
	UpdateEmailBySession(context.Context, *UpdateEmailData) (*Nothing, error)
	mustEmbedUnimplementedAuthSessionServiceServer()
}

// UnimplementedAuthSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthSessionServiceServer struct {
}

func (UnimplementedAuthSessionServiceServer) UpdateEmailBySession(context.Context, *UpdateEmailData) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailBySession not implemented")
}
func (UnimplementedAuthSessionServiceServer) mustEmbedUnimplementedAuthSessionServiceServer() {}

// UnsafeAuthSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthSessionServiceServer will
// result in compilation errors.
type UnsafeAuthSessionServiceServer interface {
	mustEmbedUnimplementedAuthSessionServiceServer()
}

func RegisterAuthSessionServiceServer(s grpc.ServiceRegistrar, srv AuthSessionServiceServer) {
	s.RegisterService(&AuthSessionService_ServiceDesc, srv)
}

func _AuthSessionService_UpdateEmailBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSessionServiceServer).UpdateEmailBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AuthSessionService/UpdateEmailBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSessionServiceServer).UpdateEmailBySession(ctx, req.(*UpdateEmailData))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthSessionService_ServiceDesc is the grpc.ServiceDesc for AuthSessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthSessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.AuthSessionService",
	HandlerType: (*AuthSessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateEmailBySession",
			Handler:    _AuthSessionService_UpdateEmailBySession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
