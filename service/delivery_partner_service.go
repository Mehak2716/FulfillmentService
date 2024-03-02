package service

import (
	"fulfillment/models"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryPartnerService struct {
	Repo repository.DeliveryPartnerRepository
}

func (service *DeliveryPartnerService) Register(req *pb.RegisterDeliveryPartnerRequest) (*pb.DeliveryPartnerResponse, error) {

	deliveryPartner := models.DeliveryPartner{
		Username: req.Username,
		Password: req.Password,
		Location: models.Location{
			XCordinate: float64(req.Location.XCordinate),
			YCordinate: float64(req.Location.YCordinate),
		}}

	registeredDeliveryPartner, err := service.Repo.Save(&deliveryPartner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create menu item")
	}

	response := &pb.DeliveryPartnerResponse{
		Id:       int64(registeredDeliveryPartner.ID),
		Username: registeredDeliveryPartner.Username,
		Location: &pb.Location{
			XCordinate: float32(registeredDeliveryPartner.Location.XCordinate),
			YCordinate: float32(registeredDeliveryPartner.Location.YCordinate),
		}}
	return response, nil
}
