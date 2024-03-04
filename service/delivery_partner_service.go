package service

import (
	"fulfillment/mapper"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryPartnerService struct {
	Repo repository.DeliveryPartnerRepository
}

func (service *DeliveryPartnerService) Register(req *pb.RegisterDeliveryPartnerRequest) (*pb.DeliveryPartnerResponse, error) {

	deliveryPartner := mapper.MapToDeliveryPartner(req)

	if service.Repo.IsExists(deliveryPartner.Username) {
		return nil, status.Errorf(codes.AlreadyExists, "Try with different username")
	}

	registeredDeliveryPartner, err := service.Repo.Save(&deliveryPartner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to register delivery partner")
	}

	response := mapper.MapToDeliveryPartnerResponse(*registeredDeliveryPartner)
	return response, nil
}
