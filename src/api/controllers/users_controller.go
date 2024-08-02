package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/api/services"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/utils"
)

type UsersController struct {
	logger       *lib.Logger
	usersService *services.UsersService
}

func NewUsersController(
	logger *lib.Logger,
	usersService *services.UsersService,
) *UsersController {
	return &UsersController{
		logger,
		usersService,
	}
}

func (u UsersController) GetUserList(c *fiber.Ctx) error {
	limit, page := utils.GetPaginationQuery(c)
	items, count, err := u.usersService.SetPaginationScope(utils.Paginate(limit, page)).GetUserList()
	if err != nil {
		return err
	}

	resp := utils.CreatePagination(items, count, limit, page)
	return utils.SuccessJSON(c, resp)
}
func (u UsersController) GetUserListCursor(c *fiber.Ctx) error {
	limit, cursor := utils.GetPaginationCursorQuery(c)
	items, err := u.usersService.SetPaginationScope(utils.PaginateCursor(limit)).GetUserListCursor(cursor)
	if err != nil {
		return err
	}

	resp := utils.CreatePaginationCursor(items, limit, cursor)
	return utils.SuccessJSON(c, resp)
}

func (u UsersController) GetUserByID(c *fiber.Ctx) error {
	params, err := utils.BindParams[dto.GetUserByIDParams](c)
	if err != nil {
		return err
	}

	user, err := u.usersService.GetUserByID(params)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessJSON(c, user)
}
