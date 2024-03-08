package mapper

import (
	"fulfillment/models"
	pb "fulfillment/proto/fulfillment"
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
