package handlers

import (
	"errors"
	"log"
	"strconv"

	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const pageSize = 50

func GetGames(c *fiber.Ctx) error {
	id := c.Params("id")

	pageIndex, err := getPageIndex(c)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	var games []entities.Game

	if id != "" {
		err := getGames().Find(&games, id).Error
		if err != nil {
			return c.SendStatus(404)
		}
	} else {
		err := paginate(config.DB, pageIndex).Find(&games).Error
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
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

func SearchGames(c *fiber.Ctx) error {
	var games []entities.Game

	filter := c.Query("filter", "")

	pageIndex, err := getPageIndex(c)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	log.Println(filter)

	paginate(getGames().Where("name LIKE ?", "%"+filter+"%").Find(&games), pageIndex)

	return c.JSON(games)
}

func addGame(game *entities.Game) (*entities.Game, error) {
	res := config.DB.Create(&game)

	if res.Error != nil {
		return nil, errors.New(res.Error.Error())
	}

	res = config.DB.Save(&game)

	if res.Error != nil {
		return nil, errors.New("problem while adding data")
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
		return nil, errors.New("problem while adding data")
	}

	return *games, nil
}

func getGames() *gorm.DB {
	return config.DB.Model(&entities.Game{}).Preload("Genres")
}

func paginate(req *gorm.DB, pageIndex int) *gorm.DB {
	return req.Limit(pageSize).Offset(pageSize * pageIndex)
}

func getPageIndex(c *fiber.Ctx) (int, error) {
	pageIndex, err := strconv.Atoi(c.Query("pageIndex", "0"))
	if err != nil {
		return 0, errors.New("pageIndex should be a positive integer")
	}

	return pageIndex, nil
}
