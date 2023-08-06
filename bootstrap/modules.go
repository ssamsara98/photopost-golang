package bootstrap

import (
	"photopost/api/controllers"
	"photopost/api/middlewares"
	"photopost/api/routes"
	"photopost/api/services"
	"photopost/infrastructure"
	"photopost/lib"

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
