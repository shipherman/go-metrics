// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.2
// source: internal/grpcapi/protometrics/grpcapi.proto

package protometrics

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
	MetricsService_GetGauge_FullMethodName        = "/grpcapi.protometrics.MetricsService/GetGauge"
	MetricsService_UpdateGauge_FullMethodName     = "/grpcapi.protometrics.MetricsService/UpdateGauge"
	MetricsService_GetCounter_FullMethodName      = "/grpcapi.protometrics.MetricsService/GetCounter"
	MetricsService_UpdateCounter_FullMethodName   = "/grpcapi.protometrics.MetricsService/UpdateCounter"
	MetricsService_UpdateJSONBatch_FullMethodName = "/grpcapi.protometrics.MetricsService/UpdateJSONBatch"
)

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	GetGauge(ctx context.Context, in *GetGaugeRequest, opts ...grpc.CallOption) (*GetGaugeResponse, error)
	UpdateGauge(ctx context.Context, in *UpdateGaugeRequest, opts ...grpc.CallOption) (*UpdateGaugeResponse, error)
	GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error)
	UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*UpdateCounterResponse, error)
	UpdateJSONBatch(ctx context.Context, in *JSONRequest, opts ...grpc.CallOption) (*JSONResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetGauge(ctx context.Context, in *GetGaugeRequest, opts ...grpc.CallOption) (*GetGaugeResponse, error) {
	out := new(GetGaugeResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetGauge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateGauge(ctx context.Context, in *UpdateGaugeRequest, opts ...grpc.CallOption) (*UpdateGaugeResponse, error) {
	out := new(UpdateGaugeResponse)
	err := c.cc.Invoke(ctx, MetricsService_UpdateGauge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error) {
	out := new(GetCounterResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetCounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateCounter(ctx context.Context, in *UpdateCounterRequest, opts ...grpc.CallOption) (*UpdateCounterResponse, error) {
	out := new(UpdateCounterResponse)
	err := c.cc.Invoke(ctx, MetricsService_UpdateCounter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UpdateJSONBatch(ctx context.Context, in *JSONRequest, opts ...grpc.CallOption) (*JSONResponse, error) {
	out := new(JSONResponse)
	err := c.cc.Invoke(ctx, MetricsService_UpdateJSONBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	GetGauge(context.Context, *GetGaugeRequest) (*GetGaugeResponse, error)
	UpdateGauge(context.Context, *UpdateGaugeRequest) (*UpdateGaugeResponse, error)
	GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error)
	UpdateCounter(context.Context, *UpdateCounterRequest) (*UpdateCounterResponse, error)
	UpdateJSONBatch(context.Context, *JSONRequest) (*JSONResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) GetGauge(context.Context, *GetGaugeRequest) (*GetGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGauge not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateGauge(context.Context, *UpdateGaugeRequest) (*UpdateGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGauge not implemented")
}
func (UnimplementedMetricsServiceServer) GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounter not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateCounter(context.Context, *UpdateCounterRequest) (*UpdateCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCounter not implemented")
}
func (UnimplementedMetricsServiceServer) UpdateJSONBatch(context.Context, *JSONRequest) (*JSONResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJSONBatch not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_GetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetGauge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetGauge(ctx, req.(*GetGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateGauge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateGauge(ctx, req.(*UpdateGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetCounter(ctx, req.(*GetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateCounter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateCounter(ctx, req.(*UpdateCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UpdateJSONBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UpdateJSONBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UpdateJSONBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UpdateJSONBatch(ctx, req.(*JSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.protometrics.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGauge",
			Handler:    _MetricsService_GetGauge_Handler,
		},
		{
			MethodName: "UpdateGauge",
			Handler:    _MetricsService_UpdateGauge_Handler,
		},
		{
			MethodName: "GetCounter",
			Handler:    _MetricsService_GetCounter_Handler,
		},
		{
			MethodName: "UpdateCounter",
			Handler:    _MetricsService_UpdateCounter_Handler,
		},
		{
			MethodName: "UpdateJSONBatch",
			Handler:    _MetricsService_UpdateJSONBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpcapi/protometrics/grpcapi.proto",
}
