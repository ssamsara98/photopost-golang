package routes

import (
	"photopost/api/controllers"
	"photopost/api/middlewares"

	"github.com/gin-gonic/gin"
)

type PostsRoutes struct {
	postsController         *controllers.PostsController
	jwtAuthMiddleware       *middlewares.JWTAuthMiddleware
	paginationMiddleware    *middlewares.PaginationMiddleware
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware
}

func NewPostsRoutes(
	postsController *controllers.PostsController,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	paginationMiddleware *middlewares.PaginationMiddleware,
	dbTransactionMiddleware *middlewares.DBTransactionMiddleware,
) *PostsRoutes {
	return &PostsRoutes{
		postsController,
		jwtAuthMiddleware,
		paginationMiddleware,
		dbTransactionMiddleware,
	}
}

func (p PostsRoutes) Run(handler *gin.RouterGroup) {
	router := handler.Group("posts")

	router.POST("/", p.jwtAuthMiddleware.Handle(), p.dbTransactionMiddleware.Handle(), p.postsController.CreatePost)
	router.POST("/upload", p.jwtAuthMiddleware.Handle(), p.postsController.UploadPhoto)
	router.GET("/", p.paginationMiddleware.Handle(), p.postsController.GetPostList)
	router.GET("/mine", p.jwtAuthMiddleware.Handle(), p.paginationMiddleware.Handle(), p.postsController.GetMyPostList)
	router.GET("/p/:postId", p.postsController.GetPost)
	router.GET("/u/:userId", p.paginationMiddleware.Handle(), p.postsController.GetUserPostList)
}
