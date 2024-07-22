package controllers

import (
	"net/http"

	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/api/services"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/helpers"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"
	"github.com/ssamsara98/photopost-golang/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppController struct {
	logger     *lib.Logger
	appService *services.AppService
}

func NewAppController(
	logger *lib.Logger,
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
	body, err := utils.BindBody[dto.RegisterUserDto](c)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	err = app.appService.FindEmailUsername(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusConflict, err)
		return
	}

	trxHandle, _ := c.MustGet(constants.DBTransaction).(*gorm.DB)

	result, err := app.appService.WithTrx(trxHandle).Register(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessJSON(c, http.StatusCreated, result)
}

func (app AppController) Login(c *gin.Context) {
	body, err := utils.BindBody[dto.LoginUserDto](c)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	token, err := app.appService.Login(body)
	if err != nil {
		utils.ErrorJSON(c, http.StatusUnauthorized, err)
		return
	}

	utils.SuccessJSON(c, http.StatusCreated, token)
}

func (app AppController) Me(c *gin.Context) {
	user, _ := c.MustGet(constants.User).(*models.User)
	utils.SuccessJSON(c, http.StatusOK, user)
}

func (app AppController) UpdateProfile(c *gin.Context) {
	body, err := utils.BindBody[dto.UpdateProfileDto](c)
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

func (app AppController) TokenCheck(c *gin.Context) {
	claims, _ := c.MustGet(constants.User).(*helpers.Claims)
	utils.SuccessJSON(c, http.StatusOK, claims)
}

func (app AppController) TokenRefresh(c *gin.Context) {
	user, _ := c.MustGet(constants.User).(*models.User)
	tokens, err := app.appService.TokenRefresh(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusUnauthorized, err)
		return
	}
	utils.SuccessJSON(c, http.StatusOK, tokens)
}
