package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	UpdateBalance(c *fiber.Ctx) error
	GuessAndUpdateBalance(*fiber.Ctx) error
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

func (h handler) UpdateBalance(c *fiber.Ctx) error {
	// Kullanıcı ID'sini al
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(Response{Error: "Invalid ID"})
	}

	// Request body'den eklemek istediğiniz miktarı alın
	var request struct {
		Amount int `json:"amount"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(Response{Error: "Invalid request body"})
	}

	// Servis katmanında hesaplama ve güncelleme işlemini gerçekleştir
	updatedModel, err := h.service.UpdateBalance(uint(id), request.Amount)
	if err != nil {
		// Güncelleme sırasında bir hata oluşursa
		return c.Status(500).JSON(Response{Error: "Failed to update balance"})
	}

	// Başarılı yanıt dön
	return c.Status(200).JSON(Response{Data: updatedModel})
}

func (h handler) GuessAndUpdateBalance(c *fiber.Ctx) error {
	// Kullanıcı ID'sini al
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Error: "Invalid user ID"})
	}

	// Request body'den miktarı ve hedef değeri alın
	var request struct {
		Amount  int   `json:"amount"`
		Targets []int `json:"targets"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Error: "Invalid request body"})
	}

	// Servis katmanında GuessAndUpdateBalance işlemini gerçekleştir
	result, err := h.service.GuessAndUpdateBalance(uint(id), request.Amount, request.Targets)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Error: "Failed to update balance"})
	}

	// Başarılı yanıt dön
	return c.Status(fiber.StatusOK).JSON(Response{Data: result})
}
