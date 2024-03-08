package models

import "time"

type Delivery struct {
	ID                int64    `gorm:"primaryKey"`
	OrderID           int64    `gorm:"not null"`
	PickUpLocation    Location `gorm:"embedded;embeddedPrefix:pickup_"`
	DeliveryLocation  Location `gorm:"embedded;embeddedPrefix:delivery_"`
	DeliveryPartnerID int64
	DeliveryPartner   DeliveryPartner `gorm:"foreignKey:DeliveryPartnerID"`
	DeliveredAt       time.Time       `gorm:"default:NULL"`
}

func (delivery *Delivery) AssignDeliveryPartner(deliveryPartner DeliveryPartner) {
	delivery.DeliveryPartner = deliveryPartner
}

func (delivery *Delivery) MarkDelivered() {
	delivery.DeliveredAt = time.Now()
}
