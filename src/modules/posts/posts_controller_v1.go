package posts

import (
	"go-photopost/src/entities"
	"go-photopost/src/lib"
	"go-photopost/src/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsControllerV1Inf interface {
	Run(router *gin.RouterGroup)
	UploadPhoto(c *gin.Context)
	CreatePost(c *gin.Context)
	GetPostList(c *gin.Context)
	GetPost(c *gin.Context)
}

type PostsControllerV1 struct {
	Log               *log.Logger
	S3Service         *lib.S3Service
	JWTAuthMiddleware *middlewares.JWTAuthMiddleware
	PostsServiceV1    PostsServiceV1Inf
}

func NewPostsControllerV1(
	log *log.Logger,
	s3Service *lib.S3Service,
	jwtAuthMiddleware *middlewares.JWTAuthMiddleware,
	postsServiceV1 PostsServiceV1Inf,
) *PostsControllerV1 {
	return &PostsControllerV1{
		log,
		s3Service,
		jwtAuthMiddleware,
		postsServiceV1,
	}
}

func (pc PostsControllerV1) Run(router *gin.RouterGroup) {
	router.POST("/", pc.JWTAuthMiddleware.Handler(), pc.CreatePost)
	router.POST("/upload", pc.JWTAuthMiddleware.Handler(), pc.UploadPhoto)
	router.GET("/", pc.GetPostList)
	router.GET("/mine", pc.JWTAuthMiddleware.Handler(), pc.GetMyPostList)
	router.GET("/p/:id", pc.GetPost)
	router.GET("/u/:id", pc.GetUserPostList)
}

func (pc PostsControllerV1) UploadPhoto(c *gin.Context) {
	var body UploadPhotoReqDto
	c.Bind(&body)

	// file, err := body.Image.Open()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// defer file.Close()
	// fileBuffer := make([]byte, body.Image.Size)
	// file.Read(fileBuffer)

	// // c.String(http.StatusOK, "", body.Image)

	// c.Writer.WriteHeader(http.StatusOK)
	// c.Header("Content-Type", body.Image.Header.Get("Content-Type"))
	// c.Header("Content-Length", fmt.Sprintf("%d", body.Image.Size))
	// c.Writer.Write(fileBuffer) //the memory take up 1.2~1.7G

	// creds := credentials.NewStaticCredentialsProvider(pc.S3Service.Env.AWSAccessKeyId, pc.S3Service.Env.AWSSecretAccessKey, "")
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(pc.S3Service.Env.AWSRegion))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// client := s3.NewFromConfig(cfg)
	// uploader := manager.NewUploader(client)

	// result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket:        aws.String(pc.S3Service.Env.AWSS3Bucket),
	// 	Key:           aws.String(fmt.Sprintf("img/photopost/%s", body.Image.Filename)),
	// 	Body:          file,
	// 	ACL:           types.ObjectCannedACLPublicRead,
	// 	ContentType:   aws.String(body.Image.Header.Get("Content-Type")),
	// 	ContentLength: body.Image.Size,
	// 	// ContentDisposition:   aws.String("attachment"),
	// 	// ServerSideEncryption: types.ServerSideEncryptionAes256,
	// })
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	s3, err := pc.S3Service.UploadPhoto(&body.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	keypath := (*s3.Key)[4:]
	photo := pc.PostsServiceV1.UploadPhoto(&keypath)

	c.JSON(http.StatusOK, gin.H{
		"s3":    s3,
		"photo": photo,
	})
}

func (pc PostsControllerV1) CreatePost(c *gin.Context) {
	var body CreatePostReqDto
	c.Bind(&body)

	user := c.MustGet("user").(*entities.User)

	result := pc.PostsServiceV1.CreatePost(user, &body)
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPostList(c *gin.Context) {
	result := pc.PostsServiceV1.GetPostList()
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetPost(c *gin.Context) {
	var uri GetPostByIdParams
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := pc.PostsServiceV1.GetPost(&uri)
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetMyPostList(c *gin.Context) {
	user := c.MustGet("user").(*entities.User)

	result := pc.PostsServiceV1.GetMyPostList(user)
	c.JSON(http.StatusOK, result)
}

func (pc PostsControllerV1) GetUserPostList(c *gin.Context) {
	var uri GetPostByUserIdParams
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := pc.PostsServiceV1.GetUserPostList(&uri)
	c.JSON(http.StatusOK, result)
}
