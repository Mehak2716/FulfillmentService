package service

import (
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"
)

type DeliveryService struct {
	Repo                   repository.DeliveryRepository
	DeliveryPartnerService DeliveryPartnerService
}

func (service *DeliveryService) Initiate(req *pb.DeliveryRequest) (*pb.Response, error) {

	return nil, nil
}

func (service *DeliveryService) UpdateStatus(req *pb.DeliveryStatusRequest) (*pb.Response, error) {

	return nil, nil
}
