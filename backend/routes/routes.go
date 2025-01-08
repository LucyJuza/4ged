package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	// Games routes
	app.Get("/games/:id?", handlers.GetGames)
	// Search games
	app.Post("/games", handlers.AddGame)
	app.Post("/batch/games", handlers.AddGames)

	// Get users (by id)

	// Login -> user

	// Register -> user already exists

	// Post/Delete games for user

	// Post/Delete plays for user

	// Post/Delete players for user

	// (Stats)

	// Genres routes
	app.Get("/genres", handlers.GetGenres)
}
