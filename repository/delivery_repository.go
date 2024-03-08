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

func (repo *DeliveryRepository) Fetch(orderID int64) (*models.Delivery, error) {

	var delivery models.Delivery
	res := repo.DB.Where("order_id = ?", orderID).First(&delivery)

	if res.Error != nil {
		return nil, res.Error
	}
	return &delivery, nil
}

func (repo *DeliveryRepository) Update(delivery *models.Delivery) (*models.Delivery, error) {

	res := repo.DB.Save(delivery)

	if res.Error != nil {
		return nil, res.Error
	}

	return delivery, nil
}
