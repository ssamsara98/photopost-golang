package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/utils"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter(
	env *lib.Env,
	logger *lib.Logger,
) *Router {

	// if (env.Environment != "local" && env.Environment != "development") && env.SentryDSN != "" {
	// 	if err := sentry.Init(sentry.ClientOptions{
	// 		Dsn:         env.SentryDSN,
	// 		Environment: `clean-backend-` + env.Environment,
	// 	}); err != nil {
	// 		logger.Infof("Sentry initialization failed: %v\n", err)
	// 	}
	// }

	gin.DefaultWriter = logger.GetGinLogger()
	appEnv := env.Environment
	if appEnv == production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.Default()

	httpRouter.MaxMultipartMemory = env.MaxMultipartMemory

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	httpRouter.Use(gin.Recovery())
	// httpRouter.Use(gin.Logger())

	// // Attach sentry middleware
	// httpRouter.Use(sentrygin.New(sentrygin.Options{
	// 	Repanic: true,
	// }))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		utils.SuccessJSON(c, http.StatusOK, "clean architecture 📺 API Up and Running")
	})

	router := &Router{
		httpRouter,
	}
	return router
}
