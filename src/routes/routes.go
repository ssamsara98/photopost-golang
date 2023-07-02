package routes

import (
	"errors"
	"net/http"
	"photopost/infrastructure"
	"photopost/src/middlewares"
	"photopost/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewAppRoutes),
	fx.Provide(NewUsersRoutes),
	fx.Provide(NewPostsRoutes),
)

// Route interface
type IRoute interface {
	Setup()
}

// Routes contains multiple routes
type Routes struct {
	handler             infrastructure.Router
	rateLimitMiddleware *middlewares.RateLimitMiddleware
	appRoutes           *AppRoutes
	usersRoutes         *UsersRoutes
	postsRoutes         *PostsRoutes
}

// NewRoutes sets up routes
func NewRoutes(
	handler infrastructure.Router,
	rateLimitMiddleware *middlewares.RateLimitMiddleware,
	appRoutes *AppRoutes,
	usersRoutes *UsersRoutes,
	postsRoutes *PostsRoutes,
) Routes {
	return Routes{
		handler,
		rateLimitMiddleware,
		appRoutes,
		usersRoutes,
		postsRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	r.handler.Use(r.rateLimitMiddleware.Handle())

	apiV1 := r.handler.Group("v1")

	r.appRoutes.Run(r.handler)
	r.usersRoutes.Run(apiV1)
	r.postsRoutes.Run(apiV1)

	// Not Found route
	r.handler.NoRoute(func(c *gin.Context) {
		utils.ErrorJSON(c, http.StatusNotFound, errors.New("not found"))
	})
}
