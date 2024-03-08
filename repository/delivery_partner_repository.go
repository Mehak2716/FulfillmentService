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

func (repo *DeliveryPartnerRepository) IsExists(username string) bool {
	var count int64
	repo.DB.Model(&models.DeliveryPartner{}).
		Where("delivery_partners.username = ? AND delivery_partners.deleted_at IS NULL", username).
		Count(&count)

	return count > 0
}

func (repo *DeliveryPartnerRepository) FetchNearest(location models.Location) (*models.DeliveryPartner, error) {
	var nearestDeliveryPartner models.DeliveryPartner

	res := repo.DB.Raw(`
    SELECT delivery_partners.*
    FROM delivery_partners
    WHERE delivery_partners.availability = 'available'
    ORDER BY ST_Distance(
        ST_MakePoint(?, ?)::geography,
        ST_MakePoint(delivery_partners.x_cordinate, delivery_partners.y_cordinate)::geography
    ) LIMIT 1
`, location.XCordinate, location.YCordinate).Scan(&nearestDeliveryPartner)

	if res.Error != nil {
		return nil, res.Error
	}

	return &nearestDeliveryPartner, nil

}

func (repo *DeliveryPartnerRepository) Fetch(ID int64) (*models.DeliveryPartner, error) {

	var deliveryPartner models.DeliveryPartner
	res := repo.DB.Where("id = ?", ID).First(&deliveryPartner)

	if res.Error != nil {
		return nil, res.Error
	}
	return &deliveryPartner, nil
}

func (repo *DeliveryPartnerRepository) Update(deliveryPartner *models.DeliveryPartner) error {

	res := repo.DB.Save(deliveryPartner)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
