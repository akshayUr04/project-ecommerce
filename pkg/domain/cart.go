package domain

type Cart struct {
	Id      uint `gorm:"primaryKey;unique;not null"`
	User_id uint
	Users   Users `gorm:"foreignKey:User_id"`
	Tottal  int
}

type CartItems struct {
	Id             uint `gorm:"primaryKey;unique;not null"`
	Cart_id        uint
	Cart           Cart        `gorm:"foreignKey:Cart_id"`
	ProductItem_id uint        `gorm:"unique;not null"`
	ProductItem    ProductItem `gorm:"foreignKey:ProductItem_id"`
	Quantity       int
}
