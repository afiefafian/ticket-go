package rest

import "github.com/gofiber/fiber/v2"

type PaymentController struct {
}

func NewPaymentController() *PaymentController {
	return &PaymentController{}
}

func (c *PaymentController) Route(app *fiber.App) {
	app.Post("api/payment/notification", c.notification)
}

func (c *PaymentController) notification(ctx *fiber.Ctx) error {
	return nil
}
