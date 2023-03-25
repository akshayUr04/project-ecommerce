package interfaces

import "github.com/akshayur04/project-ecommerce/pkg/common/response"

type FavouritesRepository interface {
	AddToFavourites(productId, userId int) error
	RemoveFromFav(userId, productId int) error
	ViewFavourites(userId int) ([]response.ProductItem, error)
}
