package rest

import "github.com/gofiber/fiber/v2"

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) Route(app *fiber.App) {
	app.Get("api/orders", c.findAll)
	app.Get("api/orders/:id", c.findByID)
	app.Post("api/orders", c.create)
	app.Delete("api/orders/:id", c.cancelOrder)
}

func (c *OrderController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) cancelOrder(ctx *fiber.Ctx) error {
	return nil
}
