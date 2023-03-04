package interfaces

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admis helperStruct.CreateAdmin, createrId int) (response.AdminData, error)
	AdminLogin(admin helperStruct.LoginReq) (string, error)
	BlockUser(body helperStruct.BlockData, adminId int) error
	UnblockUser(id int) error
}
