package controllers

import (
	"github.com/gofiber/fiber/v2"
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

func (p PostsController) UploadPhoto(c *fiber.Ctx) error {
	body, err := utils.BindBody[dto.UploadPhotoDto](c)
	if err != nil {
		return err
	}

	s3, err := p.s3Service.UploadPhoto(&body.Image)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	keypath := (*s3.Key)[4:]
	photo := p.postsService.UploadPhoto(&keypath)

	return utils.SuccessJSON(c, fiber.Map{
		"s3":    s3,
		"photo": photo,
	})

	// return utils.SuccessJSON(c, fiber.Map{})
}

func (p PostsController) CreatePost(c *fiber.Ctx) error {
	body, err := utils.BindBody[dto.CreatePostDto](c)
	if err != nil {
		return err
	}

	user, _ := utils.GetUser[models.User](c)
	trxHandle, _ := c.Locals(constants.DBTransaction).(*gorm.DB)

	result := p.postsService.WithTrx(trxHandle).CreatePost(user, body)
	return utils.SuccessJSON(c, result)
}

func (p PostsController) GetPostList(c *fiber.Ctx) error {
	limit, cursor := utils.GetPaginationCursorQuery(c)
	items, err := p.postsService.SetPaginationScope(utils.PaginateCursor(limit)).GetPostList(cursor)
	if err != nil {
		return err
	}

	resp := utils.CreatePaginationCursor(items, limit, cursor)
	return utils.SuccessJSON(c, resp)
}

func (p PostsController) GetPostById(c *fiber.Ctx) error {
	params, err := utils.BindParams[dto.GetPostByIDParams](c)
	if err != nil {
		return err
	}

	resp, err := p.postsService.GetPostById(params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessJSON(c, resp)
}

func (p PostsController) GetMyPostList(c *fiber.Ctx) error {
	user, _ := utils.GetUser[models.User](c)

	limit, page := utils.GetPaginationQuery(c)
	items, count, err := p.postsService.SetPaginationScope(utils.Paginate(limit, page)).GetMyPostList(user)
	if err != nil {
		return err
	}

	resp := utils.CreatePagination(&items, count, limit, page)
	return utils.SuccessJSON(c, resp)
}

func (p PostsController) GetUserPostList(c *fiber.Ctx) error {
	params, err := utils.BindParams[dto.GetPostByUserIDParams](c)
	if err != nil {
		return err
	}

	limit, page := utils.GetPaginationQuery(c)
	items, count, err := p.postsService.SetPaginationScope(utils.Paginate(limit, page)).GetUserPostList(params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp := utils.CreatePagination(&items, count, limit, page)
	return utils.SuccessJSON(c, resp)
}

func (p PostsController) DeletePost(c *fiber.Ctx) error {
	user, _ := utils.GetUser[models.User](c)

	params, err := utils.BindParams[dto.GetPostByIDParams](c)
	if err != nil {
		return err
	}

	post, err := p.postsService.GetPostById(params)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if *post.AuthorID != user.ID {
		return fiber.NewError(fiber.StatusForbidden, "author_id != user.id")
	}

	p.postsService.DeletePost(&post, user, params)
	return utils.SuccessJSON(c, fiber.Map{})
}
