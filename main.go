package main

import (
	"fulfillment/config"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"
	"fulfillment/server"
	"fulfillment/service"
	"fulfillment/validation"
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
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(validation.RequestValidationHandler),
	)

	// givenPoint := geom.NewPointFlat(geom.XY, []float64{1.0, 2.0})

	// user := models.DeliveryPartner{
	// 	Username: "koplu",
	// 	Password: "abc",
	// 	Location: models.Location{
	// 		XCordinate: 30,
	// 		YCordinate: 20},
	// }

	// db.Create(&user)

	// var nearestLocation models.Location
	// err = db.Raw("SELECT * FROM locations ORDER BY ST_Distance(ST_MakePoint(?, ?)::geography, ST_MakePoint(locations.x_cordinate, locations.y_cordinate)::geography) LIMIT 1", givenPoint.Coords()[0], givenPoint.Coords()[1]).Scan(&nearestLocation).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(nearestLocation)

	// var locations models.Location

	// // Find locations by UserID
	// if err := db.Where("user_id = ?", 1).Find(&locations).Error; err != nil {
	// }
	// fmt.Print(locations)

	// // userToDelete := models.DeliveryPartner{Model: gorm.Model{ID: 1}}
	// fmt.Print(user.ID)
	// if err := db.Delete(&user).Error; err != nil {
	// 	fmt.Print(err)
	// }

	deliveryPartnerRepo := repository.DeliveryPartnerRepository{DB: db}
	deliveryPartnerService := service.DeliveryPartnerService{Repo: deliveryPartnerRepo}

	server := server.NewFulfillmentServer(deliveryPartnerService)
	pb.RegisterFulfillmentServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
