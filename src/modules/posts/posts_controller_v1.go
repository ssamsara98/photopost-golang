package posts

import (
	"fmt"
	"go-photopost/src/entities"
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsControllerV1Interface interface {
	Run(router *gin.RouterGroup)
	UploadPhoto(c *gin.Context)
	CreatePost(c *gin.Context)
	GetPostList(c *gin.Context)
	GetPost(c *gin.Context)
}

type PostsControllerV1 struct {
	Log               *log.Logger
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	PostsServiceV1    PostsServiceV1Interface
}

func NewPostsControllerV1(
	log *log.Logger,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	postsServiceV1 *PostsServiceV1,
) *PostsControllerV1 {
	return &PostsControllerV1{
		log,
		jwtAuthMiddleware,
		postsServiceV1,
	}
}

func (pc PostsControllerV1) Run(router *gin.RouterGroup) {
	router.POST("/", pc.JWTAuthMiddleware.Handler(), pc.CreatePost)
	router.POST("/upload", pc.JWTAuthMiddleware.Handler(), pc.UploadPhoto)
	router.GET("/", pc.GetPostList)
	router.GET("/p/:id", pc.GetPost)
}

func (pc PostsControllerV1) UploadPhoto(c *gin.Context) {
	var body UploadPhotoDto
	c.Bind(&body)

	file, err := body.Image.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()
	fileBuffer := make([]byte, body.Image.Size)
	file.Read(fileBuffer)

	// c.String(http.StatusOK, "", body.Image)
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", body.Image.Header.Get("Content-Type"))
	c.Header("Content-Length", fmt.Sprintf("%d", body.Image.Size))
	c.Writer.Write(fileBuffer) //the memory take up 1.2~1.7G
}

func (pc PostsControllerV1) CreatePost(c *gin.Context) {
	var body CreatePostDto
	c.Bind(&body)

	userAny, _ := c.Get("user")
	user := userAny.(*entities.User)

	result := pc.PostsServiceV1.CreatePost(user, &body)
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPostList(c *gin.Context) {
	result := pc.PostsServiceV1.GetPostList()
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPost(c *gin.Context) {
	var uri GetPostByIdUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := pc.PostsServiceV1.GetPost(&uri)
	c.JSON(http.StatusOK, result)
}
