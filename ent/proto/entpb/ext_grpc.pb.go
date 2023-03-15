// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: entpb/ext.proto

package entpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExtService_TopUser_FullMethodName = "/entpb.ExtService/TopUser"
)

// ExtServiceClient is the client API for ExtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtServiceClient interface {
	TopUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error)
}

type extServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtServiceClient(cc grpc.ClientConnInterface) ExtServiceClient {
	return &extServiceClient{cc}
}

func (c *extServiceClient) TopUser(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, ExtService_TopUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtServiceServer is the server API for ExtService service.
// All implementations must embed UnimplementedExtServiceServer
// for forward compatibility
type ExtServiceServer interface {
	TopUser(context.Context, *emptypb.Empty) (*User, error)
	mustEmbedUnimplementedExtServiceServer()
}

// UnimplementedExtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtServiceServer struct {
}

func (UnimplementedExtServiceServer) TopUser(context.Context, *emptypb.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopUser not implemented")
}
func (UnimplementedExtServiceServer) mustEmbedUnimplementedExtServiceServer() {}

// UnsafeExtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtServiceServer will
// result in compilation errors.
type UnsafeExtServiceServer interface {
	mustEmbedUnimplementedExtServiceServer()
}

func RegisterExtServiceServer(s grpc.ServiceRegistrar, srv ExtServiceServer) {
	s.RegisterService(&ExtService_ServiceDesc, srv)
}

func _ExtService_TopUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtServiceServer).TopUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtService_TopUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtServiceServer).TopUser(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtService_ServiceDesc is the grpc.ServiceDesc for ExtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.ExtService",
	HandlerType: (*ExtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TopUser",
			Handler:    _ExtService_TopUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/ext.proto",
}
