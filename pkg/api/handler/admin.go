package handler

import (
	"encoding/csv"
	"fmt"
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
// @Router /adminPanel/create-admin [post]
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

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
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

func (cr *AdminHandler) UploadImage(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()
	fmt.Println(form)

	files := form.File["images"]

	fmt.Println(files)

	for _, file := range files {
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "asset/uploads/"+file.Filename)

		err := cr.adminUseCase.UploadImage(file.Filename, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "cant upload images",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}
}
