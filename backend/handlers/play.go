package handlers

import (
	"strconv"

	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
)

func AddUserPlay(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	play := entities.Play{}

	if err = c.BodyParser(&play); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Bad play format")
	}

	err = config.DB.Model(&user).Association("Plays").Append(&play)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error while adding play")
	}

	config.DB.Save(&play)

	return c.JSON(play)
}

func RemoveUserPlay(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	playId := 0

	if playId, err = strconv.Atoi(c.Params("playId")); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid play id")
	}

	play := entities.Play{}

	config.DB.Find(&play, playId)

	if err = config.DB.Model(&user).Association("Plays").Delete(play); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString("An error occured while deleting the person")
	}
	return c.JSON(play)
}
