// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: counter.proto

package counter

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

const (
	CounterService_Counter_FullMethodName = "/clicker.CounterService/Counter"
)

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	Counter(ctx context.Context, in *CounterRequest, opts ...grpc.CallOption) (*CounterResponse, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) Counter(ctx context.Context, in *CounterRequest, opts ...grpc.CallOption) (*CounterResponse, error) {
	out := new(CounterResponse)
	err := c.cc.Invoke(ctx, CounterService_Counter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations must embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	Counter(context.Context, *CounterRequest) (*CounterResponse, error)
	mustEmbedUnimplementedCounterServiceServer()
}

// UnimplementedCounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) Counter(context.Context, *CounterRequest) (*CounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Counter not implemented")
}
func (UnimplementedCounterServiceServer) mustEmbedUnimplementedCounterServiceServer() {}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_Counter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).Counter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CounterService_Counter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).Counter(ctx, req.(*CounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clicker.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Counter",
			Handler:    _CounterService_Counter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
