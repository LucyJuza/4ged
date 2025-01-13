package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	// Games routes
	app.Get("/games/:id<int>?", handlers.GetGames)
	app.Get("/games/search", handlers.SearchGames)
	// Search games
	app.Post("/games", handlers.AddGame)
	app.Post("/batch/games", handlers.AddGames)

	// Get users (by id)
	app.Get("/users/:id<int>?", handlers.GetUsers)

	// Login -> user
	app.Post("/login", handlers.Login)

	// Register -> user already exists
	app.Post("/register", handlers.Register)

	userID := app.Group("/users/:id<int>/")

	// Post/Delete games for user
	userID.Post("games", handlers.AddUserGame)
	userID.Delete("games", handlers.AddUserGame)

	// userID.Delete("games/:id", handlers)

	// Post/Delete plays for user

	// Post/Delete players for user

	// (Stats)

	// Genres routes
	app.Get("/genres", handlers.GetGenres)
}
