package handlers

import (
	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
)

func GetGenres(c *fiber.Ctx) error {
	var genres []entities.Genre
	config.DB.Find(&genres)

	return c.Status(200).JSON(genres)
}
