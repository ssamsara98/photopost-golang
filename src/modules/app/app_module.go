package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AppModuleInf interface {
	Router(r *gin.Engine)
}

type AppModule struct {
	AppController AppControllerInf
}

func NewAppModule(
	appController AppControllerInf,
) AppModuleInf {
	return &AppModule{
		appController,
	}
}

func (app AppModule) Router(r *gin.Engine) {
	app.AppController.Run(r)
}

var AppModuleFx = fx.Options(
	fx.Provide(
		NewAppModule,
		NewAppController,
		NewAppService,
	),
)
