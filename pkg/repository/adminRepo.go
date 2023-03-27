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
	query := `INSERT INTO admins (name,email,password,is_super_admin,created_at)
								  VALUES($1,$2,$3,$4,NOW())
								  RETURNING id,name,email,is_super_admin`

	err := c.DB.Raw(query, admin.Name, admin.Email, admin.Password, admin.IsSuper).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) BlockUser(body helperStruct.BlockData, adminId int) error {
	// Start a transaction
	tx := c.DB.Begin()
	// Execute the first SQL command (UPDATE)
	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Execute the second SQL command (INSERT)
	if err := tx.Exec("INSERT INTO user_infos (users_id, reason_for_blocking, blocked_at, blocked_by) VALUES (?, ?, NOW(), ?)", body.UserId, body.Reason, adminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	// If all commands were executed successfully, return nil
	return nil

}

func (c *adminDatabase) UnblockUser(id int) error {
	tx := c.DB.Begin()
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_infos SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *adminDatabase) FindUser(id int) (response.UserDetails, error) {
	var userDetails response.UserDetails
	qury := `SELECT users.name,
			 users.email, 
			 users.mobile,  
			 users.is_blocked, 
			 infos.blocked_by,
			 infos.blocked_at,
			 infos.reason_for_blocking 
			 FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id
			 WHERE users.id = $1;`

	err := c.DB.Raw(qury, id).Scan(&userDetails).Error
	return userDetails, err
}

func (c *adminDatabase) FindAll() ([]response.UserDetails, error) {
	var users []response.UserDetails
	qury := `SELECT users.name,
			 users.email, 
			 users.mobile,  
			 users.is_blocked, 
			 infos.blocked_by,
			 infos.blocked_at,
			 infos.reason_for_blocking 
			 FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id`
	err := c.DB.Raw(qury).Scan(&users).Error
	return users, err
}

func (c *adminDatabase) GetDashBoard() (response.DashBoard, error) {
	tx := c.DB.Begin()
	var dashBoard response.DashBoard
	getDasheBoard := `SELECT SUM(oi.quantity*oi.price)as Total_Revenue,
			SUM (oi.quantity)as Total_Products_Selled,
			COUNT(DISTINCT o.id)as Total_Orders FROM orders o
			JOIN order_items oi on o.id=oi.orders_id
			WHERE o.order_status_id=$1`
	if err := tx.Raw(getDasheBoard, 1).Scan(&dashBoard).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	// getDasheBoard := `SELECT SUM(quantity*price)as Total_Revenue,
	// 		SUM (quantity)as Total_Products_Selled FROM order_items`
	// if err := tx.Raw(getDasheBoard).Scan(&dashBoard).Error; err != nil {
	// 	tx.Rollback()
	// 	return response.DashBoard{}, err
	// }

	// getOrderNo := `SELECT COUNT(id)FROM orders WHERE order_status_id=$1`
	// if err := tx.Raw(getOrderNo, 1).Scan(&dashBoard.TotalOrders).Error; err != nil {
	// 	tx.Rollback()
	// 	return response.DashBoard{}, err
	// }

	getTotalUsers := `SELECT COUNT(id)AS TotalUsers FROM users`
	if err := tx.Raw(getTotalUsers).Scan(&dashBoard.TotalUsers).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	return dashBoard, nil
}

func (c *adminDatabase) ViewSalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}

func (c *adminDatabase) UploadImage(filepath string, productId int) error {
	uploadImage := `INSERT INTO images (product_item_id,file_name)VALUES($1,$2)`
	err := c.DB.Exec(uploadImage, productId, filepath).Error
	return err
}
