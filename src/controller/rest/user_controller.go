package rest

import "github.com/gofiber/fiber/v2"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Route(app *fiber.App) {
	app.Post("api/users/me", c.userInfo)
}

func (c *UserController) userInfo(ctx *fiber.Ctx) error {
	return nil
}
