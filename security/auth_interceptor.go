package security

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type Interceptor struct {
	Auth Auth
}

func (interceptor *Interceptor) RequestHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	// switch info.FullMethod {
	// case "/Fulfillment/RegisterDeliveryPartner":

	// 	r := req.(*pb.RegisterDeliveryPartnerRequest)
	// 	if r.Username == "" || r.Password == "" {
	// 		return nil, status.Error(codes.InvalidArgument, "Username Passwor must be provided")
	// 	}
	// 	if r.Location == nil || r.Location.XCordinate == 0 || r.Location.YCordinate == 0 {
	// 		return nil, status.Error(codes.InvalidArgument, "Location cannot be empty")
	// 	}
	// 	if r.Availability != "Available" && r.Availability != "Unavailable" {
	// 		return nil, status.Error(codes.InvalidArgument, "Enter valid availability status ie. available or unavailable.")
	// 	}

	// case "/Fulfillment/GetNearestDeliveryPartner":

	// 	r := req.(*pb.Location)
	// 	if r == nil || r.XCordinate <= 0 || r.YCordinate <= 0 {
	// 		return nil, status.Error(codes.InvalidArgument, "Provide valid location")
	// 	}

	// case "/Fulfillment/UpdateDeliveryStatus":

	// 	if err := interceptor.Auth.isAuthenticated(ctx); err != nil {
	// 		return nil, err
	// 	}
	// 	r := req.(*pb.DeliveryStatus)
	// 	if r == nil || (r.Status != "Delivered" && r.Status != "Assigned") {
	// 		return nil, status.Error(codes.InvalidArgument, "Provide valid status i.e Delivered or Assigned")
	// 	}
	// }
	resp, err := handler(ctx, req)
	return resp, err

}
