package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	cartUseCase services.CartUsecase
}

func NewUserHandler(usecase services.UserUseCase, cartUseCase services.CartUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		cartUseCase: cartUseCase,
	}
}

// --------------------------------------------------UserSignUp------------------------------------------------------------

func (cr *UserHandler) UserSignUp(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	// defer cancel()
	var user helperStruct.UserReq
	err := c.Bind(&user)
	fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userData, err := cr.userUseCase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.cartUseCase.CreateCart(userData.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable create cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}

// --------------------------------------------------UserLogin------------------------------------------------------------

func (cr *UserHandler) UserLogin(c *gin.Context) {

	var user helperStruct.LoginReq
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ss, err := cr.userUseCase.UserLogin(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
}

func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}
