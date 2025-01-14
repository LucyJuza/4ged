package main

import (
	"log"

	"backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	err := config.Connect()
	if err != nil {
		log.Fatal("Couldn't initialize configuration")
		return
	}

	routes.Init(app)
	log.Fatal(app.Listen(":8080"))
}
