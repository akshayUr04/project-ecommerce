package repository

import (
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

	//finding cart id coresponding to the user
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Check whether the product exists in the cart_items
	var cartItemId int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id = $1 AND product_item_id = $2 LIMIT 1`
	err = tx.Raw(cartItemCheck, cartId, productId).Scan(&cartItemId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if cartItemId == 0 {
		addToCart := `INSERT INTO cart_items (carts_id,product_item_id,quantity)VALUES($1,$2,1)`
		err = tx.Exec(addToCart, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updatCart := `UPDATE cart_items SET quantity = cart_items.quantity+1 WHERE id = $1 `
		err = tx.Exec(updatCart, cartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//finding the price of the product
	var price int
	findPrice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Updating the tottal in cart table
	updateTottal := `UPDATE carts SET tottal=carts.tottal+$1 WHERE user_id=$2`
	err = tx.Exec(updateTottal, price, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) RemoveFromCart(userId, productId int) error {
	tx := c.DB.Begin()

	//Find cart id
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Find the qty of the product in cart
	var qty int
	findQty := `SELECT quantity FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
	err = tx.Raw(findQty, cartId, productId).Scan(&qty).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//If the qty is 1 dlt the product from the cart
	if qty == 1 {
		dltItem := `DELET FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
		err := tx.Exec(dltItem).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else { // If there is  more than one product reduce the qty by 1
		updateQty := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Exec(updateQty, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//Find the price of the product item
	var price int
	productPrice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(productPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Reduce the price of the cart tottal with price of the product
	updateTottal := `UPDATE carts SET tottal=tottal-$1 WHERE user_id=$2`
	err = tx.Exec(updateTottal, price, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *CartDatabase) ListCart(userId int) ([]response.Cart, error) {

	tx := c.DB.Begin()
	var items []response.Cart

	//Find the cart id of the user
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}

	//Get the details of the items in the cart and theri total
	cartItems := `SELECT pi.sku,pi.color,pi.price,ci.quantity,c.tottal FROM product_items pi JOIN 
			cart_items ci ON pi.id=ci.product_item_id 
			JOIN carts c ON ci.carts_id=c.id 
			WHERE c.user_id= $1`
	err = tx.Raw(cartItems, userId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return []response.Cart{}, err
	}
	return items, nil

}
