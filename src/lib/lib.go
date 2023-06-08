package lib

import (
	"log"

	"go.uber.org/fx"
)

var LibModuleFx = fx.Options(
	fx.Provide(
		log.Default,
		NewEnv,
		NewDatabase,
		NewJWTAuthHelper,
		NewS3Service,
	),
)
