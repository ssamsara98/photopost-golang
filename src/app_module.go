package src

import (
	"errors"
	"go-photopost/src/entities"
	"go-photopost/src/helpers"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppModuleInf interface {
	Router(r *gin.Engine)
}

type AppModule struct {
	Log               *log.Logger
	DB                *gorm.DB
	JWTAuthHelper     *lib.JWTAuthHelper
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
}

func NewAppModule(
	log *log.Logger,
	db *gorm.DB,
	jwtAuthHelper *lib.JWTAuthHelper,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
) *AppModule {
	return &AppModule{
		log,
		db,
		jwtAuthHelper,
		jwtAuthMiddleware,
	}
}

func (app AppModule) Router(r *gin.Engine) {
	r.GET("/", app.greet)
	r.POST("/register", app.register)
	r.POST("/login", app.login)
	r.GET("/me", app.JWTAuthMiddleware.Handler(), app.me)
}

func (app AppModule) greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (app AppModule) register(c *gin.Context) {
	var body RegisterUserDto
	c.Bind(&body)

	hashedPassword := helpers.HashPassword([]byte(body.Password))

	user := entities.User{
		Email:    &body.Email,
		Username: &body.Username,
		Password: string(hashedPassword),
		Name:     body.Name,
	}
	app.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func (app AppModule) login(c *gin.Context) {
	var body LoginUserDto
	c.Bind(&body)

	var user entities.User
	res := app.DB.Where("email = ? OR username = ?", body.UserSession, body.UserSession).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Email/Username or Password",
		})
		return
	}

	err := helpers.CompareHash([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode": http.StatusUnauthorized,
			"message":    "Email/Username or Password",
		})
		return
	}

	// create token
	token := app.JWTAuthHelper.CreateToken(user)

	c.JSON(http.StatusCreated, token)
}

func (app AppModule) me(c *gin.Context) {
	userAny, _ := c.Get("user")
	user := userAny.(*entities.User)

	app.Log.Println(user)

	c.JSON(http.StatusOK, user)
}
