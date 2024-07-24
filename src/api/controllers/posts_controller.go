package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/api/services"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"
	"github.com/ssamsara98/photopost-golang/src/utils"
	"gorm.io/gorm"
)

type PostsController struct {
	logger       *lib.Logger
	s3Service    *services.S3Service
	postsService *services.PostsService
}

func NewPostsController(
	logger *lib.Logger,
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
	body := utils.BindBody[dto.UploadPhotoDto](c)
	if body == nil {
		return
	}

	s3, err := p.s3Service.UploadPhoto(&body.Image)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	keypath := (*s3.Key)[4:]
	photo := p.postsService.UploadPhoto(&keypath)

	utils.SuccessJSON(c, http.StatusOK, gin.H{
		"s3":    s3,
		"photo": photo,
	})

	// utils.SuccessJSON(c, http.StatusOK, gin.H{})
}

func (p PostsController) CreatePost(c *gin.Context) {
	body := utils.BindBody[dto.CreatePostDto](c)
	if body == nil {
		return
	}

	user, _ := c.MustGet(constants.User).(*models.User)
	trxHandle, _ := c.MustGet(constants.DBTransaction).(*gorm.DB)

	result := p.postsService.WithTrx(trxHandle).CreatePost(user, body)
	utils.SuccessJSON(c, http.StatusOK, result)
}

func (p PostsController) GetPostList(c *gin.Context) {
	limit, cursor := utils.GetPaginationCursorQuery(c)
	items, err := p.postsService.SetPaginationScope(utils.PaginateCursor(limit)).GetPostList(cursor)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	resp := utils.CreatePaginationCursor(items, limit, cursor)
	utils.SuccessJSON(c, http.StatusOK, resp)
}

func (p PostsController) GetPostById(c *gin.Context) {
	uri := utils.BindUri[dto.GetPostByIDParams](c)
	if uri == nil {
		return
	}

	resp, err := p.postsService.GetPostById(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	utils.SuccessJSON(c, http.StatusOK, resp)
}

func (p PostsController) GetMyPostList(c *gin.Context) {
	user, _ := utils.GetUser[models.User](c)

	limit, page := utils.GetPaginationQuery(c)
	items, count, err := p.postsService.SetPaginationScope(utils.Paginate(limit, page)).GetMyPostList(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	resp := utils.CreatePagination(&items, count, limit, page)
	utils.SuccessJSON(c, http.StatusOK, resp)
}

func (p PostsController) GetUserPostList(c *gin.Context) {
	uri := utils.BindUri[dto.GetPostByUserIDParams](c)
	if uri == nil {
		return
	}

	limit, page := utils.GetPaginationQuery(c)
	items, count, err := p.postsService.SetPaginationScope(utils.Paginate(limit, page)).GetUserPostList(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
	}

	resp := utils.CreatePagination(&items, count, limit, page)
	utils.SuccessJSON(c, http.StatusOK, resp)
}

func (p PostsController) DeletePost(c *gin.Context) {
	user, _ := utils.GetUser[models.User](c)

	uri := utils.BindUri[dto.GetPostByIDParams](c)
	if uri == nil {
		return
	}

	post, err := p.postsService.GetPostById(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusNotFound, err)
		return
	}
	if post.AuthorID != &user.ID {
		utils.ErrorJSON(c, http.StatusForbidden, errors.New("author_id != user.id"))
		return
	}

	p.postsService.DeletePost(&post, user, uri)
	utils.SuccessJSON(c, http.StatusNoContent, gin.H{})
}
