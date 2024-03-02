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

// FulfillmentClient is the client API for Fulfillment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulfillmentClient interface {
	RegisterDeliveryPartner(ctx context.Context, in *RegisterDeliveryPartnerRequest, opts ...grpc.CallOption) (*DeliveryPartnerResponse, error)
}

type fulfillmentClient struct {
	cc grpc.ClientConnInterface
}

func NewFulfillmentClient(cc grpc.ClientConnInterface) FulfillmentClient {
	return &fulfillmentClient{cc}
}

func (c *fulfillmentClient) RegisterDeliveryPartner(ctx context.Context, in *RegisterDeliveryPartnerRequest, opts ...grpc.CallOption) (*DeliveryPartnerResponse, error) {
	out := new(DeliveryPartnerResponse)
	err := c.cc.Invoke(ctx, "/Fulfillment/RegisterDeliveryPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulfillmentServer is the server API for Fulfillment service.
// All implementations must embed UnimplementedFulfillmentServer
// for forward compatibility
type FulfillmentServer interface {
	RegisterDeliveryPartner(context.Context, *RegisterDeliveryPartnerRequest) (*DeliveryPartnerResponse, error)
	mustEmbedUnimplementedFulfillmentServer()
}

// UnimplementedFulfillmentServer must be embedded to have forward compatible implementations.
type UnimplementedFulfillmentServer struct {
}

func (UnimplementedFulfillmentServer) RegisterDeliveryPartner(context.Context, *RegisterDeliveryPartnerRequest) (*DeliveryPartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDeliveryPartner not implemented")
}
func (UnimplementedFulfillmentServer) mustEmbedUnimplementedFulfillmentServer() {}

// UnsafeFulfillmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulfillmentServer will
// result in compilation errors.
type UnsafeFulfillmentServer interface {
	mustEmbedUnimplementedFulfillmentServer()
}

func RegisterFulfillmentServer(s grpc.ServiceRegistrar, srv FulfillmentServer) {
	s.RegisterService(&Fulfillment_ServiceDesc, srv)
}

func _Fulfillment_RegisterDeliveryPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeliveryPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulfillmentServer).RegisterDeliveryPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fulfillment/RegisterDeliveryPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulfillmentServer).RegisterDeliveryPartner(ctx, req.(*RegisterDeliveryPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fulfillment_ServiceDesc is the grpc.ServiceDesc for Fulfillment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fulfillment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Fulfillment",
	HandlerType: (*FulfillmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDeliveryPartner",
			Handler:    _Fulfillment_RegisterDeliveryPartner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fulfillment.proto",
}
