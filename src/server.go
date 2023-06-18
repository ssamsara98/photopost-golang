package src

import (
	"context"
	"fmt"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/app"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewServer(
	lc fx.Lifecycle,
	log *zap.Logger,
	env *lib.Env,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appModule app.AppModuleInf,
	usersModule users.UsersModuleInf,
	postsModule posts.PostsModuleInf,
) *http.Server {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// version 1
	apiV1 := r.Group("v1")

	// routes
	appModule.Router(r)
	usersModule.Router(apiV1)
	postsModule.Router(apiV1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
	})

	port := ":8080"
	if env.ServerPort != "" {
		port = fmt.Sprintf(":%v", env.ServerPort)
	}
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = fmt.Sprintf(":%v", envPort)
	}
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Sugar().Infof("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Sugar().Infoln("Shutting down the server")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

var ServerModuleFx = fx.Options(
	fx.Provide(
		NewServer,
		zap.NewExample,
	),
)
