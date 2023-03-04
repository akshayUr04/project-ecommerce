package repository

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (c *adminDatabase) IsSuperAdmin(createrId int) (bool, error) {
	var isSuper bool
	query := "SELECT is_super_admin FROM admins WHERE id=$1"
	err := c.DB.Raw(query, createrId).Scan(&isSuper).Error
	return isSuper, err
}

func (c *adminDatabase) CreateAdmin(admin helperStruct.CreateAdmin) (response.AdminData, error) {
	var adminData response.AdminData
	query := "INSERT INTO admins (name,email,password,is_super_admin)VALUES($1,$2,$3,$4)RETURNING id,name,email,is_super_admin"
	err := c.DB.Raw(query, admin.Name, admin.Email, admin.Password, admin.IsSuper).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}
