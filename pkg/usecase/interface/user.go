package interfaces

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(ctx context.Context, user helperStruct.UserReq) (response.UserData, error)
	UserLogin(ctx context.Context, user helperStruct.LoginReq) (string, error)
	OtpLogin(phno string) (string, error)
	IsSignIn(phno string) (bool, error)
}
