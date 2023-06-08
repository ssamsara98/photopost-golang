package posts

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PostsModuleInf interface {
	Router(rg *gin.RouterGroup)
}

type PostsModule struct {
	PostsControllerV1 PostsControllerV1Inf
}

func NewPostsModule(
	postsControllerV1 PostsControllerV1Inf,
) PostsModuleInf {
	return &PostsModule{
		postsControllerV1,
	}
}

func (pc PostsModule) Router(rg *gin.RouterGroup) {
	postsRoutesV1 := rg.Group("posts")

	pc.PostsControllerV1.Run(postsRoutesV1)
}

var PostsModulefx = fx.Options(
	fx.Provide(
		NewPostsModule,
		NewPostsControllerV1,
		NewPostsServiceV1,
	),
)
