package bootstrap

import (
	"github.com/ssamsara98/photopost-golang/src/api/controllers"
	"github.com/ssamsara98/photopost-golang/src/api/middlewares"
	"github.com/ssamsara98/photopost-golang/src/api/routes"
	"github.com/ssamsara98/photopost-golang/src/api/services"
	"github.com/ssamsara98/photopost-golang/src/helpers"
	"github.com/ssamsara98/photopost-golang/src/infrastructure"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	lib.Module,
	infrastructure.Module,
	helpers.Module,
	services.Module,
	controllers.Module,
	middlewares.Module,
	routes.Module,
)
