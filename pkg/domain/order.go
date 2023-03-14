package domain

import (
	"time"
)

type PaymentType struct {
	Id   uint   `gorm:"primaryKey;unique;not null"`
	Type string `gorm:"unique;not null"`
}

type Orders struct {
	Id              uint `gorm:"primaryKey;unique;not null"`
	UserId          uint
	OrderDate       time.Time
	PaymentTypeId   uint
	PaymentType     PaymentType `gorm:"foreignKey:PaymentTypeId"`
	ShippingAddress uint
	OrderTotal      int
	OrderStatus     string
}

type OrderItem struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	OrdersId      uint
	Orders        Orders `gorm:"foreignKey:OrdersId"`
	ProductItemId uint
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemId"`
	Quantity      int
	Price         int
}
