package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/api/controllers"
	"github.com/ssamsara98/photopost-golang/src/api/middlewares"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/lib"
)

type PostsRoutes struct {
	logger                  *lib.Logger
	paginationMiddleware    *middlewares.PaginationMiddleware
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware
	jwtAuthMiddleware       *middlewares.JWTAuthMiddleware
	postsController         *controllers.PostsController
}

func NewPostsRoutes(
	logger *lib.Logger,
	paginationMiddleware *middlewares.PaginationMiddleware,
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	postsController *controllers.PostsController,
) *PostsRoutes {
	return &PostsRoutes{
		logger,
		paginationMiddleware,
		dbTransactionMiddleware,
		jwtAuthMiddleware,
		postsController,
	}
}

func (p PostsRoutes) Run(handler *gin.RouterGroup) {
	router := handler.Group("posts")

	router.GET("", p.paginationMiddleware.HandleCursor(), p.postsController.GetPostList)
	router.GET("p/:postId", p.postsController.GetPostById)
	router.GET("u/:userId", p.paginationMiddleware.Handle(), p.postsController.GetUserPostList)

	router.Use(p.jwtAuthMiddleware.Handle(constants.TokenAccess, true))
	router.POST("", p.dbTransactionMiddleware.Handle(), p.postsController.CreatePost)
	router.POST("upload", p.postsController.UploadPhoto)
	router.GET("mine", p.paginationMiddleware.Handle(), p.postsController.GetMyPostList)
	router.DELETE("p/:postId", p.postsController.DeletePost)
}
