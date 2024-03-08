package mapper

import (
	"fulfillment/models"
	pb "fulfillment/proto/fulfillment"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapToDelivery(req *pb.DeliveryRequest) models.Delivery {

	return models.Delivery{
		OrderID: req.OrderId,
		PickUpLocation: models.Location{
			XCordinate: float64(req.PickUpLocation.XCordinate),
			YCordinate: float64(req.PickUpLocation.YCordinate),
		},
		DeliveryLocation: models.Location{
			XCordinate: float64(req.DeliveryLocation.XCordinate),
			YCordinate: float64(req.DeliveryLocation.YCordinate),
		},
	}

}

func MapToResponse() *pb.Response {

	return &pb.Response{
		Message: "Success",
	}
}

func MapToDeliveredResponse(delivery *models.Delivery) *pb.DeliveredResponse {
	return &pb.DeliveredResponse{
		OrderId:          delivery.OrderID,
		DeliveryParterId: delivery.DeliveryPartnerID,
		DeliveryLocation: &pb.Location{
			XCordinate: float32(delivery.DeliveryLocation.XCordinate),
			YCordinate: float32(delivery.DeliveryLocation.YCordinate),
		},
		DeliveredAt: timestamppb.New(delivery.DeliveredAt),
	}
}
