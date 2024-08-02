package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssamsara98/photopost-golang/src/constants"
)

func BindBody[T any](c *fiber.Ctx) (*T, error) {
	body := new(T)
	if err := c.BodyParser(body); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return body, nil
}

func BindParams[T any](c *fiber.Ctx) (*T, error) {
	params := new(T)
	if err := c.ParamsParser(params); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return params, nil
}

func GetUser[T any](c *fiber.Ctx) (*T, bool) {
	user, boolean := c.Locals(constants.User).(*T)
	return user, boolean
}
