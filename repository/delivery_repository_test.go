package repository

import (
	"fulfillment/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDeliveryTest() (sqlmock.Sqlmock, *DeliveryRepository) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := DeliveryRepository{DB: gormDB}

	return mock, &repo
}

func TestSaveDeliverySuccessfully(t *testing.T) {
	mock, deliveryRepo := setUpDeliveryTest()
	orderID := int64(123)
	delivery := models.Delivery{
		OrderID:          orderID,
		PickUpLocation:   models.Location{XCordinate: 40.2, YCordinate: 30.2},
		DeliveryLocation: models.Location{XCordinate: 10.2, YCordinate: 32.2},
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "order_id", "delivery_partner_id"}).
		AddRow(1, 2, 3)
	mock.ExpectQuery("INSERT INTO \"deliveries\"").WillReturnRows(rows)
	mock.ExpectCommit()
	err := deliveryRepo.Save(&delivery)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchDeliverySuccessfully(t *testing.T) {
	mock, deliveryRepo := setUpDeliveryTest()
	orderID := int64(123)

	mock.ExpectQuery("SELECT (.+) FROM \"deliveries\" WHERE order_id = (.+) ORDER BY \"deliveries\".\"id\" LIMIT (.+)").
		WithArgs(orderID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "delivery_partner_id"}).
			AddRow(1, orderID, 2))

	result, err := deliveryRepo.Fetch(orderID)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if result.OrderID != orderID {
		t.Fatalf("Unexpected Result: The fetched delivery should have OrderID %d", orderID)
	}
}

func TestUpdateDeliverySuccessfully(t *testing.T) {
	mock, deliveryRepo := setUpDeliveryTest()
	orderID := int64(123)
	delivery := models.Delivery{
		ID:                1,
		OrderID:           orderID,
		DeliveryPartnerID: 2,
		PickUpLocation:    models.Location{XCordinate: 40.2, YCordinate: 30.2},
		DeliveryLocation:  models.Location{XCordinate: 10.2, YCordinate: 32.2},
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"deliveries\" SET \"order_id\"=\\$1,\"pickup_x_cordinate\"=\\$2,\"pickup_y_cordinate\"=\\$3,\"delivery_x_cordinate\"=\\$4,\"delivery_y_cordinate\"=\\$5,\"delivery_partner_id\"=\\$6,\"delivered_at\"=\\$7 WHERE \"id\" = \\$8").
		WithArgs(delivery.OrderID, delivery.PickUpLocation.XCordinate, delivery.PickUpLocation.YCordinate, delivery.DeliveryLocation.XCordinate, delivery.DeliveryLocation.YCordinate, delivery.DeliveryPartnerID, delivery.DeliveredAt, delivery.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updatedDelivery, err := deliveryRepo.Update(&delivery)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if updatedDelivery.ID != delivery.ID {
		t.Fatalf("Unexpected Result: The updated delivery should have ID %d", delivery.ID)
	}
}
