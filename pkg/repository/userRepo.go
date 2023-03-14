package repository

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// --------------------------------------------------UserSignUp------------------------------------------------------------

func (c *userDatabase) UserSignUp(ctx context.Context, user helperStruct.UserReq) (response.UserData, error) {
	var userData response.UserData
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}

// --------------------------------------------------UserLogin------------------------------------------------------------

func (c *userDatabase) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}

// to check whether there is any user exists coresponding to the given mobile number
func (c *userDatabase) IsSignIn(phno string) (bool, error) {
	quwery := "select exists(select 1 from users where mobile=?)"
	var isSignIng bool
	err := c.DB.Raw(quwery, phno).Scan(&isSignIng).Error
	return isSignIng, err

}

func (c *userDatabase) OtpLogin(phno string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE mobile=?"
	err := c.DB.Raw(query, phno).Scan(&id).Error
	return id, err
}

func (c *userDatabase) AddAddress(id int, address helperStruct.Address) error {

	//Check if the new address is being set as default
	if address.IsDefault { //Change the default address into false
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE users_id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, false, id, true).Error

		if err != nil {
			return err
		}
	}
	//Insert the new address
	insert := `INSERT INTO addresses (users_id,house_number,street,city, district,landmark,pincode,is_default)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	err := c.DB.Exec(insert, id, address.
		House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault).Error
	return err
}

func (c *userDatabase) UpdateAddress(id, addressId int, address helperStruct.Address) error {
	//Check if the new address is being set as default
	if address.IsDefault { //Change the default address into false
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE users_id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, false, id, true).Error

		if err != nil {
			return err
		}
	}
	//Update the address
	update := `UPDATE addresses SET 
		house_number=$1,street=$2,city=$3, district=$4,landmark=$5,pincode=$6,is_default=$7 WHERE users_id=$8 AND id=$9`
	err := c.DB.Exec(update,
		address.House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault,
		id,
		addressId).Error
	return err
}
