package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/akshayur04/project-ecommerce/pkg/api/handlerUtil"
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

// @Summary UserSignUp
// @ID user-signup
// @Description Create a new user with the specified details.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body  helperStruct.UserReq true "User details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/signup [post]
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

// LoginWithEmail
// @Summary User Login
// @ID user-login-email
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body helperStruct.LoginReq true "User details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/login [post]
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
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})
}

// UserLogout
// @Summary User Logout
// @ID user-logout
// @Description Logs out a logged-in user from the E-commerce web api
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /user/logout [post]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}

// AddAddress
// @Summary User can add address
// @ID add-address
// @Description Add address
// @Tags Users
// @Accept json
// @Produce json
// @Param user_address body helperStruct.Address true "User address"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/address/add [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	Id, err := handlerUtil.GetUserIdFromContext(c)
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

// UpdateAddress
// @Summary User can update existing address
// @ID update-address
// @Description Update address
// @Tags Users
// @Accept json
// @Produce json
// @Param addressId path string true "addressId"
// @Param user_address body helperStruct.Address true "User address"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/address/update/{addressId} [patch]
func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	paramsId := c.Param("addressId")
	addressId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find ProductId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Id, err := handlerUtil.GetUserIdFromContext(c)
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
	err = cr.userUseCase.UpdateAddress(Id, addressId, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "address updated",
		Data:       nil,
		Errors:     nil,
	})
}

// UserProfile
// @Summary User can view their profile
// @ID user-profile
// @Description Users can visit their profile
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/view [get]
func (cr *UserHandler) Viewprfile(c *gin.Context) {
	Id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	Profile, err := cr.userUseCase.Viewprfile(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find userprofile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Profile",
		Data:       Profile,
		Errors:     nil,
	})
}

// UpdateUserProfile
// @Summary User can update their profile
// @ID update-user-profile
// @Description Users can update their profile
// @Tags Users
// @Accept json
// @Produce json
// @Param user_profile body helperStruct.UserReq true "User profile"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/edite [patch]
func (cr *UserHandler) UserEditProfile(c *gin.Context) {
	Id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var updatingDetails helperStruct.UserReq
	err = c.Bind(&updatingDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind details",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	updatedProfile, err := cr.userUseCase.UserEditProfile(Id, updatingDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find userprofile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Profile updated",
		Data:       updatedProfile,
		Errors:     nil,
	})
}

// UpdateUserPassword
// @Summary User can update their Password
// @ID update-user-Password
// @Description Users can update their Password
// @Tags Users
// @Accept json
// @Produce json
// @Param user_profile body helperStruct.UpdatePassword true "User password"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/profile/updatepassword [patch]
func (cr *UserHandler) UpdatePassword(c *gin.Context) {
	Id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var Passwords helperStruct.UpdatePassword
	err = c.Bind(&Passwords)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.UpdatePassword(Id, Passwords)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update password",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Password updated",
		Data:       nil,
		Errors:     nil,
	})
}
