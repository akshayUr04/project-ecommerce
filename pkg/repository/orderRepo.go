package repository

import (
	"fmt"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}

func (c *OrderDatabase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	tx := c.DB.Begin()

	//Find the cart id and tottal of the cart
	var cart helperStruct.Cart
	findCart := `SELECT id,tottal FROM carts WHERE user_id=? `
	err := tx.Raw(findCart, id).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Find the default address of the user
	var addressId int
	address := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(address, id).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Add the details to the orders and return the orderid
	var order domain.Orders
	insetOrder := `INSERT INTO orders (user_id,order_date,payment_type_id,shipping_address,order_total,order_status)
		VALUES($1,NOW(),$2,$3,$4,'shipped') RETURNING *`
	err = tx.Raw(insetOrder, id, paymentTypeId, addressId, cart.Tottal).Scan(&order).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Get the cart item details of the user
	var cartItmes []helperStruct.CartItems
	cartDetail := `select ci.product_item_id,ci.quantity,pi.price,pi.qty_in_stock  from cart_items ci join product_items pi on ci.product_item_id = pi.id where ci.carts_id=$1`
	err = tx.Raw(cartDetail, cart.Id).Scan(&cartItmes).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Add the items in the cart into the orderitems one by one
	for _, items := range cartItmes {
		//check whether the item is available
		if items.Quantity > items.QtyInStock {
			return domain.Orders{}, fmt.Errorf("out of stock")
		}
		insetOrderItems := `INSERT INTO order_items (orders_id,product_item_id,quantity,price) VALUES($1,$2,$3,$4)`
		err = tx.Exec(insetOrderItems, order.Id, items.ProductItemId, items.Quantity, items.Price).Error

		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//Remove the items from the cart_items
	for _, items := range cartItmes {
		removeCartItems := `DELETE FROM cart_items WHERE carts_id =$1 AND product_item_id=$2`
		err = tx.Exec(removeCartItems, cart.Id, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//Reduce the product qty in stock details
	for _, items := range cartItmes {
		updateQty := `UPDATE product_items SET qty_in_stock=product_items.qty_in_stock-$1 WHERE id=$2`
		err = tx.Exec(updateQty, items.Quantity, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	return order, nil
}
