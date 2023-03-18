package handler

import (
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase  services.AdminUsecase
	findIdUseCase services.FindIdUseCase
}

func NewAdminHandler(adminUseCae services.AdminUsecase, findIdUseCase services.FindIdUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase:  adminUseCae,
		findIdUseCase: findIdUseCase,
	}
}

func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var adminData helperStruct.CreateAdmin
	err := c.Bind(&adminData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	createrId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	admin, err := cr.adminUseCase.CreateAdmin(c.Request.Context(), adminData, createrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Create Admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Admin created",
		Data:       admin,
		Errors:     nil,
	})

}

func (cr *AdminHandler) AdminLoging(c *gin.Context) {
	var admin helperStruct.LoginReq
	err := c.Bind(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ss, err := cr.adminUseCase.AdminLogin(admin)
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
	c.SetCookie("AdminAuth", ss, 3600*24*30, "", "", false, true)
}

func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout succesfully",
	})
}

func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body helperStruct.BlockData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	cookie, err := c.Cookie("AdminAuth")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminId, err := cr.findIdUseCase.FindId(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.BlockUser(body, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *AdminHandler) FindUser(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	user, err := cr.adminUseCase.FindUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user details",
		Data:       user,
		Errors:     nil,
	})
}

func (cr *AdminHandler) FindAllUsers(c *gin.Context) {
	users, err := cr.adminUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "users are",
		Data:       users,
		Errors:     nil,
	})

}

func (cr *AdminHandler) AdminDashBoard(c *gin.Context) {
	dashBoard, err := cr.adminUseCase.GetDashBoard()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get dashboard",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Dash board",
		Data:       dashBoard,
		Errors:     nil,
	})
}
