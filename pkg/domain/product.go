package domain

import "time"

type Category struct {
	Id         uint   `gorm:"primaryKey;unique;not null"`
	Name       string `gorm:"unique;not null"`
	Created_at time.Time
	Updated_at time.Time
}

type Product struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	Name        string `gorm:"unique;not null"`
	Description string
	Brand       string `gorm:"unique;not null"`
	Category_id uint
	Category    Category `gorm:"foreignKey:Category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}

type ProductItem struct {
	Id           uint `gorm:"primaryKey;unique;not null"`
	Product_id   uint
	Product      Product `gorm:"foreignKey:Product_id"`
	Sku          string  `gorm:"unique;not null"`
	Qty_in_stock int
	Color        string
	Ram          int
	Battery      int
	Screen_size  float64
	Storage      int
	Camera       int
	Price        int
	Imag         string
	Created_at   time.Time
	Updated_at   time.Time
}
