package main

import (
	"fulfillment/config"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"
	"fulfillment/security"
	"fulfillment/server"
	"fulfillment/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const servePort string = ":9000"

func main() {
	lis, err := net.Listen("tcp", servePort)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	log.Printf("Listening on %s\n", servePort)
	db := config.DatabaseConnection()
	deliveryPartnerRepo := repository.DeliveryPartnerRepository{DB: db}
	deliveryPartnerService := service.DeliveryPartnerService{Repo: deliveryPartnerRepo}

	interceptor := &security.Interceptor{
		Auth: security.Auth{Repo: deliveryPartnerRepo},
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.RequestHandler),
	)

	server := server.NewFulfillmentServer(deliveryPartnerService)
	pb.RegisterFulfillmentServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
