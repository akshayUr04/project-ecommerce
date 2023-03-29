package handler

import (
	// "context"
	// "net/http"

	"fmt"
	"net/http"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	config "github.com/akshayur04/project-ecommerce/pkg/config"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase  services.OtpUseCase
	userUseCase services.UserUseCase
	cfg         config.Config
}

func NewOtpHandler(cfg config.Config, otpUseCase services.OtpUseCase, userUseCase services.UserUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase:  otpUseCase,
		userUseCase: userUseCase,
		cfg:         cfg,
	}
}

// SendOtp
// @Summary Send OTP to user's mobile
// @ID send-otp
// @Description Send OTP to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param user_mobile body  helperStruct.OTPData true "User mobile number"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/otp/send [post]
func (cr *OtpHandler) SendOtp(c *gin.Context) {
	var phno helperStruct.OTPData
	err := c.Bind(&phno)
	if err != nil {
		fmt.Println("e1")
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		fmt.Println("e2")
		return
	}

	isSignIn, err := cr.userUseCase.IsSignIn(phno.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "No user with this phonenumber",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	fmt.Println(isSignIn)

	if !isSignIn {
		fmt.Println("login err")
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "no user found",
			Data:       nil,
			Errors:     nil,
		})
		fmt.Println("login err2")
		return
	}

	err = cr.otpUseCase.SendOtp(c.Request.Context(), phno)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "creatingfailed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "otp send",
		Data:       nil,
		Errors:     nil,
	})
}

// ValidateOtp
// @Summary Validate the OTP to user's mobile
// @ID validate-otp
// @Description Validate the  OTP sent to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param otp body helperStruct.VerifyOtp true "OTP sent to user's mobile number"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/otp/verif [post]
func (cr *OtpHandler) ValidateOtp(c *gin.Context) {
	var otpDetails helperStruct.VerifyOtp
	err := c.Bind(&otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	resp, err := cr.otpUseCase.ValidateOtp(otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	} else if *resp.Status != "approved" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "incorect",
			Data:       nil,
			Errors:     "incorect",
		})
		return
	}
	ss, err := cr.userUseCase.OtpLogin(otpDetails.User.PhoneNumber)
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
