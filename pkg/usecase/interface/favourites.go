package interfaces

type FavouritesUsecase interface {
	AddToFavourites(productId, userId int) error
	RemoveFromFav(userId, productId int) error
}
