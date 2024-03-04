package validation

import (
	"context"
	"log"

	pb "fulfillment/proto/fulfillment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RequestValidationHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	switch r := req.(type) {
	case *pb.RegisterDeliveryPartnerRequest:
		if r.Username == "" {
			return nil, status.Error(codes.InvalidArgument, "Username cannot be empty")
		}
		if r.Password == "" {
			return nil, status.Error(codes.InvalidArgument, "Password cannot be empty")
		}
		if r.Location == nil || r.Location.XCordinate == 0 || r.Location.YCordinate == 0 {
			return nil, status.Error(codes.InvalidArgument, "Location cannot be empty")
		}

		if r.Availability != "available" && r.Availability != "unavailable" {
			return nil, status.Error(codes.InvalidArgument, "Enter valid availability status ie. available or unavailable.")
		}
	}
	resp, err := handler(ctx, req)
	return resp, err
}
