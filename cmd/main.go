package main

import (
	"fmt"
	"log"
	"test/db"
	"test/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		fmt.Println("Database connection error:", err)
	}

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)

	app.Listen(":8000")
}
