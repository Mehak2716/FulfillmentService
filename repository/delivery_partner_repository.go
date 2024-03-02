package repository

import (
	"fulfillment/models"

	"gorm.io/gorm"
)

type DeliveryPartnerRepository struct {
	DB *gorm.DB
}

func (repo *DeliveryPartnerRepository) Save(deliveryPartner *models.DeliveryPartner) (*models.DeliveryPartner, error) {
	res := repo.DB.Create(deliveryPartner)

	if res.Error != nil {
		return nil, res.Error
	}
	return deliveryPartner, nil
}
