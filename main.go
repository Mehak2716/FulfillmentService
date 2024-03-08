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
	"os"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

const servePort string = ":9000"

func main() {
	err := godotenv.Load(".env")
	servePort := os.Getenv("SERVER_PORT")

	if err != nil {
		log.Fatalf("Error loading env file %v", err)
	}

	lis, err := net.Listen("tcp", servePort)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	log.Printf("Listening on %s\n", servePort)
	db := config.DatabaseConnection()
	deliveryPartnerRepo := repository.DeliveryPartnerRepository{DB: db}
	deliveryPartnerService := service.DeliveryPartnerService{Repo: deliveryPartnerRepo}

	deliveryRepo := repository.DeliveryRepository{DB: db}
	deliveryService := service.DeliveryService{
		Repo:                   deliveryRepo,
		DeliveryPartnerService: deliveryPartnerService,
	}

	interceptor := &security.Interceptor{
		Auth: security.Auth{Repo: deliveryPartnerRepo},
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.RequestHandler),
	)

	deliveryPartnerServiceServer := server.NewDeliveryPartnerServiceServer(deliveryPartnerService)
	deliveryServiceServer := server.NewDeliveryServiceServer(deliveryService)
	pb.RegisterDeliveryPartnerServiceServer(grpcServer, deliveryPartnerServiceServer)
	pb.RegisterDeliveryServiceServer(grpcServer, deliveryServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
