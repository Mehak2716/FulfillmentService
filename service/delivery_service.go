package service

import (
	"fulfillment/mapper"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryService struct {
	Repo                   repository.DeliveryRepository
	DeliveryPartnerService DeliveryPartnerService
}

func (service *DeliveryService) Initiate(req *pb.DeliveryRequest) (*pb.Response, error) {

	delivery := mapper.MapToDelivery(req)
	deliveryPartner, err := service.DeliveryPartnerService.GetNearest(delivery.PickUpLocation)
	if err != nil {
		return nil, err
	}
	delivery.AssignDeliveryPartner(*deliveryPartner)
	error := service.Repo.Save(&delivery)
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Failed to initiate delivery")
	}

	res := mapper.MapToResponse()
	return res, nil
}

func (service *DeliveryService) UpdateStatus(req *pb.DeliveryStatusRequest) (*pb.Response, error) {

	return nil, nil
}
