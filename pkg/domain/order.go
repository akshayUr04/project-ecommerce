package domain

import (
	"time"
)

type PaymentType struct {
	Id        uint   `gorm:"primaryKey;unique;not null"`
	Type      string `gorm:"unique;not null"`
	IsDefault bool
}

type Orderes struct {
	Id              uint `gorm:"primaryKey;unique;not null"`
	UserId          uint
	OrderDdate      time.Time
	PaymentTypeId   uint
	PaymentType     PaymentType
	ShippingAddress uint
	OrderTotal      int
	OrderStatus     string
}

type OrderItem struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	OrderId       uint
	Orderes       Orderes `gorm:"foreignKey:OrderId"`
	ProductItemId uint
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemId"`
	Quantity      int
	Price         int
}
