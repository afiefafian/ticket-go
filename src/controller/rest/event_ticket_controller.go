package rest

import "github.com/gofiber/fiber/v2"

type EventTicketController struct {
}

func NewEventTicketController() *EventTicketController {
	return &EventTicketController{}
}

func (c *EventTicketController) Route(app *fiber.App) {
	app.Get("api/events/:eventId", c.findAll)
	app.Get("api/events/:eventId/tickets/:id", c.findByID)
	app.Post("api/events/:eventId/tickets/:id", c.create)
	app.Put("api/events/:eventId/tickets/:id", c.updateByID)
	app.Delete("api/events/:eventId/tickets/:id", c.deleteByID)
}

func (c *EventTicketController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventTicketController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventTicketController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventTicketController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventTicketController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}
