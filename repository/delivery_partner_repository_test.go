package repository

import (
	"fulfillment/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDeliveryPartnerTest() (sqlmock.Sqlmock, *DeliveryPartnerRepository) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := DeliveryPartnerRepository{DB: gormDB}

	return mock, &repo
}

func TestDeliveryPartnerCreatedSuccessfully(t *testing.T) {
	mock, deliveryPartnerRepo := setUpDeliveryPartnerTest()
	deliveryPartner := models.DeliveryPartner{
		Username: "testName",
		Password: "abc",
		Location: models.Location{},
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "testName")
	mock.ExpectQuery("INSERT INTO \"delivery_partners\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := deliveryPartnerRepo.Save(&deliveryPartner)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if res.ID != 1 || res.Username != "testName" {
		t.Fatal("Unexpected Result")
	}
}
