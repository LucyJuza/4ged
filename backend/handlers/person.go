package handlers

import (
	"strconv"

	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
)

func AddUserPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	player := entities.Player{}

	if err = c.BodyParser(&player); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Bad person format")
	}

	err = config.DB.Model(&user).Association("Players").Append(&player)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error while adding person")
	}
	config.DB.Save(&player)

	return c.JSON(player)
}

func RemoveUserPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	playerId := 0

	if playerId, err = strconv.Atoi(c.Params("personId")); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid person id")
	}

	if uint(playerId) == user.PlayerId {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Cannot delete the users person!")
	}
	player := entities.Player{}

	config.DB.Find(&player, playerId)

	if err = config.DB.Model(&user).Association("Players").Delete(player); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString("An error occured while deleting the person")
	}
	return c.JSON(player)
}
