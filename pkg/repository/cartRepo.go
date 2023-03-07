package repository

import (
	"fmt"

	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDatabase{DB}
}

func (c *CartDatabase) CreateCart(id int) error {
	query := `INSERT INTO carts (user_id, tottal) VALUES($1,0)`
	err := c.DB.Exec(query, id).Error
	return err
}

func (c *CartDatabase) AddToCart(productId, userId int) error {
	tx := c.DB.Begin()
	fmt.Println(userId)
	var cartId int
	query1 := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(query1, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(1)

	query2 := `INSERT INTO cart_items (cart_id,product_item_id,quantity)VALUES($1,$2,1)
	ON CONFLICT (product_item_id) DO UPDATE SET quantity=cart_items.quantity+1
	WHERE cart_items.cart_id=$1 AND cart_items.product_item_id=$2 `
	err = tx.Exec(query2, cartId, productId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(2)
	var price int
	query3 := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(query3, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(price)

	query4 := `UPDATE carts SET tottal=carts.tottal+$1 WHERE user_id=$2`
	err = tx.Exec(query4, price, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(3)
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) RemoveFromCart(userId, productId int) error {
	tx := c.DB.Begin()
	fmt.Println(userId)
	var cartId int
	query1 := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(query1, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(cartId)

	query2 := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE cart_id=$1 AND product_item_id=$2`
	err = tx.Exec(query2, cartId, productId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("2")

	var price int
	query3 := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(query3, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(price)
	fmt.Println("3")

	query4 := `UPDATE carts SET tottal=tottal-$1 WHERE user_id=$2`
	err = tx.Exec(query4, price, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("4")

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) ListCart(userId int) ([]response.Cart, error) {
	fmt.Println(userId)
	var items []response.Cart
	tx := c.DB.Begin()
	var cartId int
	query1 := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(query1, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}
	fmt.Println(2)

	query2 := `SELECT c.product_item_id, c.quantity, t.tottal 
	FROM cart_items c JOIN carts t ON c.cart_id = t.id 
	WHERE t.user_id = $1`
	err = tx.Raw(query2, userId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}

	fmt.Println(3)
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}
	return items, nil

}
