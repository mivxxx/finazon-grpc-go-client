// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: time_series.proto

package finazon_grpc_go_client

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
	TimeSeriesService_GetTimeSeries_FullMethodName         = "/finazon.TimeSeriesService/GetTimeSeries"
	TimeSeriesService_GetTimeSeriesAtr_FullMethodName      = "/finazon.TimeSeriesService/GetTimeSeriesAtr"
	TimeSeriesService_GetTimeSeriesBBands_FullMethodName   = "/finazon.TimeSeriesService/GetTimeSeriesBBands"
	TimeSeriesService_GetTimeSeriesIchimoku_FullMethodName = "/finazon.TimeSeriesService/GetTimeSeriesIchimoku"
	TimeSeriesService_GetTimeSeriesMa_FullMethodName       = "/finazon.TimeSeriesService/GetTimeSeriesMa"
	TimeSeriesService_GetTimeSeriesMacd_FullMethodName     = "/finazon.TimeSeriesService/GetTimeSeriesMacd"
	TimeSeriesService_GetTimeSeriesObv_FullMethodName      = "/finazon.TimeSeriesService/GetTimeSeriesObv"
	TimeSeriesService_GetTimeSeriesRsi_FullMethodName      = "/finazon.TimeSeriesService/GetTimeSeriesRsi"
	TimeSeriesService_GetTimeSeriesSar_FullMethodName      = "/finazon.TimeSeriesService/GetTimeSeriesSar"
	TimeSeriesService_GetTimeSeriesStoch_FullMethodName    = "/finazon.TimeSeriesService/GetTimeSeriesStoch"
)

// TimeSeriesServiceClient is the client API for TimeSeriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeSeriesServiceClient interface {
	// Get time series data without technical indicators
	GetTimeSeries(ctx context.Context, in *GetTimeSeriesRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error)
	// Get time series data for ATR technical indicator.
	GetTimeSeriesAtr(ctx context.Context, in *GetTimeSeriesAtrRequest, opts ...grpc.CallOption) (*GetTimeSeriesAtrResponse, error)
	// Get time series data for BBands technical indicator.
	GetTimeSeriesBBands(ctx context.Context, in *GetTimeSeriesBBandsRequest, opts ...grpc.CallOption) (*GetTimeSeriesBBandsResponse, error)
	// Get time series data for Ichimoku technical indicator.
	GetTimeSeriesIchimoku(ctx context.Context, in *GetTimeSeriesIchimokuRequest, opts ...grpc.CallOption) (*GetTimeSeriesIchimokuResponse, error)
	// Get time series data for Ma technical indicator.
	GetTimeSeriesMa(ctx context.Context, in *GetTimeSeriesMaRequest, opts ...grpc.CallOption) (*GetTimeSeriesMaResponse, error)
	// Get time series data for Macd technical indicator.
	GetTimeSeriesMacd(ctx context.Context, in *GetTimeSeriesMacdRequest, opts ...grpc.CallOption) (*GetTimeSeriesMacdResponse, error)
	// Get time series data for Obv technical indicator.
	GetTimeSeriesObv(ctx context.Context, in *GetTimeSeriesObvRequest, opts ...grpc.CallOption) (*GetTimeSeriesObvResponse, error)
	// Get time series data for Rsi technical indicator.
	GetTimeSeriesRsi(ctx context.Context, in *GetTimeSeriesRsiRequest, opts ...grpc.CallOption) (*GetTimeSeriesRsiResponse, error)
	// Get time series data for Sar technical indicator.
	GetTimeSeriesSar(ctx context.Context, in *GetTimeSeriesSarRequest, opts ...grpc.CallOption) (*GetTimeSeriesSarResponse, error)
	// Get time series data for Stoch technical indicator.
	GetTimeSeriesStoch(ctx context.Context, in *GetTimeSeriesStochRequest, opts ...grpc.CallOption) (*GetTimeSeriesStochResponse, error)
}

type timeSeriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSeriesServiceClient(cc grpc.ClientConnInterface) TimeSeriesServiceClient {
	return &timeSeriesServiceClient{cc}
}

