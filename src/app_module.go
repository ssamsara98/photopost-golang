package src

import (
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppModuleInf interface {
	Router(r *gin.Engine)
}

type AppModule struct {
	Log               *log.Logger
	DB                *gorm.DB
	JWTAuthHelper     *lib.JWTAuthHelper
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	AppController     AppControllerInf
}

func NewAppModule(
	log *log.Logger,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appController AppControllerInf,
) *AppModule {
	return &AppModule{
		log,
		db,
		jwtAuthHelper,
		jwtAuthMiddleware,
		appController,
	}
}

func (app AppModule) Router(r *gin.Engine) {
	app.AppController.Run(r)
}
