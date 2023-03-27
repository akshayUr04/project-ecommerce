package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type AdminRepository interface {
	IsSuperAdmin(createrId int) (bool, error)
	CreateAdmin(admin helperStruct.CreateAdmin) (response.AdminData, error)
	AdminLogin(email string) (domain.Admins, error)
	BlockUser(body helperStruct.BlockData, adminId int) error
	UnblockUser(id int) error
	FindUser(id int) (response.UserDetails, error)
	FindAll() ([]response.UserDetails, error)
	GetDashBoard() (response.DashBoard, error)
	ViewSalesReport() ([]response.SalesReport, error)
	UploadImage(filepath string, productId int) error
}
