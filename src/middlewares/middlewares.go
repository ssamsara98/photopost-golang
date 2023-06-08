package middlewares

import "go.uber.org/fx"

var MiddlewaresModuleFx = fx.Options(
	fx.Provide(
		NewJWTAuthMiddleware,
	),
)
