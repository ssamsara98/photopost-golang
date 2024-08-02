package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssamsara98/photopost-golang/src/api/dto"
	"github.com/ssamsara98/photopost-golang/src/api/services"
	"github.com/ssamsara98/photopost-golang/src/constants"
	"github.com/ssamsara98/photopost-golang/src/helpers"
	"github.com/ssamsara98/photopost-golang/src/lib"
	"github.com/ssamsara98/photopost-golang/src/models"
	"github.com/ssamsara98/photopost-golang/src/utils"
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

func (app AppController) Home(c *fiber.Ctx) error {
	message := app.appService.Home()
	return utils.SuccessJSON(c, message)
}

func (app AppController) Register(c *fiber.Ctx) error {
	body, err := utils.BindBody[dto.RegisterUserDto](c)
	if err != nil {
		return err
	}

	if err := app.appService.FindEmailUsername(body); err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	trxHandle, _ := c.Locals(constants.DBTransaction).(*gorm.DB)

	result, err := app.appService.WithTrx(trxHandle).Register(body)
	if err != nil {
		return err
	}

	return utils.SuccessJSON(c, result)
}

func (app AppController) Login(c *fiber.Ctx) error {
	body, err := utils.BindBody[dto.LoginUserDto](c)
	if err != nil {
		return err
	}

	token, err := app.appService.Login(body)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return utils.SuccessJSON(c, token)
}

func (app AppController) Me(c *fiber.Ctx) error {
	user, _ := utils.GetUser[models.User](c)
	return utils.SuccessJSON(c, user)
}

func (app AppController) UpdateProfile(c *fiber.Ctx) error {
	body, err := utils.BindBody[dto.UpdateProfileDto](c)
	if err != nil {
		return err
	}

	user, _ := utils.GetUser[models.User](c)
	if err := app.appService.UpdateProfile(user.ID, body); err != nil {
		return err
	}

	return utils.SuccessJSON(c, "success")
}

func (app AppController) TokenCheck(c *fiber.Ctx) error {
	claims, _ := utils.GetUser[helpers.Claims](c)
	return utils.SuccessJSON(c, claims)
}

func (app AppController) TokenRefresh(c *fiber.Ctx) error {
	user, _ := utils.GetUser[models.User](c)
	tokens, err := app.appService.TokenRefresh(user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return utils.SuccessJSON(c, tokens)
}
