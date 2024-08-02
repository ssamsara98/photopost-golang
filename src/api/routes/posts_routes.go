package routes

import (
	"github.com/gofiber/fiber/v2"
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

func (p PostsRoutes) Run(handler fiber.Router) {
	router := handler.Group("posts")

	router.Get("", p.paginationMiddleware.HandleCursor(), p.postsController.GetPostList)
	router.Get("p/:postId", p.postsController.GetPostById)
	router.Get("u/:userId", p.paginationMiddleware.Handle(), p.postsController.GetUserPostList)

	router.Use(p.jwtAuthMiddleware.Handle(constants.TokenAccess, true))
	router.Post("", p.dbTransactionMiddleware.Handle(), p.postsController.CreatePost)
	router.Post("upload", p.postsController.UploadPhoto)
	router.Get("mine", p.paginationMiddleware.Handle(), p.postsController.GetMyPostList)
	router.Delete("p/:postId", p.postsController.DeletePost)
}
