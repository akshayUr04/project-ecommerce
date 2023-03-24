package interfaces

type FavouritesRepository interface {
	AddToFavourites(productId, userId int) error
	RemoveFromFav(userId, productId int) error
}
