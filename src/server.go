package src

import (
	"context"
	"fmt"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// type Server struct {
// 	Log               *log.Logger
// 	Env               *lib.Env
// 	DB                *gorm.DB
// 	JWTAuthHelper     *lib.JWTAuthHelper
// 	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
// 	AppModule         *AppModule
// 	UsersModule       *users.UsersModule
// 	PostsModule       *posts.PostsModule
// }

// func NewServer(
// 	log *log.Logger,
// 	env *lib.Env,
// 	db *gorm.DB,
// 	jwtAuthHelper *lib.JWTAuthHelper,
// 	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
// 	appModule AppModuleInf,
// 	usersModule users.UsersModuleInf,
// 	postsModule posts.PostsModuleInf,
// ) *Server {
// 	return &Server{
// 		log,
// 		env,
// 		db,
// 		jwtAuthHelper,
// 		jwtAuthMiddleware,
// 		appModule,
// 		usersModule,
// 		postsModule,
// 	}
// }

// func (server Server) Start() {
// 	r := gin.Default()
// 	r.Use(favicon.New("./favicon.ico"))
// 	r.Use(gin.Recovery())
// 	r.Use(gin.Logger())

// 	// version 1
// 	apiV1 := r.Group("v1")

// 	// routes
// 	server.AppModule.Router(r)
// 	server.UsersModule.Router(apiV1)
// 	server.PostsModule.Router(apiV1)

// 	r.NoRoute(func(c *gin.Context) {
// 		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
// 	})

// 	// r.Run()

// 	port := ":8080"
// 	if server.Env.ServerPort != "" {
// 		port = fmt.Sprintf(":%v", server.Env.ServerPort)
// 	}
// 	srv := &http.Server{
// 		Addr:    port,
// 		Handler: r,
// 	}

// 	go func() {
// 		// service connections
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s\n", err)
// 		}
// 	}()

// 	// Wait for interrupt signal to gracefully shutdown the server with
// 	// a timeout of 5 seconds.
// 	quit := make(chan os.Signal)
// 	// kill (no param) default send syscanll.SIGTERM
// 	// kill -2 is syscall.SIGINT
// 	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	server.Log.Println("Shutdown Server ...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := srv.Shutdown(ctx); err != nil {
// 		server.Log.Fatal("Server Shutdown:", err)
// 	}
// 	// catching ctx.Done(). timeout of 5 seconds.
// 	select {
// 	case <-ctx.Done():
// 		server.Log.Println("timeout of 5 seconds.")
// 	}
// 	server.Log.Println("Server exiting")
// }

func NewServer(
	lc fx.Lifecycle,
	log *log.Logger,
	env *lib.Env,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appModule *AppModule,
	usersModule *users.UsersModule,
	postsModule *posts.PostsModule,
) *http.Server {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello",
	// 	})
	// })

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
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

var ServerModuleFx = fx.Options(
	fx.Provide(
		NewServer,
		NewAppModule,
		NewAppController,
		NewAppService,
	),
)
