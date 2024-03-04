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

func (repo *DeliveryPartnerRepository) FetchNearest(xCordinate float64, yCordinate float64) (*models.DeliveryPartner, error) {
	var nearestDeliveryPartner models.DeliveryPartner

	res := repo.DB.Raw(`
    SELECT delivery_partners.*
    FROM locations
    JOIN delivery_partners ON delivery_partners.id = locations.user_id
    WHERE delivery_partners.availability = 'available'
    ORDER BY ST_Distance(
        ST_MakePoint(?, ?)::geography,
        ST_MakePoint(locations.x_cordinate, locations.y_cordinate)::geography
    ) LIMIT 1
`, xCordinate, yCordinate).Scan(&nearestDeliveryPartner)

	if res.Error != nil {
		return nil, res.Error
	}

	return &nearestDeliveryPartner, nil

}
