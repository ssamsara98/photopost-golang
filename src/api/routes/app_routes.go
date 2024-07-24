package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/api/controllers"
	"github.com/ssamsara98/photopost-golang/src/api/middlewares"
	"github.com/ssamsara98/photopost-golang/src/constants"
)

type AppRoutes struct {
	appController           *controllers.AppController
	jwtAuthMiddleware       *middlewares.JWTAuthMiddleware
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware
}

func NewAppRoutes(
	appController *controllers.AppController,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware,
) *AppRoutes {
	return &AppRoutes{
		appController,
		jwtAuthMiddleware,
		dbTransactionMiddleware,
	}
}

func (app AppRoutes) Run(handler *gin.RouterGroup) {
	handler.GET("", app.appController.Home)
	handler.POST("register", app.dbTransactionMiddleware.Handle(), app.appController.Register)
	handler.POST("login", app.appController.Login)
	handler.GET("me", app.jwtAuthMiddleware.Handle(constants.TokenAccess, true), app.appController.Me)
	handler.PATCH("me", app.jwtAuthMiddleware.Handle(constants.TokenAccess, true), app.appController.UpdateProfile)
	handler.GET("token/check", app.jwtAuthMiddleware.Handle(constants.TokenAccess, false), app.appController.TokenCheck)
	handler.GET("token/refresh", app.jwtAuthMiddleware.Handle(constants.TokenRefresh, true), app.appController.TokenRefresh)
}
