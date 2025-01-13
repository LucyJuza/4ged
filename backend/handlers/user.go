package handlers

import (
	"errors"
	"strconv"

	"backend/config"
	"backend/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	id := c.Params("id")

	pageIndex, err := getPageIndex(c)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid parameters")
	}

	var users []entities.User

	if id != "" {
		err := getUsers().Find(&users, id).Error
		if err != nil {
			return c.SendStatus(404)
		}
	} else {
		err := paginate(getUsers().Find(&users), pageIndex).Error
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}

	return c.Status(200).JSON(users)
}

func Login(c *fiber.Ctx) error {
	credentials := entities.UserCredentials{}

	if err := c.BodyParser(&credentials); err != nil {
		return c.SendString(err.Error())
	}

	log.Info(credentials)
	user := entities.User{Name: credentials.Name, Password: credentials.Password}

	if credentials.Name == "" {
		return loginError(c)
	}

	getUsers().Where("name = ?", credentials.Name).First(&user)

	if user.Name == "" || user.Password != credentials.Password {
		return loginError(c)
	}

	return c.JSON(user.ID)
}

func Register(c *fiber.Ctx) error {
	credentials := entities.UserCredentials{}

	if err := c.BodyParser(&credentials); err != nil ||
		credentials.Name == "" ||
		credentials.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid parameters")
	}

	user := entities.User{Name: credentials.Name, Password: credentials.Password}

	if res := config.DB.Create(&user); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.ErrBadRequest.Code).SendString("Username already exists")
		}
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	if res := config.DB.Save(&user); res.Error != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	return c.JSON(user.ID)
}

func AddUserGame(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	game := entities.UGame{UserID: user.ID}

	if err = c.BodyParser(&game); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Bad game format")
	}

	err = config.DB.Model(&user).Association("UserGames").Append(&game)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error while adding game")
	}
	config.DB.Save(&game)

	return c.JSON(user)
}

func RemoveUserGame(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := getUser(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	ugame := entities.UGame{}

	if err = c.BodyParser(&ugame); err == nil {
		if err = config.DB.Model(&user).Association("UserGames").Delete(ugame); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString("An error occured while deleting the game")
		}

		return c.JSON(ugame)
	}

	game := entities.Game{}

	if err = c.BodyParser(&game); err != nil {
		return c.Status(fiber.ErrNotFound.Code).SendString("No such game")
	}

	if err = config.DB.Model(&user).Association("Games").Delete(&game); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString("An error occured while deleting the game")
	}

	return c.JSON(game)
}

func loginError(c *fiber.Ctx) error {
	return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid username or password")
}

func getUsers() *gorm.DB {
	return config.DB.
		Model(&entities.User{}).
		Preload("Games").
		Preload("UserGames").
		Preload("Plays").
		Preload("Players")
}

func getUser(id string) (*entities.User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	user := entities.User{}
	res := getUsers().Find(&user, userId)
	if res.Error != nil || user.Name == "" {
		return nil, errors.New("no such user")
	}
	return &user, nil
}
