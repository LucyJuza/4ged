package handlers

import (
	"errors"

	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetGames(c *fiber.Ctx) error {
	id := c.Params("id")

	var games []entities.Game
	var result *gorm.DB

	if id != "" {
		result = config.DB.Find(&games, id)
		if result.RowsAffected == 0 {
			return c.SendStatus(404)
		}
	} else {
		config.DB.Find(&games)
	}

	return c.Status(200).JSON(games)
}

func AddGame(c *fiber.Ctx) error {
	game := new(entities.Game)

	if err := c.BodyParser(game); err != nil {
		return c.Status(400).SendString("Invalid parameters.")
	}

	res, err := addGame(game)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
	}

	return c.JSON(res)
}

func AddGames(c *fiber.Ctx) error {
	var games *[]entities.Game

	if err := c.BodyParser(&games); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	res, err := addGames(games)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
	}

	return c.JSON(res)
}

func addGame(game *entities.Game) (*entities.Game, error) {
	res := config.DB.Create(&game)

	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}

	res = config.DB.Save(&game)

	if res.Error != nil {
		return nil, errors.New("Problem while adding data")
	}

	return game, nil
}

func addGames(games *[]entities.Game) ([]entities.Game, error) {
	res := config.DB.Create(&games)

	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}

	res = config.DB.Save(&games)

	if res.Error != nil {
		return nil, errors.New("Problem while adding data")
	}

	return *games, nil
}
