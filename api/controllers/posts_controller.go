package controllers

import (
	"net/http"
	"photopost/api/dto"
	"photopost/api/services"
	"photopost/constants"
	"photopost/lib"
	"photopost/models"
	"photopost/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsController struct {
	logger       lib.Logger
	s3Service    *services.S3Service
	postsService *services.PostsService
}

func NewPostsController(
	logger lib.Logger,
	s3Service *services.S3Service,
	postsService *services.PostsService,
) *PostsController {
	return &PostsController{
		logger,
		s3Service,
		postsService,
	}
}

func (p PostsController) UploadPhoto(c *gin.Context) {
	body := new(dto.UploadPhotoReqDto)
	err := c.Bind(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	s3, err := p.s3Service.UploadPhoto(&body.Image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	keypath := (*s3.Key)[4:]
	photo := p.postsService.UploadPhoto(&keypath)

	c.JSON(http.StatusOK, gin.H{
		"s3":    s3,
		"photo": photo,
	})
}

func (p PostsController) CreatePost(c *gin.Context) {
	body := new(dto.CreatePostReqDto)
	err := c.Bind(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	user, _ := c.MustGet(constants.User).(*models.User)
	trxHandle, _ := c.MustGet(constants.DBTransaction).(*gorm.DB)

	result := p.postsService.WithTrx(trxHandle).CreatePost(user, body)
	c.JSON(http.StatusOK, result)
}

func (p PostsController) GetPostList(c *gin.Context) {
	result, err := p.postsService.SetPaginationScope(utils.Paginate(c)).GetPostList()
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.JSONWithPagination(c, http.StatusOK, result)
}

func (p PostsController) GetPost(c *gin.Context) {
	uri := new(dto.GetPostByIDParams)
	err := c.ShouldBindUri(uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	result := p.postsService.GetPost(uri)
	c.JSON(http.StatusOK, result)
}

func (p PostsController) GetMyPostList(c *gin.Context) {
	user, _ := c.MustGet(constants.User).(*models.User)

	result, err := p.postsService.SetPaginationScope(utils.Paginate(c)).GetMyPostList(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}
	utils.JSONWithPagination(c, http.StatusOK, result)
}

func (p PostsController) GetUserPostList(c *gin.Context) {
	uri := new(dto.GetPostByUserIDParams)
	err := c.ShouldBindUri(uri)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	// result := p.postsService.GetUserPostList(uri)
	result, err := p.postsService.SetPaginationScope(utils.Paginate(c)).GetUserPostList(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}
	utils.JSONWithPagination(c, http.StatusOK, result)
}
