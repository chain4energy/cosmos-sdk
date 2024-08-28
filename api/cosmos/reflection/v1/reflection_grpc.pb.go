// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/reflection/v1/reflection.proto

package reflectionv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReflectionService_FileDescriptors_FullMethodName = "/cosmos.reflection.v1.ReflectionService/FileDescriptors"
)

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Package cosmos.reflection.v1 provides support for inspecting protobuf
// file descriptors.
type ReflectionServiceClient interface {
	// FileDescriptors queries all the file descriptors in the app in order
	// to enable easier generation of dynamic clients.
	FileDescriptors(ctx context.Context, in *FileDescriptorsRequest, opts ...grpc.CallOption) (*FileDescriptorsResponse, error)
}

type reflectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReflectionServiceClient(cc grpc.ClientConnInterface) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) FileDescriptors(ctx context.Context, in *FileDescriptorsRequest, opts ...grpc.CallOption) (*FileDescriptorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileDescriptorsResponse)
	err := c.cc.Invoke(ctx, ReflectionService_FileDescriptors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
// All implementations must embed UnimplementedReflectionServiceServer
// for forward compatibility.
//
// Package cosmos.reflection.v1 provides support for inspecting protobuf
// file descriptors.
type ReflectionServiceServer interface {
	// FileDescriptors queries all the file descriptors in the app in order
	// to enable easier generation of dynamic clients.
	FileDescriptors(context.Context, *FileDescriptorsRequest) (*FileDescriptorsResponse, error)
	mustEmbedUnimplementedReflectionServiceServer()
}

// UnimplementedReflectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReflectionServiceServer struct{}

func (UnimplementedReflectionServiceServer) FileDescriptors(context.Context, *FileDescriptorsRequest) (*FileDescriptorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDescriptors not implemented")
}
func (UnimplementedReflectionServiceServer) mustEmbedUnimplementedReflectionServiceServer() {}
func (UnimplementedReflectionServiceServer) testEmbeddedByValue()                           {}

// UnsafeReflectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflectionServiceServer will
// result in compilation errors.
type UnsafeReflectionServiceServer interface {
	mustEmbedUnimplementedReflectionServiceServer()
}

func RegisterReflectionServiceServer(s grpc.ServiceRegistrar, srv ReflectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedReflectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReflectionService_ServiceDesc, srv)
}

func _ReflectionService_FileDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).FileDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReflectionService_FileDescriptors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).FileDescriptors(ctx, req.(*FileDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflectionService_ServiceDesc is the grpc.ServiceDesc for ReflectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.reflection.v1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileDescriptors",
			Handler:    _ReflectionService_FileDescriptors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/reflection/v1/reflection.proto",
}