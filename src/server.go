package src

import (
	"context"
	"fmt"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"go-photopost/src/modules/posts"
	"go-photopost/src/modules/users"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"gorm.io/gorm"
)

type Server struct {
	Log               *log.Logger
	Env               *lib.Env
	DB                *gorm.DB
	JWTAuthHelper     *lib.JWTAuthHelper
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	AppModule         *AppModule
	UsersModule       *users.UsersModule
	PostsModule       *posts.PostsModule
}

func NewServer(
	log *log.Logger,
	env *lib.Env,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appModule *AppModule,
	usersModule *users.UsersModule,
	postsModule *posts.PostsModule,
) *Server {
	return &Server{
		log,
		env,
		db,
		jwtAuthHelper,
		jwtAuthMiddleware,
		appModule,
		usersModule,
		postsModule,
	}
}

func (server Server) Start() {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// version 1
	apiV1 := r.Group("v1")

	// routes
	server.AppModule.Router(r)
	server.UsersModule.Router(apiV1)
	server.PostsModule.Router(apiV1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": http.StatusNotFound, "message": "Not Found"})
	})

	// r.Run()

	port := ":8080"
	if server.Env.ServerPort != "" {
		port = fmt.Sprintf(":%v", server.Env.ServerPort)
	}
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.Log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		server.Log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		server.Log.Println("timeout of 5 seconds.")
	}
	server.Log.Println("Server exiting")
}
