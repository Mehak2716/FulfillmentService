package service

import (
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

// func TestGetNearestDeliveryPartnerSuccessfully(t *testing.T) {
// 	mock, service := setUpDeliveryPartnerServiceTest()
// 	req := &pb.Location{
// 		XCordinate: 40,
// 		YCordinate: 30,
// 	}

// 	rows := sqlmock.NewRows([]string{"id", "username", "password", "availability"}).
// 		AddRow(1, "testUsername", "testPassword", "available")
// 	mock.ExpectQuery("SELECT delivery_partners.* FROM locations").
// 		WillReturnRows(rows)
// 	res, err := service.GetNearest(req)

// 	if err != nil {
// 		t.Fatalf("Error not expected but encountered: %v", err)
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("Unfulfilled expectations: %s", err)
// 	}
// 	if res == nil || res.Id != 1 || res.Username != "testUsername" {
// 		t.Fatal("Unexpected Result")
// 	}
// }

// func TestGetNearestDeliveryPartnerForNoResult(t *testing.T) {
// 	mock, service := setUpDeliveryPartnerServiceTest()
// 	req := &pb.Location{
// 		XCordinate: 40,
// 		YCordinate: 30,
// 	}

// 	mock.ExpectQuery("SELECT delivery_partners.* FROM locations").
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "availability"}))
// 	res, err := service.GetNearest(req)

// 	if err == nil {
// 		t.Fatalf("Error expected but not encountered")
// 	}
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("Unfulfilled expectations: %s", err)
// 	}
// 	if res != nil {
// 		t.Fatal("Expected nil result, but got a non-nil result")
// 	}
// }
