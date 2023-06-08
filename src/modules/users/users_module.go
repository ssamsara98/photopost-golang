package users

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UsersModuleInf interface {
	Router(rg *gin.RouterGroup)
}

type UsersModule struct {
	UsersControllerV1 UsersControllerV1Inf
}

func NewUsersModule(
	usersControllerV1 UsersControllerV1Inf,
) UsersModuleInf {
	return &UsersModule{
		UsersControllerV1: usersControllerV1,
	}
}

func (um UsersModule) Router(rg *gin.RouterGroup) {
	usersRoutesV1 := rg.Group("users")

	um.UsersControllerV1.Run(usersRoutesV1)
}

var UsersModuleFx = fx.Options(
	fx.Provide(
		NewUsersModule,
		NewUsersControllerV1,
		NewUsersServiceV1,
	),
)
