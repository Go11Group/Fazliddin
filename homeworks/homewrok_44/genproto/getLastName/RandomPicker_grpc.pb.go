// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: RandomPicker.proto

package getLastName

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

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorClient interface {
	GetLastName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type generatorClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorClient(cc grpc.ClientConnInterface) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) GetLastName(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Generator/GetLastName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
// All implementations must embed UnimplementedGeneratorServer
// for forward compatibility
type GeneratorServer interface {
	GetLastName(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGeneratorServer()
}

// UnimplementedGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (UnimplementedGeneratorServer) GetLastName(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastName not implemented")
}
func (UnimplementedGeneratorServer) mustEmbedUnimplementedGeneratorServer() {}

// UnsafeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServer will
// result in compilation errors.
type UnsafeGeneratorServer interface {
	mustEmbedUnimplementedGeneratorServer()
}

func RegisterGeneratorServer(s grpc.ServiceRegistrar, srv GeneratorServer) {
	s.RegisterService(&Generator_ServiceDesc, srv)
}

func _Generator_GetLastName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).GetLastName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Generator/GetLastName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).GetLastName(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Generator_ServiceDesc is the grpc.ServiceDesc for Generator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastName",
			Handler:    _Generator_GetLastName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RandomPicker.proto",
}
