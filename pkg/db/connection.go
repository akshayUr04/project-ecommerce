package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/akshayur04/project-ecommerce/pkg/config"
	domain "github.com/akshayur04/project-ecommerce/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Address{},
		&domain.Admins{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductItem{},
		&domain.PaymentType{},
		&domain.Orders{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.Carts{},
		&domain.CartItem{},
		&domain.PaymentDetails{},
		&domain.PaymentStatus{},
		&domain.Coupons{},
	)

	return db, dbErr
}
