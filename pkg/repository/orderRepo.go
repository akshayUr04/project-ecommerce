package repository

import (
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}

func (c *OrderDatabase) PlaceOrder(id int) error {
	tx := c.DB.Begin()

	type cart struct {
		Id     int
		Tottal int
	}
	var carts cart
	query1 := `SELECT id,tottal FROM carts WHERE user_id=? `
	err := tx.Raw(query1, id).Scan(&carts).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var addressId int
	query2 := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(query2, id).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// var orderId int
	query3 := `INSERT INTO orderes (user_id,order_ddate,shipping_address,order_total,order_status)
		VALUES($1,NOW(),$2,$3,'shipped')`
	err = tx.Exec(query3, id, addressId, carts.Tottal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// type cart2 struct {
	// 	Quantity      int
	// 	ProductItemId int
	// }
	// var carts2 []cart2
	// // var quantity int

	// query4 := `SELECT quantity,product_item_id FROM cart_items WHERE cart_id =?`
	// err = tx.Raw(query4, carts.Id).Scan(&carts2).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// query5 := `INSERT INTO order_items (order_id,product_item_id,quantity) VALUES($1,$2,$3)`
	// err = tx.Exec(query5, orderId, carts2).Error

	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
