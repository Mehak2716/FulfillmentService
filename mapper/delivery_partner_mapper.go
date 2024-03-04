package mapper

import (
	"fulfillment/models"
	pb "fulfillment/proto/fulfillment"
)

func MapToDeliveryPartner(req *pb.RegisterDeliveryPartnerRequest) models.DeliveryPartner {

	return models.DeliveryPartner{
		Username: req.Username,
		Password: req.Password,
		Location: models.Location{
			XCordinate: float64(req.Location.XCordinate),
			YCordinate: float64(req.Location.YCordinate),
		},
		Availability: req.Availability,
	}

}

func MapToDeliveryPartnerResponse(deliveryPartner models.DeliveryPartner) *pb.DeliveryPartnerResponse {

	return &pb.DeliveryPartnerResponse{
		Id:       int64(deliveryPartner.ID),
		Username: deliveryPartner.Username,
		Location: &pb.Location{
			XCordinate: float32(deliveryPartner.Location.XCordinate),
			YCordinate: float32(deliveryPartner.Location.YCordinate),
		},
		Availability: deliveryPartner.Availability,
	}
}

func MapToNearestDeliveryPartnerResponse(deliveryPartner models.DeliveryPartner) *pb.NearestDeliveryPartnerResponse {

	return &pb.NearestDeliveryPartnerResponse{
		Id:       int64(deliveryPartner.ID),
		Username: deliveryPartner.Username,
	}
}
