package user

import (
	"fmt"
	"net/http/httptest"
	"test/db"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err) // err nil değilse error f fonksiyonunu çağırıyor
	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)
	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	id, err := repo.Create(Model{Name: "test1", Email: "test1@mail.com"})
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
