// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: fulfillment.proto

package fulfillment

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

// DeliveryPartnerServiceClient is the client API for DeliveryPartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryPartnerServiceClient interface {
	RegisterDeliveryPartner(ctx context.Context, in *RegisterDeliveryPartnerRequest, opts ...grpc.CallOption) (*DeliveryPartner, error)
}

type deliveryPartnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryPartnerServiceClient(cc grpc.ClientConnInterface) DeliveryPartnerServiceClient {
	return &deliveryPartnerServiceClient{cc}
}

func (c *deliveryPartnerServiceClient) RegisterDeliveryPartner(ctx context.Context, in *RegisterDeliveryPartnerRequest, opts ...grpc.CallOption) (*DeliveryPartner, error) {
	out := new(DeliveryPartner)
	err := c.cc.Invoke(ctx, "/DeliveryPartnerService/RegisterDeliveryPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryPartnerServiceServer is the server API for DeliveryPartnerService service.
// All implementations must embed UnimplementedDeliveryPartnerServiceServer
// for forward compatibility
type DeliveryPartnerServiceServer interface {
	RegisterDeliveryPartner(context.Context, *RegisterDeliveryPartnerRequest) (*DeliveryPartner, error)
	mustEmbedUnimplementedDeliveryPartnerServiceServer()
}

// UnimplementedDeliveryPartnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryPartnerServiceServer struct {
}

func (UnimplementedDeliveryPartnerServiceServer) RegisterDeliveryPartner(context.Context, *RegisterDeliveryPartnerRequest) (*DeliveryPartner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDeliveryPartner not implemented")
}
func (UnimplementedDeliveryPartnerServiceServer) mustEmbedUnimplementedDeliveryPartnerServiceServer() {
}

// UnsafeDeliveryPartnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryPartnerServiceServer will
// result in compilation errors.
type UnsafeDeliveryPartnerServiceServer interface {
	mustEmbedUnimplementedDeliveryPartnerServiceServer()
}

func RegisterDeliveryPartnerServiceServer(s grpc.ServiceRegistrar, srv DeliveryPartnerServiceServer) {
	s.RegisterService(&DeliveryPartnerService_ServiceDesc, srv)
}

func _DeliveryPartnerService_RegisterDeliveryPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeliveryPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryPartnerServiceServer).RegisterDeliveryPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeliveryPartnerService/RegisterDeliveryPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryPartnerServiceServer).RegisterDeliveryPartner(ctx, req.(*RegisterDeliveryPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryPartnerService_ServiceDesc is the grpc.ServiceDesc for DeliveryPartnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryPartnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeliveryPartnerService",
	HandlerType: (*DeliveryPartnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDeliveryPartner",
			Handler:    _DeliveryPartnerService_RegisterDeliveryPartner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulfillment.proto",
}

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	InitiateDelivery(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*Response, error)
	MarkDelivered(ctx context.Context, in *DeliveredRequest, opts ...grpc.CallOption) (*DeliveredResponse, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) InitiateDelivery(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/DeliveryService/InitiateDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) MarkDelivered(ctx context.Context, in *DeliveredRequest, opts ...grpc.CallOption) (*DeliveredResponse, error) {
	out := new(DeliveredResponse)
	err := c.cc.Invoke(ctx, "/DeliveryService/MarkDelivered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
// All implementations must embed UnimplementedDeliveryServiceServer
// for forward compatibility
type DeliveryServiceServer interface {
	InitiateDelivery(context.Context, *DeliveryRequest) (*Response, error)
	MarkDelivered(context.Context, *DeliveredRequest) (*DeliveredResponse, error)
	mustEmbedUnimplementedDeliveryServiceServer()
}

// UnimplementedDeliveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeliveryServiceServer struct {
}

func (UnimplementedDeliveryServiceServer) InitiateDelivery(context.Context, *DeliveryRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) MarkDelivered(context.Context, *DeliveredRequest) (*DeliveredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkDelivered not implemented")
}
func (UnimplementedDeliveryServiceServer) mustEmbedUnimplementedDeliveryServiceServer() {}

// UnsafeDeliveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryServiceServer will
// result in compilation errors.
type UnsafeDeliveryServiceServer interface {
	mustEmbedUnimplementedDeliveryServiceServer()
}

func RegisterDeliveryServiceServer(s grpc.ServiceRegistrar, srv DeliveryServiceServer) {
	s.RegisterService(&DeliveryService_ServiceDesc, srv)
}

func _DeliveryService_InitiateDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).InitiateDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeliveryService/InitiateDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).InitiateDelivery(ctx, req.(*DeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_MarkDelivered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).MarkDelivered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeliveryService/MarkDelivered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).MarkDelivered(ctx, req.(*DeliveredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService_ServiceDesc is the grpc.ServiceDesc for DeliveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitiateDelivery",
			Handler:    _DeliveryService_InitiateDelivery_Handler,
		},
		{
			MethodName: "MarkDelivered",
			Handler:    _DeliveryService_MarkDelivered_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulfillment.proto",
}
