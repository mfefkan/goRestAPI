package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(*fiber.Ctx) error
	Create(*fiber.Ctx) error
}

type handler struct {
	service Service
}

var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func (h handler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	model, err := h.service.Get(uint(id))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	return c.Status(200).JSON(Response{Data: model})

}

func (h handler) Create(c *fiber.Ctx) error {
	model := Model{}

	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	_, err = h.service.Create(model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	return c.SendStatus(201)
}
