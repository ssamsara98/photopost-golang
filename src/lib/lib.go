package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(GetEnv),
	fx.Provide(GetLogger),
)
