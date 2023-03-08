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
	userUseCase   services.UserUseCase
	cartUseCase   services.CartUsecase
	findIdUseCase services.FindIdUseCase
}

func NewUserHandler(usecase services.UserUseCase, cartUseCase services.CartUsecase, findIdUseCase services.FindIdUseCase) *UserHandler {
	return &UserHandler{
		userUseCase:   usecase,
		cartUseCase:   cartUseCase,
		findIdUseCase: findIdUseCase,
	}
}

// --------------------------------------------------UserSignUp------------------------------------------------------------

// @Summary user signup
// @ID user-signup
// @accept json
// @Param user_details body helperStruct.UserReq true "User Data"
// @Success 200 {object} response.UserData
// @Failure 400 {object} response.Response
// @Router /signup [post]
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
// @Summary user login
// @ID user-login
// @accept json
// @Param user_details body  helperStruct.LoginReq true "User Data"
// @Success 200 {object}
// @Failure 400 {object} response.Response
// @Router /userlogin [post]
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

// @Summary user logout
// @ID user-logout
// @Produce json
// @Success 200 {object}
// @Router /logout [get]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}

// @Summary user add address
// @ID user-add-address
// @accept json
// @Param user_address body  helperStruct.Address true "User address"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /addaddress [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Id, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address helperStruct.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.userUseCase.AddAddress(Id, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address added",
		Data:       nil,
		Errors:     nil,
	})
}

// func (cr *UserHandler) UpdateAddress(c *gin.Context) {
// 	cookie, err := c.Cookie("UserAuth")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't find Id",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	Id, err := cr.findIdUseCase.FindId(cookie)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't find Id",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	var address helperStruct.Address
// 	err = c.Bind(&address)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't bind",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	err = cr.userUseCase.UpdateAddress(Id, address)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "Can't update address",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.Response{
// 		StatusCode: 200,
// 		Message:    "address updated",
// 		Data:       nil,
// 		Errors:     nil,
// 	})
// }
