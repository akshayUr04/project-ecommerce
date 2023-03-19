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
	if cart.Tottal == 0 {
		return domain.Orders{}, fmt.Errorf("no items in cart")
	}
	//Find the default address of the user
	var addressId int
	address := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(address, id).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if addressId == 0 {
		return domain.Orders{}, fmt.Errorf("add address pls")
	}

	//Add the details to the orders and return the orderid
	var order domain.Orders
	insetOrder := `INSERT INTO orders (user_id,order_date,payment_type_id,shipping_address,order_total,order_status_id)
		VALUES($1,NOW(),$2,$3,$4,1) RETURNING *`
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

	//Update the cart total
	updateCart := `UPDATE carts SET tottal=0 WHERE user_id=?`
	err = tx.Exec(updateCart, id).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
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

	//update the PaymentDetails table with OrdersID, OrderTotal, PaymentTypeID, PaymentStatusID
	createPaymentDetails := `INSERT INTO payment_details
			(orders_id,
			order_total,
			payment_type_id,
			payment_status_id,
			updated_at)
			VALUES($1,$2,$3,$4,NOW())`
	if err = tx.Exec(createPaymentDetails, order.Id, order.OrderTotal, paymentTypeId, 1).Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	return order, nil
}

func (c *OrderDatabase) UserCancelOrder(orderId, userId int) error {
	tx := c.DB.Begin()

	//find the orderd product and qty and update the product_items with those
	var items []helperStruct.CartItems
	findProducts := `SELECT product_item_id,quantity FROM order_items WHERE orders_id=?`
	err := tx.Raw(findProducts, orderId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(items) == 0 {
		return fmt.Errorf("no order found with this id")
	}
	for _, item := range items {
		updateProductItem := `UPDATE product_items SET qty_in_stock=qty_in_stock+$1 WHERE id=$2`
		err = tx.Exec(updateProductItem, item.Quantity, item.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//Remove the items from order_items
	removeItems := `DELETE FROM order_items WHERE orders_id=$1`
	err = tx.Exec(removeItems, orderId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update the order status as canceled
	cancelOrder := `UPDATE orders SET order_status_id=$1 WHERE id=$2 AND user_id=$3`
	err = tx.Exec(cancelOrder, 4, orderId, userId).Error
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

func (c *OrderDatabase) ListOrder(userId, orderId int) (domain.Orders, error) {
	var order domain.Orders
	findOrder := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err := c.DB.Raw(findOrder, userId, orderId).Scan(&order).Error
	return order, err
}

func (c *OrderDatabase) ListAllOrders(userId int) ([]domain.Orders, error) {
	var orders []domain.Orders

	findOrders := `SELECT * FROM orders WHERE user_id=?`
	err := c.DB.Raw(findOrders, userId).Scan(&orders).Error
	return orders, err
}
