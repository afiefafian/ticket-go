package rest

import "github.com/gofiber/fiber/v2"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Route(app *fiber.App) {
	app.Post("api/authentication", c.auth)
}

func (c *AuthController) auth(ctx *fiber.Ctx) error {
	return nil
}
