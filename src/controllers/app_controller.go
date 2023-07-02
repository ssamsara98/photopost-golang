package controllers

import (
	"net/http"
	"photopost/constants"
	"photopost/lib"
	"photopost/models"
	"photopost/src/dto"
	"photopost/src/services"
	"photopost/utils"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	logger     lib.Logger
	appService *services.AppService
}

func NewAppController(
	logger lib.Logger,
	appService *services.AppService,
) *AppController {
	return &AppController{
		logger,
		appService,
	}
}

func (app AppController) Home(c *gin.Context) {
	message := app.appService.Home()
	utils.SuccessJSON(c, http.StatusOK, message)
}

func (app AppController) Register(c *gin.Context) {
	body := new(dto.RegisterUserDto)
	err := c.Bind(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	err = app.appService.FindEmailUsername(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusConflict, err)
		return
	}

	user, err := app.appService.Register(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (app AppController) Login(c *gin.Context) {
	body := new(dto.LoginUserDto)
	err := c.Bind(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	token, err := app.appService.Login(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func (app AppController) Me(c *gin.Context) {
	user, _ := c.MustGet(constants.User).(*models.User)

	c.JSON(http.StatusOK, user)
}

func (app AppController) UpdateProfile(c *gin.Context) {
	body := new(dto.UpdateProfile)
	err := c.Bind(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	user, _ := c.MustGet(constants.User).(*models.User)
	err = app.appService.UpdateProfile(user.ID, body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessJSON(c, http.StatusOK, "success")
}
