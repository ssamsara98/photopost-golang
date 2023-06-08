package posts

import "github.com/gin-gonic/gin"

type PostsModuleInf interface {
	Router(rg *gin.RouterGroup)
}

type PostsModule struct {
	PostsControllerV1 PostsControllerV1Inf
}

func NewPostsModule(
	postsControllerV1 PostsControllerV1Inf,
) *PostsModule {
	return &PostsModule{
		postsControllerV1,
	}
}

func (pc PostsModule) Router(rg *gin.RouterGroup) {
	postsRoutesV1 := rg.Group("posts")

	pc.PostsControllerV1.Run(postsRoutesV1)
}
