package usecase

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type FavouritesUsecase struct {
	favouritesRepository interfaces.FavouritesRepository
}

func NewFavouritesUsecase(favouritesRepository interfaces.FavouritesRepository) services.FavouritesUsecase {
	return &FavouritesUsecase{
		favouritesRepository: favouritesRepository,
	}
}

func (c *FavouritesUsecase) AddToFavourites(productId, userId int) error {
	err := c.favouritesRepository.AddToFavourites(productId, userId)
	return err
}

func (c *FavouritesUsecase) RemoveFromFav(userId, productId int) error {
	err := c.favouritesRepository.RemoveFromFav(userId, productId)
	return err
}

func (c *FavouritesUsecase) ViewFavourites(userId int) ([]response.ProductItem, error) {
	favourites, err := c.favouritesRepository.ViewFavourites(userId)
	return favourites, err
}
