package routes

import (
	"github.com/ssamsara98/photopost-golang/src/api/controllers"
	"github.com/ssamsara98/photopost-golang/src/api/middlewares"
	"github.com/ssamsara98/photopost-golang/src/lib"

	"github.com/gin-gonic/gin"
)

type UsersRoutes struct {
	logger               *lib.Logger
	paginationMiddleware *middlewares.PaginationMiddleware
	usersController      *controllers.UsersController
}

func NewUsersRoutes(
	logger *lib.Logger,
	paginationMiddleware *middlewares.PaginationMiddleware,
	usersController *controllers.UsersController,
) *UsersRoutes {
	return &UsersRoutes{
		logger,
		paginationMiddleware,
		usersController,
	}
}

func (u UsersRoutes) Run(handler *gin.RouterGroup) {
	router := handler.Group("users")

	router.GET("", u.paginationMiddleware.Handle(), u.usersController.GetUserList)
	router.GET("cursor", u.paginationMiddleware.HandleCursor(), u.usersController.GetUserListCursor)
	router.GET("u/:userId", u.usersController.GetUserByID)
}
