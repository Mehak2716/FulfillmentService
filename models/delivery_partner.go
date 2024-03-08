package models

import "gorm.io/gorm"

type DeliveryPartner struct {
	gorm.Model
	Username     string   `gorm:"unique;not null"`
	Password     string   `gorm:"not null"`
	Location     Location `gorm:"embedded"`
	Availability string   `gorm:"not null"`
}

func (deliveryPartner *DeliveryPartner) UpdateAvailability(availability string) {
	deliveryPartner.Availability = availability
}
