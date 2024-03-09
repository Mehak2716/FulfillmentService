package repository

import (
	"fmt"
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

func TestIsExistsForExistingDeliveryPartnerSuccessfully(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	username := "testUsername"

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(username)

	if !result {
		t.Fatalf("Expected IsExists to return true, but got false")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestIsExistsForNonExistingDeliveryPartnerSuccessfully(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	username := "testUsername"

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(username)

	if result {
		t.Fatalf("Expected IsExists to return false, but got true")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchNearestDeliveryPartner(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
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

	result, err := repo.FetchNearest(location)

	fmt.Print(result.ID)
	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if result.ID != 1 {
		t.Fatal("Unexpected Result: The nearest partner should have ID 1")
	}
}

func TestFetchNearestDeliveryPartnerWithNoResult(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	location := models.Location{
		XCordinate: 40.0,
		YCordinate: 30.0,
	}

	mock.ExpectQuery("SELECT delivery_partners.* FROM delivery_partners").
		WithArgs(location.XCordinate, location.YCordinate).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}))

	result, err := repo.FetchNearest(location)

	fmt.Print(result)
	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if result.ID != 0 {
		t.Fatal("Expected no result, but got a non nil result")
	}
}

func TestFetchDeliveryPartnerByID(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	partnerID := int64(1)

	mock.ExpectQuery("SELECT \\* FROM \"delivery_partners\" WHERE id = \\$1 AND \"delivery_partners\".\"deleted_at\" IS NULL ORDER BY \"delivery_partners\".\"id\" LIMIT \\$2").
		WithArgs(partnerID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "x_cordinate", "y_cordinate", "availability"}).
			AddRow(1, "testName1", 40.2, 30.2, "available"))

	result, err := repo.Fetch(partnerID)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if int64(result.ID) != partnerID {
		t.Fatalf("Unexpected Result: The fetched partner should have ID %d", partnerID)
	}
}
