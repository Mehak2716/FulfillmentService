package repository

import (
	"fulfillment/models"

	"github.com/jackc/pgx/v5/pgconn"

	"gorm.io/gorm"
)

type DeliveryPartnerRepository struct {
	DB *gorm.DB
}

func (repo *DeliveryPartnerRepository) Save(deliveryPartner *models.DeliveryPartner) (*models.DeliveryPartner, error) {
	res := repo.DB.Create(deliveryPartner)

	if res.Error != nil {
		return nil, res.Error.(*pgconn.PgError)
	}
	return deliveryPartner, nil
}