func (c *timeSeriesServiceClient) GetTimeSeries(ctx context.Context, in *GetTimeSeriesRequest, opts ...grpc.CallOption) (*GetTimeSeriesResponse, error) {
	out := new(GetTimeSeriesResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesAtr(ctx context.Context, in *GetTimeSeriesAtrRequest, opts ...grpc.CallOption) (*GetTimeSeriesAtrResponse, error) {
	out := new(GetTimeSeriesAtrResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesAtr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesBBands(ctx context.Context, in *GetTimeSeriesBBandsRequest, opts ...grpc.CallOption) (*GetTimeSeriesBBandsResponse, error) {
	out := new(GetTimeSeriesBBandsResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesBBands_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesIchimoku(ctx context.Context, in *GetTimeSeriesIchimokuRequest, opts ...grpc.CallOption) (*GetTimeSeriesIchimokuResponse, error) {
	out := new(GetTimeSeriesIchimokuResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesIchimoku_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesMa(ctx context.Context, in *GetTimeSeriesMaRequest, opts ...grpc.CallOption) (*GetTimeSeriesMaResponse, error) {
	out := new(GetTimeSeriesMaResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesMa_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesMacd(ctx context.Context, in *GetTimeSeriesMacdRequest, opts ...grpc.CallOption) (*GetTimeSeriesMacdResponse, error) {
	out := new(GetTimeSeriesMacdResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesMacd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesObv(ctx context.Context, in *GetTimeSeriesObvRequest, opts ...grpc.CallOption) (*GetTimeSeriesObvResponse, error) {
	out := new(GetTimeSeriesObvResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesObv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesRsi(ctx context.Context, in *GetTimeSeriesRsiRequest, opts ...grpc.CallOption) (*GetTimeSeriesRsiResponse, error) {
	out := new(GetTimeSeriesRsiResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesRsi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesSar(ctx context.Context, in *GetTimeSeriesSarRequest, opts ...grpc.CallOption) (*GetTimeSeriesSarResponse, error) {
	out := new(GetTimeSeriesSarResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesSar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSeriesServiceClient) GetTimeSeriesStoch(ctx context.Context, in *GetTimeSeriesStochRequest, opts ...grpc.CallOption) (*GetTimeSeriesStochResponse, error) {
	out := new(GetTimeSeriesStochResponse)
	err := c.cc.Invoke(ctx, TimeSeriesService_GetTimeSeriesStoch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSeriesServiceServer is the server API for TimeSeriesService service.
// All implementations must embed UnimplementedTimeSeriesServiceServer
// for forward compatibility
type TimeSeriesServiceServer interface {
	// Get time series data without technical indicators
	GetTimeSeries(context.Context, *GetTimeSeriesRequest) (*GetTimeSeriesResponse, error)
	// Get time series data for ATR technical indicator.
	GetTimeSeriesAtr(context.Context, *GetTimeSeriesAtrRequest) (*GetTimeSeriesAtrResponse, error)
	// Get time series data for BBands technical indicator.
	GetTimeSeriesBBands(context.Context, *GetTimeSeriesBBandsRequest) (*GetTimeSeriesBBandsResponse, error)
	// Get time series data for Ichimoku technical indicator.
	GetTimeSeriesIchimoku(context.Context, *GetTimeSeriesIchimokuRequest) (*GetTimeSeriesIchimokuResponse, error)
	// Get time series data for Ma technical indicator.
	GetTimeSeriesMa(context.Context, *GetTimeSeriesMaRequest) (*GetTimeSeriesMaResponse, error)
	// Get time series data for Macd technical indicator.
	GetTimeSeriesMacd(context.Context, *GetTimeSeriesMacdRequest) (*GetTimeSeriesMacdResponse, error)
	// Get time series data for Obv technical indicator.
	GetTimeSeriesObv(context.Context, *GetTimeSeriesObvRequest) (*GetTimeSeriesObvResponse, error)
	// Get time series data for Rsi technical indicator.
	GetTimeSeriesRsi(context.Context, *GetTimeSeriesRsiRequest) (*GetTimeSeriesRsiResponse, error)
	// Get time series data for Sar technical indicator.
	GetTimeSeriesSar(context.Context, *GetTimeSeriesSarRequest) (*GetTimeSeriesSarResponse, error)
	// Get time series data for Stoch technical indicator.
	GetTimeSeriesStoch(context.Context, *GetTimeSeriesStochRequest) (*GetTimeSeriesStochResponse, error)
	mustEmbedUnimplementedTimeSeriesServiceServer()
}

// UnimplementedTimeSeriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeSeriesServiceServer struct {
}

func (UnimplementedTimeSeriesServiceServer) GetTimeSeries(context.Context, *GetTimeSeriesRequest) (*GetTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeries not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesAtr(context.Context, *GetTimeSeriesAtrRequest) (*GetTimeSeriesAtrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesAtr not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesBBands(context.Context, *GetTimeSeriesBBandsRequest) (*GetTimeSeriesBBandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesBBands not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesIchimoku(context.Context, *GetTimeSeriesIchimokuRequest) (*GetTimeSeriesIchimokuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesIchimoku not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesMa(context.Context, *GetTimeSeriesMaRequest) (*GetTimeSeriesMaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesMa not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesMacd(context.Context, *GetTimeSeriesMacdRequest) (*GetTimeSeriesMacdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesMacd not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesObv(context.Context, *GetTimeSeriesObvRequest) (*GetTimeSeriesObvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesObv not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesRsi(context.Context, *GetTimeSeriesRsiRequest) (*GetTimeSeriesRsiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesRsi not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesSar(context.Context, *GetTimeSeriesSarRequest) (*GetTimeSeriesSarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesSar not implemented")
}
func (UnimplementedTimeSeriesServiceServer) GetTimeSeriesStoch(context.Context, *GetTimeSeriesStochRequest) (*GetTimeSeriesStochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesStoch not implemented")
}
func (UnimplementedTimeSeriesServiceServer) mustEmbedUnimplementedTimeSeriesServiceServer() {}

// UnsafeTimeSeriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeSeriesServiceServer will
// result in compilation errors.
type UnsafeTimeSeriesServiceServer interface {
	mustEmbedUnimplementedTimeSeriesServiceServer()
}

func RegisterTimeSeriesServiceServer(s grpc.ServiceRegistrar, srv TimeSeriesServiceServer) {
	s.RegisterService(&TimeSeriesService_ServiceDesc, srv)
}

func _TimeSeriesService_GetTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeries(ctx, req.(*GetTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesAtr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesAtrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesAtr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesAtr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesAtr(ctx, req.(*GetTimeSeriesAtrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesBBands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesBBandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesBBands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesBBands_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesBBands(ctx, req.(*GetTimeSeriesBBandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesIchimoku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesIchimokuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesIchimoku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesIchimoku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesIchimoku(ctx, req.(*GetTimeSeriesIchimokuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesMa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesMaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesMa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesMa_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesMa(ctx, req.(*GetTimeSeriesMaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesMacd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesMacdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesMacd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesMacd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesMacd(ctx, req.(*GetTimeSeriesMacdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesObv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesObvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesObv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesObv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesObv(ctx, req.(*GetTimeSeriesObvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesRsi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesRsiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesRsi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesRsi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesRsi(ctx, req.(*GetTimeSeriesRsiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesSar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesSarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesSar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesSar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesSar(ctx, req.(*GetTimeSeriesSarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSeriesService_GetTimeSeriesStoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesStochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesStoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeSeriesService_GetTimeSeriesStoch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSeriesServiceServer).GetTimeSeriesStoch(ctx, req.(*GetTimeSeriesStochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeSeriesService_ServiceDesc is the grpc.ServiceDesc for TimeSeriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeSeriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finazon.TimeSeriesService",
	HandlerType: (*TimeSeriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeSeries",
			Handler:    _TimeSeriesService_GetTimeSeries_Handler,
		},
		{
			MethodName: "GetTimeSeriesAtr",
			Handler:    _TimeSeriesService_GetTimeSeriesAtr_Handler,
		},
		{
			MethodName: "GetTimeSeriesBBands",
			Handler:    _TimeSeriesService_GetTimeSeriesBBands_Handler,
		},
		{
			MethodName: "GetTimeSeriesIchimoku",
			Handler:    _TimeSeriesService_GetTimeSeriesIchimoku_Handler,
		},
		{
			MethodName: "GetTimeSeriesMa",
			Handler:    _TimeSeriesService_GetTimeSeriesMa_Handler,
		},
		{
			MethodName: "GetTimeSeriesMacd",
			Handler:    _TimeSeriesService_GetTimeSeriesMacd_Handler,
		},
		{
			MethodName: "GetTimeSeriesObv",
			Handler:    _TimeSeriesService_GetTimeSeriesObv_Handler,
		},
		{
			MethodName: "GetTimeSeriesRsi",
			Handler:    _TimeSeriesService_GetTimeSeriesRsi_Handler,
		},
		{
			MethodName: "GetTimeSeriesSar",
			Handler:    _TimeSeriesService_GetTimeSeriesSar_Handler,
		},
		{
			MethodName: "GetTimeSeriesStoch",
			Handler:    _TimeSeriesService_GetTimeSeriesStoch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time_series.proto",
}
