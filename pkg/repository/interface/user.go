package interfaces

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type UserRepository interface {
	UserSignUp(ctx context.Context, user helperStruct.UserReq) (response.UserData, error)
	UserLogin(ctx context.Context, email string) (domain.Users, error)
	IsSignIn(phno string) (bool, error)
	OtpLogin(phno string) (int, error)
	AddAddress(id int, address helperStruct.Address) error
	UpdateAddress(id, addressId int, address helperStruct.Address) error
}
