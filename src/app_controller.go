package src

import (
	"go-photopost/src/entities"
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppControllerInf interface {
	Run(router *gin.Engine)
	Greet(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type AppController struct {
	Log               *log.Logger
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	AppService        AppServiceInf
}

func NewAppController(
	log *log.Logger,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	appService AppServiceInf,
) *AppController {
	return &AppController{
		log,
		jwtAuthMiddleware,
		appService,
	}
}

func (app AppController) Run(router *gin.Engine) {
	router.GET("/", app.Greet)
	router.POST("/register", app.Register)
	router.POST("/login", app.Login)
	router.GET("/me", app.JWTAuthMiddleware.Handler(), app.Me)
}

func (app AppController) Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (app AppController) Register(c *gin.Context) {
	var body RegisterUserDto
	c.Bind(&body)

	user := app.AppService.Register(&body)

	c.JSON(http.StatusCreated, user)
}

func (app AppController) Login(c *gin.Context) {
	var body LoginUserDto
	c.Bind(&body)

	token, err := app.AppService.Login(&body)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Email/Username or Password",
		})
	}

	c.JSON(http.StatusCreated, token)
}

func (app AppController) Me(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*entities.User)

	app.Log.Println(user)

	c.JSON(http.StatusOK, user)
}
