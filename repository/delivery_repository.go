package repository

import (
	"fulfillment/models"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	DB *gorm.DB
}

func (repo *DeliveryRepository) Save(delivery *models.Delivery) error {

	return nil
}
