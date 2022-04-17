package rest

import "github.com/gofiber/fiber/v2"

type EventController struct {
}

func NewEventController() *EventController {
	return &EventController{}
}

func (c *EventController) Route(app *fiber.App) {
	app.Get("api/events", c.findAll)
	app.Get("api/events/:id", c.findByID)
	app.Post("api/events/:id", c.create)
	app.Put("api/events/:id", c.updateByID)
	app.Delete("api/events/:id", c.deleteByID)
}

func (c *EventController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}
