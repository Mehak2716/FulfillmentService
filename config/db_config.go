package config

import (
	"fmt"
	"log"

	"fulfillment/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", "localhost", "5432", "postgres", "fulfillment", "postgres")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	log.Println("Database connection successful...")
	db.AutoMigrate(&models.DeliveryPartner{}, &models.Delivery{})
	return db
}
