package usecase

import (
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
