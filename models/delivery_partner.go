package models

import "gorm.io/gorm"

type DeliveryPartner struct {
	gorm.Model
	Username     string   `gorm:"unique;not null"`
	Password     string   `gorm:"not null"`
	Location     Location `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Availability string
}
