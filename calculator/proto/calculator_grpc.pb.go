// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: calculator.proto

package proto

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
	CalculatorService_Add_FullMethodName        = "/calculator.CalculatorService/Add"
	CalculatorService_Subtract_FullMethodName   = "/calculator.CalculatorService/Subtract"
	CalculatorService_Multiply_FullMethodName   = "/calculator.CalculatorService/Multiply"
	CalculatorService_Divide_FullMethodName     = "/calculator.CalculatorService/Divide"
	CalculatorService_GetPrimes_FullMethodName  = "/calculator.CalculatorService/GetPrimes"
	CalculatorService_GetAverage_FullMethodName = "/calculator.CalculatorService/GetAverage"
	CalculatorService_GetMaximum_FullMethodName = "/calculator.CalculatorService/GetMaximum"
	CalculatorService_SquareRoot_FullMethodName = "/calculator.CalculatorService/SquareRoot"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error)
	GetPrimes(ctx context.Context, in *GetPrimesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPrimesResponse], error)
	GetAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[GetAverageRequest, GetAverageResponse], error)
	GetMaximum(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetMaximumRequest, GetMaximumResponse], error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubtractResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Subtract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Multiply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DivideResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Divide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) GetPrimes(ctx context.Context, in *GetPrimesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetPrimesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_GetPrimes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetPrimesRequest, GetPrimesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetPrimesClient = grpc.ServerStreamingClient[GetPrimesResponse]

func (c *calculatorServiceClient) GetAverage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[GetAverageRequest, GetAverageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], CalculatorService_GetAverage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetAverageRequest, GetAverageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetAverageClient = grpc.ClientStreamingClient[GetAverageRequest, GetAverageResponse]

func (c *calculatorServiceClient) GetMaximum(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[GetMaximumRequest, GetMaximumResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], CalculatorService_GetMaximum_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetMaximumRequest, GetMaximumResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetMaximumClient = grpc.BidiStreamingClient[GetMaximumRequest, GetMaximumResponse]

func (c *calculatorServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, CalculatorService_SquareRoot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility.
type CalculatorServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	Divide(context.Context, *DivideRequest) (*DivideResponse, error)
	GetPrimes(*GetPrimesRequest, grpc.ServerStreamingServer[GetPrimesResponse]) error
	GetAverage(grpc.ClientStreamingServer[GetAverageRequest, GetAverageResponse]) error
	GetMaximum(grpc.BidiStreamingServer[GetMaximumRequest, GetMaximumResponse]) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalculatorServiceServer struct{}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServiceServer) Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorServiceServer) Divide(context.Context, *DivideRequest) (*DivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedCalculatorServiceServer) GetPrimes(*GetPrimesRequest, grpc.ServerStreamingServer[GetPrimesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetPrimes not implemented")
}
func (UnimplementedCalculatorServiceServer) GetAverage(grpc.ClientStreamingServer[GetAverageRequest, GetAverageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) GetMaximum(grpc.BidiStreamingServer[GetMaximumRequest, GetMaximumResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetMaximum not implemented")
}
func (UnimplementedCalculatorServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}
func (UnimplementedCalculatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCalculatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_GetPrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPrimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).GetPrimes(m, &grpc.GenericServerStream[GetPrimesRequest, GetPrimesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetPrimesServer = grpc.ServerStreamingServer[GetPrimesResponse]

func _CalculatorService_GetAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).GetAverage(&grpc.GenericServerStream[GetAverageRequest, GetAverageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetAverageServer = grpc.ClientStreamingServer[GetAverageRequest, GetAverageResponse]

func _CalculatorService_GetMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).GetMaximum(&grpc.GenericServerStream[GetMaximumRequest, GetMaximumResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalculatorService_GetMaximumServer = grpc.BidiStreamingServer[GetMaximumRequest, GetMaximumResponse]

func _CalculatorService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_SquareRoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalculatorService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _CalculatorService_Divide_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculatorService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrimes",
			Handler:       _CalculatorService_GetPrimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAverage",
			Handler:       _CalculatorService_GetAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMaximum",
			Handler:       _CalculatorService_GetMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
