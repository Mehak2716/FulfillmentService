package utils

import (
	"context"
	"encoding/base64"
	"fulfillment/repository"
	"log"
	"strings"

	pb "fulfillment/proto/fulfillment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Validator struct {
	Repo repository.DeliveryPartnerRepository
}

func (v *Validator) ValidationHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	switch info.FullMethod {
	case "/Fulfillment/RegisterDeliveryPartner":
		r := req.(*pb.RegisterDeliveryPartnerRequest)
		if r.Username == "" || r.Password == "" {
			return nil, status.Error(codes.InvalidArgument, "Username Passwor must be provided")
		}
		if r.Location == nil || r.Location.XCordinate == 0 || r.Location.YCordinate == 0 {
			return nil, status.Error(codes.InvalidArgument, "Location cannot be empty")
		}
		if r.Availability != "available" && r.Availability != "unavailable" {
			return nil, status.Error(codes.InvalidArgument, "Enter valid availability status ie. available or unavailable.")
		}

	case "/Fulfillment/GetNearestDeliveryPartner":
		r := req.(*pb.Location)
		if r == nil || r.XCordinate < 0 || r.YCordinate < 0 {
			return nil, status.Error(codes.InvalidArgument, "Provide valid location")
		}

	case "/Fulfillment/UpdateDeliveryStatus":
		if err := v.isAuthenticated(ctx); err != nil {
			return nil, err
		}
		r := req.(*pb.DeliveryStatus)
		if r == nil || (r.Status != "Delivered" && r.Status != "Assigned") {
			return nil, status.Error(codes.InvalidArgument, "Provide valid status i.e Delivered or Assigned")
		}
	}
	resp, err := handler(ctx, req)
	return resp, err
}

func (v *Validator) isAuthenticated(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "missing credentials")
	}
	authHeaders, ok := md["authorization"]
	if !ok || len(authHeaders) == 0 {
		return status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaderParts := strings.Fields(authHeaders[0])
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
		return status.Error(codes.Unauthenticated, "invalid Authentication header")
	}
	decodedCredentials, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to decode credentials: %v", err)
	}
	credentials := strings.Split(string(decodedCredentials), ":")
	if len(credentials) != 2 {
		return status.Errorf(codes.Unauthenticated, "invalid credentials format")
	}
	username := credentials[0]

	if !v.Repo.IsExists(username) {
		return status.Errorf(codes.Unauthenticated, "Invalid Credentials.User does not exist.")
	}

	return nil
}
