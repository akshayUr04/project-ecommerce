package handler

type CartHandler struct {
	cartUsecase services.carttUsecase
}

func NewCartHandler(cartUsecase services.carttUsecase) *CartHandler {
	return &CartHandler{
		productUsecase: productUsecase,
	}
}
