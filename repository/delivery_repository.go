package repository

import (
	"fulfillment/models"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	DB *gorm.DB
}

func (repo *DeliveryRepository) Save(delivery *models.Delivery) error {
	res := repo.DB.Create(delivery)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
