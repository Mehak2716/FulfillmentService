package server

import (
	"context"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/service"
)

type DeliveryServiceServer struct {
	deliveryService service.DeliveryService
	pb.DeliveryServiceServer
}

func NewDeliveryServiceServer(deliveryService service.DeliveryService) *DeliveryServiceServer {
	return &DeliveryServiceServer{
		deliveryService: deliveryService,
	}
}

func (server *DeliveryServiceServer) InitiateDelivery(ctx context.Context, req *pb.DeliveryRequest) (*pb.Response, error) {

	return server.deliveryService.Initiate(req)
}

func (server *DeliveryServiceServer) MarkDelivered(ctx context.Context, req *pb.DeliveredRequest) (*pb.DeliveredResponse, error) {

	return server.deliveryService.MarkDelivered(req)
}
