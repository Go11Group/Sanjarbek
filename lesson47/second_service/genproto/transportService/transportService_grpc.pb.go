// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: transportService.proto

package transportService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TransportService_GetBusSchedule_FullMethodName   = "/transportService.TransportService/GetBusSchedule"
	TransportService_TrackBusLocation_FullMethodName = "/transportService.TransportService/TrackBusLocation"
	TransportService_ReportTrafficJam_FullMethodName = "/transportService.TransportService/ReportTrafficJam"
)

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	GetBusSchedule(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Schedule, error)
	TrackBusLocation(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Location, error)
	ReportTrafficJam(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Traffic, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) GetBusSchedule(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Schedule, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Schedule)
	err := c.cc.Invoke(ctx, TransportService_GetBusSchedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) TrackBusLocation(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Location, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Location)
	err := c.cc.Invoke(ctx, TransportService_TrackBusLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) ReportTrafficJam(ctx context.Context, in *Transport, opts ...grpc.CallOption) (*Traffic, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Traffic)
	err := c.cc.Invoke(ctx, TransportService_ReportTrafficJam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility
type TransportServiceServer interface {
	GetBusSchedule(context.Context, *Transport) (*Schedule, error)
	TrackBusLocation(context.Context, *Transport) (*Location, error)
	ReportTrafficJam(context.Context, *Transport) (*Traffic, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (UnimplementedTransportServiceServer) GetBusSchedule(context.Context, *Transport) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusSchedule not implemented")
}
func (UnimplementedTransportServiceServer) TrackBusLocation(context.Context, *Transport) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackBusLocation not implemented")
}
func (UnimplementedTransportServiceServer) ReportTrafficJam(context.Context, *Transport) (*Traffic, error) {
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
	in := new(Transport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_GetBusSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetBusSchedule(ctx, req.(*Transport))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_TrackBusLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_TrackBusLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).TrackBusLocation(ctx, req.(*Transport))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_ReportTrafficJam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_ReportTrafficJam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).ReportTrafficJam(ctx, req.(*Transport))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transportService.TransportService",
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
	Metadata: "transportService.proto",
}
