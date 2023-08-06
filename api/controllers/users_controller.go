package controllers

import (
	"net/http"
	"photopost/api/dto"
	"photopost/api/services"
	"photopost/lib"
	"photopost/utils"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	logger       lib.Logger
	usersService *services.UsersService
}

func NewUsersController(
	logger lib.Logger,
	usersService *services.UsersService,
) *UsersController {
	return &UsersController{
		logger,
		usersService,
	}
}

func (u UsersController) GetUserList(c *gin.Context) {
	result, err := u.usersService.SetPaginationScope(utils.Paginate(c)).GetUserList()
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.JSONWithPagination(c, http.StatusOK, result)
}

func (u UsersController) GetUserByID(c *gin.Context) {
	uri := new(dto.GetUserByIDParams)
	err := c.ShouldBindUri(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.usersService.GetUserByID(uri)
	if err != nil {
		utils.ErrorJSON(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
