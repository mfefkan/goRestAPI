package main

import (
	"fmt"
	"log"
	"test/db"
	"test/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Tüm domainlerden gelen isteklere izin ver
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)
	app.Put("/users/:id/balance", handler.UpdateBalance)
	app.Put("/users/:id/guess", handler.GuessAndUpdateBalance)
	app.Listen(":8000")
}
