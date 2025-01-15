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
	playJson := entities.PlayJson{}

	if err = c.BodyParser(&playJson); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Bad play format")
	}
	play := playFromPlayJson(playJson)

	game := entities.Game{}
	config.DB.Find(&game, play.GameId)

	if game.ID != play.GameId {
		return c.Status(fiber.ErrNotFound.Code).SendString("No such game")
	}

	err = config.DB.Model(&user).Association("Games").Append(&game)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error while adding play")
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

func playFromPlayJson(play entities.PlayJson) entities.Play {
	players := make([]entities.Player, 0)
	winners := make([]entities.Player, 0)

	config.DB.Find(&players, play.Players)
	config.DB.Find(&winners, play.Winners)

	return entities.Play{
		GameId:   play.GameId,
		Date:     play.Date,
		Location: play.Location,
		Duration: play.Duration,
		Players:  players,
		Winners:  winners,
	}
}

func playJsonFromPlay(play entities.Play) entities.PlayJson {
	players := make([]uint, 0)
	winners := make([]uint, 0)

	for _, player := range play.Players {
		players = append(players, player.ID)
	}
	for _, winner := range play.Winners {
		winners = append(winners, winner.ID)
	}

	return entities.PlayJson{
		GameId:   play.GameId,
		Date:     play.Date,
		Location: play.Location,
		Duration: play.Duration,
		Players:  players,
		Winners:  winners,
	}
}
