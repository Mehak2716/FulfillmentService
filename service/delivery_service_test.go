package service

import (
	"fulfillment/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDeliveryServiceTest() (sqlmock.Sqlmock, *DeliveryService) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	deliveryPartnerRepo := repository.DeliveryPartnerRepository{DB: gormDB}
	deliveryRepo := repository.DeliveryRepository{DB: gormDB}
	deliveryPartnerService := DeliveryPartnerService{Repo: deliveryPartnerRepo}
	service := DeliveryService{DeliveryPartnerService: deliveryPartnerService, Repo: deliveryRepo}
	return mock, &service
}
