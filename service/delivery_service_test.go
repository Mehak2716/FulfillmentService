package service

import (
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func TestInitiateDelivery(t *testing.T) {
	mock, service := setUpDeliveryServiceTest()

	req := &pb.DeliveryRequest{
		OrderId:          1,
		DeliveryLocation: &pb.Location{XCordinate: 10, YCordinate: 20},
		PickUpLocation:   &pb.Location{XCordinate: 20, YCordinate: 30},
	}

	mock.ExpectQuery("SELECT (.+) FROM delivery_partners").WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
		AddRow(1, "testName1", 40.05, 30.05, "available"))

	mock.ExpectQuery("SELECT \\* FROM \"delivery_partners\" WHERE id = \\$1 AND \"delivery_partners\".\"deleted_at\" IS NULL ORDER BY \"delivery_partners\".\"id\" LIMIT \\$2").
		WithArgs(int64(1), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
			AddRow(1, "testName1", 40.05, 30.05, "available"))
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"delivery_partners\" SET .*").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	service.Initiate(req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInitiateDeliveryForNoAvailablePartner(t *testing.T) {
	mock, service := setUpDeliveryServiceTest()

	req := &pb.DeliveryRequest{
		OrderId:          1,
		DeliveryLocation: &pb.Location{XCordinate: 10, YCordinate: 20},
		PickUpLocation:   &pb.Location{XCordinate: 20, YCordinate: 30},
	}

	mock.ExpectQuery("SELECT (.+) FROM delivery_partners").WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}))

	_, err := service.Initiate(req)

	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.NotFound
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestMarkDeliveredSuccessfully(t *testing.T) {
	mock, service := setUpDeliveryServiceTest()

	orderID := int64(1)
	req := &pb.DeliveredRequest{
		OrderId: orderID,
	}

	mock.ExpectQuery("SELECT (.+) FROM \"deliveries\" WHERE order_id = (.+) ORDER BY \"deliveries\".\"id\" LIMIT (.+)").
		WithArgs(orderID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "delivery_partner_id"}).
			AddRow(1, orderID, 1))

	mock.ExpectQuery("SELECT \\* FROM \"delivery_partners\" WHERE id = \\$1 AND \"delivery_partners\".\"deleted_at\" IS NULL ORDER BY \"delivery_partners\".\"id\" LIMIT \\$2").
		WithArgs(int64(1), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
			AddRow(1, "testName1", 40.05, 30.05, "Unavailable"))

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"delivery_partners\" SET .*").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"deliveries\" SET .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err := service.MarkDelivered(req)

	if err != nil {
		t.Fatalf("Error not expected, got: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestMarkDeliveredForNonExistingOrder(t *testing.T) {
	mock, service := setUpDeliveryServiceTest()

	orderID := int64(1)
	req := &pb.DeliveredRequest{
		OrderId: orderID,
	}

	mock.ExpectQuery("SELECT (.+) FROM \"deliveries\" WHERE order_id = (.+) ORDER BY \"deliveries\".\"id\" LIMIT (.+)").
		WithArgs(orderID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "delivery_partner_id"}))

	_, err := service.MarkDelivered(req)

	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.NotFound
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestMarkDeliveredForAlreadyDeliveredOrder(t *testing.T) {
	mock, service := setUpDeliveryServiceTest()

	orderID := int64(1)
	req := &pb.DeliveredRequest{
		OrderId: orderID,
	}

	mock.ExpectQuery("SELECT (.+) FROM \"deliveries\" WHERE order_id = (.+) ORDER BY \"deliveries\".\"id\" LIMIT (.+)").
		WithArgs(orderID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "order_id", "delivery_partner_id", "delivered_at"}).
			AddRow(1, orderID, 1, time.Now()))

	_, err := service.MarkDelivered(req)

	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.Canceled
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
