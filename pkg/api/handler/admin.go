package handler

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/akshayur04/project-ecommerce/pkg/api/handlerUtil"
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUsecase
}

func NewAdminHandler(adminUseCae services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUseCae,
	}
}

// CreateAdmin
// @Summary Create a new admin from admin panel
// @ID create-admin
// @Description Super admin can create a new admin from admin panel.
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_details body helperStruct.CreateAdmin true "New Admin details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/createadmin [post]
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
	createrId, err := handlerUtil.GetAdminIdFromContext(c)
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

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Admin created",
		Data:       admin,
		Errors:     nil,
	})
}

// AdminLogin
// @Summary Admin Login
// @ID admin-login
// @Description Admin login
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_credentials body helperStruct.LoginReq true "Admin login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/adminlogin [post]
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
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined success fully",
		Data:       nil,
		Errors:     nil,
	})
}

// AdminLogout
// @Summary Admin Logout
// @ID admin-logout
// @Description Logs out a logged-in admin from the E-commerce web api admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400
// @Router /admin/adminlogout [post]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})

}

// BlockUser
// @Summary Admin can bolock users
// @ID block-users
// @Description Admins can block users
// @Tags Admin
// @Accept json
// @Produce json
// @Param blocking_details body helperStruct.BlockData true "User bolocking details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/blockuser/{id} [patch]
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
	adminId, err := handlerUtil.GetAdminIdFromContext(c)
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

// UnBlockUser
// @Summary Admin can unbolock a blocked user
// @ID unblock-users
// @Description Admins can block users
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be blocked"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/unblockuser/{id} [patch]
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

// FindUser
// @Summary Admin can find a user
// @ID find-users
// @Description Admins can find users with id
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be found"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/finduser/{id} [get]
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

// FindAllUsers
// @Summary Admin can find all registered users
// @ID find-all-users
// @Description Admin can find all registered users
// @Tags Admin
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items to retrieve per page"
// @Param query query string false "Search query string"
// @Param filter query string false "Filter criteria for the users"
// @Param sort_by query string false "Sorting criteria for the users"
// @Param sort_desc query bool false "Sorting in descending order"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/findall [get]
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

// AdminDashboard
// @Summary Admin Dashboard
// @ID admin-dashboard
// @Description Admin can access dashboard and view details regarding orders, products, etc.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/dashboard/get [get]
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

// ViewSalesReport
// @Summary Admin can view sales report
// @ID view-sales-report
// @Description Admin can view the sales report
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/sales/get [get]
func (cr *AdminHandler) ViewSalesReport(c *gin.Context) {
	sales, err := cr.adminUseCase.ViewSalesReport()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Sales report",
		Data:       sales,
		Errors:     nil,
	})

}

// DownloadSalesReport
// @Summary Admin can download sales report
// @ID download-sales-report
// @Description Admin can download sales report in .csv format
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/sales/download [get]
func (cr *AdminHandler) DownloadSalesReport(c *gin.Context) {
	sales, err := cr.adminUseCase.ViewSalesReport()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Set headers so browser will download the file
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=sales.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)

	// Write CSV header row
	headers := []string{"Name", "PaymentType", "OrderDate", "OrderTotal"}
	if err := wr.Write(headers); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Write data rows
	for _, sale := range sales {
		row := []string{sale.Name, sale.PaymentType, sale.OrderDate.Format("2006-01-02 15:04:05"), strconv.Itoa(sale.OrderTotal)}
		if err := wr.Write(row); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
	if err := wr.Error(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}
