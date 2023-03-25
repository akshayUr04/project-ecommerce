package repository

import (
	"fmt"

	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type FavouritesDatabase struct {
	DB *gorm.DB
}

func NewFavouritesRepository(DB *gorm.DB) interfaces.FavouritesRepository {
	return &FavouritesDatabase{DB}
}

func (c *FavouritesDatabase) AddToFavourites(productId, userId int) error {
	tx := c.DB.Begin()
	//Check if the item is alredy present inside the fav
	var isPresent bool
	isIn := `SELECT EXISTS(SELECT 1 FROM favourites WHERE user_id = $1 AND item_id = $2)`
	err := tx.Raw(isIn, userId, productId).Scan(&isPresent).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if isPresent {
		tx.Rollback()
		return fmt.Errorf("the item is alredy added to fav")
	}
	insertToFav := `INSERT INTO favourites (user_id,item_id)VALUES($1,$2)`
	err = tx.Exec(insertToFav, userId, productId).Error
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

func (c *FavouritesDatabase) RemoveFromFav(userId, productId int) error {
	tx := c.DB.Begin()
	//Check if the item is alredy present inside the fav
	var isPresent bool
	isIn := `SELECT EXISTS(SELECT 1 FROM favourites WHERE user_id = $1 AND item_id = $2)`
	err := tx.Raw(isIn, userId, productId).Scan(&isPresent).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if !isPresent {
		tx.Rollback()
		return fmt.Errorf("the item is not present in the fav")
	}
	removeItem := `DELETE FROM favourites WHERE user_id=$1 AND item_id=$2`
	err = tx.Exec(removeItem, userId, productId).Error
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

func (c *FavouritesDatabase) ViewFavourites(userId int) ([]response.ProductItem, error) {
	var favourites []response.ProductItem
	viewWishlist := `SELECT pi.sku,
	pi.qty_in_stock,
	pi.color,
	pi.ram,
	pi.battery,
	pi.screen_size,
	pi.storage,
	pi.camera,
	pi.price,
	pi.imag,
	p.product_name,
	p.description,
	p.brand,
	c.category_name 
	FROM favourites f JOIN product_items pi ON f.item_id = pi.id 
	JOIN products p ON pi.product_id = p.id 
	JOIN categories c ON p.category_id = c.id WHERE f.user_id=$1`
	err := c.DB.Raw(viewWishlist, userId).Scan(&favourites).Error
	return favourites, err
}
