package users

import (
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersControllerV1Inf interface {
	Run(router *gin.RouterGroup)
}

type UsersControllerV1 struct {
	Log               *log.Logger
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	UsersService      UsersServiceV1Inf
}

func NewUsersControllerV1(
	log *log.Logger,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	usersService *UsersServiceV1,
) *UsersControllerV1 {
	return &UsersControllerV1{
		log,
		jwtAuthMiddleware,
		usersService,
	}
}

func (uc UsersControllerV1) Run(router *gin.RouterGroup) {
	router.POST("/", uc.CreateUser)
	router.GET("/", uc.GetUserList)
	router.GET("/u/:userId", uc.GetUser)
}

func (uc UsersControllerV1) CreateUser(c *gin.Context) {
	var body CreateUserReqDto
	c.Bind(&body)

	result := uc.UsersService.CreateUser(body)

	c.JSON(http.StatusCreated, result)
}

func (uc UsersControllerV1) GetUserList(c *gin.Context) {
	result := uc.UsersService.GetUserList()
	c.JSON(http.StatusOK, result)
}

func (uc UsersControllerV1) GetUser(c *gin.Context) {
	var uri GetUserByIdParams
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := uc.UsersService.GetUser(&uri)
	c.JSON(http.StatusOK, result)
}
