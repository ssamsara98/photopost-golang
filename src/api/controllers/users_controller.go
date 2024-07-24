package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (u UsersController) GetUserList(c *gin.Context) {
	limit, page := utils.GetPaginationQuery(c)
	items, count, err := u.usersService.SetPaginationScope(utils.Paginate(limit, page)).GetUserList()
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	resp := utils.CreatePagination(items, count, limit, page)
	utils.SuccessJSON(c, http.StatusOK, resp)
}
func (u UsersController) GetUserListCursor(c *gin.Context) {
	limit, cursor := utils.GetPaginationCursorQuery(c)
	items, err := u.usersService.SetPaginationScope(utils.PaginateCursor(limit)).GetUserListCursor(cursor)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	resp := utils.CreatePaginationCursor(items, limit, cursor)
	utils.SuccessJSON(c, http.StatusOK, resp)
}

func (u UsersController) GetUserByID(c *gin.Context) {
	uri := utils.BindUri[dto.GetUserByIDParams](c)
	if uri == nil {
		return
	}

	user, err := u.usersService.GetUserByID(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusNotFound, err)
		return
	}

	utils.SuccessJSON(c, http.StatusOK, user)
}
