package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAppService),
	fx.Provide(NewS3Service),
	fx.Provide(NewUsersService),
	fx.Provide(NewPostsService),
)
