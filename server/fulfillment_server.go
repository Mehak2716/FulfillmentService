package server

import (
	"context"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/service"
)

type FulfillmentServer struct {
	deliveryPartnerService service.DeliveryPartnerService
	pb.FulfillmentServer
}

func NewFulfillmentServer(deliveryPartnerService service.DeliveryPartnerService) *FulfillmentServer {
	return &FulfillmentServer{
		deliveryPartnerService: deliveryPartnerService,
	}
}

func (server *FulfillmentServer) RegisterDeliveryPartner(ctx context.Context, req *pb.RegisterDeliveryPartnerRequest) (*pb.DeliveryPartnerResponse, error) {
	return server.deliveryPartnerService.Register(req)

}

func (server *FulfillmentServer) GetNearestDeliveryPartner(ctx context.Context, req *pb.Location) (*pb.NearestDeliveryPartnerResponse, error) {

	return server.deliveryPartnerService.GetNearest(req)
}
