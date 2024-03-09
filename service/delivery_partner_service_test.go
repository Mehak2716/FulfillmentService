package service

import (
	"fulfillment/models"
	pb "fulfillment/proto/fulfillment"
	"fulfillment/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDeliveryPartnerServiceTest() (sqlmock.Sqlmock, *DeliveryPartnerService) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := repository.DeliveryPartnerRepository{DB: gormDB}

	service := DeliveryPartnerService{repo}
	return mock, &service
}

func TestRegisterDuplicateDeliveryPartnerExpectAlreadyExistError(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	req := &pb.RegisterDeliveryPartnerRequest{
		Username: "testUsername",
		Password: "testPassword",
		Location: &pb.Location{},
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	res, err := service.Register(req)

	if res != nil {
		t.Fatalf("Expected response to be nil")
	}
	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.AlreadyExists
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	} else {
		t.Fatal("Expected an error, but got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestRegisterDeliveryPartnerSuccessfully(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	req := &pb.RegisterDeliveryPartnerRequest{
		Username: "testUsername",
		Password: "testPassword",
		Location: &pb.Location{},
	}

	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rowsCount)
	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "testUsername", "testPassword")
	mock.ExpectQuery("INSERT INTO \"delivery_partners\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := service.Register(req)

	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got %", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetNearestDeliveryPartner(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	location := models.Location{
		XCordinate: 40.0,
		YCordinate: 30.0,
	}

	mock.ExpectQuery("SELECT delivery_partners.* FROM delivery_partners").
		WithArgs(location.XCordinate, location.YCordinate).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
			AddRow(1, "testName1", 40.05, 30.05, "available").
			AddRow(2, "testName2", 40.2, 30.2, "available").
			AddRow(3, "testName3", 40.3, 30.3, "available"))

	nearestDeliveryPartner, err := service.GetNearest(location)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if nearestDeliveryPartner.ID != 1 {
		t.Fatalf("Unexpected Result: The nearest partner should have ID 1")
	}
}

func TestGetNearestDeliveryPartnerNotFound(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	location := models.Location{
		XCordinate: 40.0,
		YCordinate: 30.0,
	}

	mock.ExpectQuery("SELECT delivery_partners.* FROM delivery_partners").
		WithArgs(location.XCordinate, location.YCordinate).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}))

	_, err := service.GetNearest(location)

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
func TestUpdateDeliveryPartnerAvailabilitySuccessfully(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	id := int64(1)
	availability := "unavailable"

	mock.ExpectQuery("SELECT \\* FROM \"delivery_partners\" WHERE id = \\$1 AND \"delivery_partners\".\"deleted_at\" IS NULL ORDER BY \"delivery_partners\".\"id\" LIMIT \\$2").
		WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
			AddRow(id, "testName1", 40.2, 30.2, "available"))

	err := service.UpdateAvailability(id, availability)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUpdateDeliveryPartnerAvailabilityForNotFoundPartner(t *testing.T) {
	mock, service := setUpDeliveryPartnerServiceTest()
	id := int64(1)
	availability := "unavailable"

	mock.ExpectQuery("SELECT \\* FROM \"delivery_partners\" WHERE id = \\$1 AND \"delivery_partners\".\"deleted_at\" IS NULL ORDER BY \"delivery_partners\".\"id\" LIMIT \\$2").
		WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}))

	err := service.UpdateAvailability(id, availability)

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
