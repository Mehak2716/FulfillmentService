package security

import (
	"context"
	"log"

	pb "fulfillment/proto/fulfillment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	Auth Auth
}

func (interceptor *Interceptor) RequestHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	switch info.FullMethod {
	case "/DeliveryPartnerService/RegisterDeliveryPartner":

		r := req.(*pb.RegisterDeliveryPartnerRequest)
		if r.Username == "" || r.Password == "" {
			return nil, status.Error(codes.InvalidArgument, "Username Password must be provided")
		}
		if r.Location == nil || r.Location.XCordinate == 0 || r.Location.YCordinate == 0 {
			return nil, status.Error(codes.InvalidArgument, "Location cannot be empty")
		}
		if r.Availability != "Available" && r.Availability != "Unavailable" {
			return nil, status.Error(codes.InvalidArgument, "Enter valid availability status ie. Available or Unavailable.")
		}

	case "/DeliveryService/InitiateDelivery":

		r := req.(*pb.DeliveryRequest)
		if r.OrderId == 0 || r.PickUpLocation.XCordinate <= 0 || r.PickUpLocation.YCordinate <= 0 || r.DeliveryLocation.XCordinate <= 0 || r.DeliveryLocation.YCordinate <= 0 {
			return nil, status.Error(codes.InvalidArgument, "Valid orderId and location must be provided")
		}

	case "/DeliveryService/MarkDelivered":

		if err := interceptor.Auth.isAuthenticated(ctx); err != nil {
			return nil, err
		}
		r := req.(*pb.DeliveredRequest)
		if r.OrderId == 0 {
			return nil, status.Error(codes.InvalidArgument, "Provide valid orderID")
		}
	}
	resp, err := handler(ctx, req)
	return resp, err

}
