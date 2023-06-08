package users

import "github.com/gin-gonic/gin"

type UsersModuleInf interface {
	Router(rg *gin.RouterGroup)
}

type UsersModule struct {
	UsersControllerV1 UsersControllerV1Inf
}

func NewUsersModule(
	usersControllerV1 UsersControllerV1Inf,
) *UsersModule {
	return &UsersModule{
		usersControllerV1,
	}
}

func (um UsersModule) Router(rg *gin.RouterGroup) {
	usersRoutesV1 := rg.Group("users")

	um.UsersControllerV1.Run(usersRoutesV1)
}
