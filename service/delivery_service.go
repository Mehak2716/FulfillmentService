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
	updateErr := service.DeliveryPartnerService.UpdateAvailability(int64(deliveryPartner.ID), "Unavailable")
	if updateErr != nil {
		return nil, updateErr
	}
	error := service.Repo.Save(&delivery)
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Failed to initiate delivery")
	}
	res := mapper.MapToResponse()
	return res, nil
}

func (service *DeliveryService) MarkDelivered(req *pb.DeliveredRequest) (*pb.DeliveredResponse, error) {

	delivery, err := service.Repo.Fetch(req.OrderId)
	if delivery == nil {
		return nil, status.Errorf(codes.NotFound, "Delivery for Order not found.")
	}
	if !delivery.DeliveredAt.IsZero() {
		return nil, status.Errorf(codes.Canceled, "Order is already delivered.")
	}
	delivery.MarkDelivered()
	updateErr := service.DeliveryPartnerService.UpdateAvailability(int64(delivery.DeliveryPartnerID), "Available")
	if updateErr != nil {
		return nil, updateErr
	}
	service.Repo.Update(delivery)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to mark delivered")
	}
	response := mapper.MapToDeliveredResponse(delivery)
	return response, nil
}
