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

func TestFetchNearestDeliveryPartnerSuccessfully(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	xCordinate, yCordinate := 40.0, 30.0

	rows := sqlmock.NewRows([]string{"id", "username", "password", "availability"}).
		AddRow(1, "testName", "abc", "available")
	mock.ExpectQuery("SELECT delivery_partners.* FROM locations").
		WillReturnRows(rows)

	result, err := repo.FetchNearest(xCordinate, yCordinate)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if result.ID != 1 || result.Username != "testName" {
		t.Fatal("Unexpected Result")
	}
}

func TestFetchNearestDeliveryPartnerNoResult(t *testing.T) {
	mock, repo := setUpDeliveryPartnerTest()
	xCordinate, yCordinate := 40.0, 30.0

	mock.ExpectQuery("SELECT delivery_partners.* FROM locations").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "availability"}))

	result, err := repo.FetchNearest(xCordinate, yCordinate)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if result == nil {
		t.Fatal("Expected non nil result, but got a nil result")
	}
}
