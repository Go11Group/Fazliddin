// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: transport.proto

package transport

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

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	GetBusSchedule(ctx context.Context, in *BusNumber, opts ...grpc.CallOption) (*BusSchudle, error)
	TrackBusLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*BusWithLocations, error)
	ReportTrafficJam(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Response, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) GetBusSchedule(ctx context.Context, in *BusNumber, opts ...grpc.CallOption) (*BusSchudle, error) {
	out := new(BusSchudle)
	err := c.cc.Invoke(ctx, "/transport.TransportService/GetBusSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) TrackBusLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*BusWithLocations, error) {
	out := new(BusWithLocations)
	err := c.cc.Invoke(ctx, "/transport.TransportService/TrackBusLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) ReportTrafficJam(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/transport.TransportService/ReportTrafficJam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility
type TransportServiceServer interface {
	GetBusSchedule(context.Context, *BusNumber) (*BusSchudle, error)
	TrackBusLocation(context.Context, *Location) (*BusWithLocations, error)
	ReportTrafficJam(context.Context, *Location) (*Response, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (UnimplementedTransportServiceServer) GetBusSchedule(context.Context, *BusNumber) (*BusSchudle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusSchedule not implemented")
}
func (UnimplementedTransportServiceServer) TrackBusLocation(context.Context, *Location) (*BusWithLocations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackBusLocation not implemented")
}
func (UnimplementedTransportServiceServer) ReportTrafficJam(context.Context, *Location) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTrafficJam not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_GetBusSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportService/GetBusSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, req.(*BusNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_TrackBusLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportService/TrackBusLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_ReportTrafficJam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TransportService/ReportTrafficJam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusSchedule",
			Handler:    _TransportService_GetBusSchedule_Handler,
		},
		{
			MethodName: "TrackBusLocation",
			Handler:    _TransportService_TrackBusLocation_Handler,
		},
		{
			MethodName: "ReportTrafficJam",
			Handler:    _TransportService_ReportTrafficJam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
