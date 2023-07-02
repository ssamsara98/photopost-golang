package bootstrap

import (
	"photopost/infrastructure"
	"photopost/lib"
	"photopost/src/controllers"
	"photopost/src/middlewares"
	"photopost/src/routes"
	"photopost/src/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	infrastructure.Module,
	services.Module,
	controllers.Module,
	middlewares.Module,
	routes.Module,
)
