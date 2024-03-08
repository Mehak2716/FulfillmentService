package server

import (
	"context"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/service"
)

type DeliveryPartnerServiceServer struct {
	deliveryPartnerService service.DeliveryPartnerService
	pb.DeliveryPartnerServiceServer
}

func NewDeliveryPartnerServiceServer(deliveryPartnerService service.DeliveryPartnerService) *DeliveryPartnerServiceServer {
	return &DeliveryPartnerServiceServer{
		deliveryPartnerService: deliveryPartnerService,
	}
}

func (server *DeliveryPartnerServiceServer) RegisterDeliveryPartner(ctx context.Context, req *pb.RegisterDeliveryPartnerRequest) (*pb.DeliveryPartner, error) {
	return server.deliveryPartnerService.Register(req)
}
