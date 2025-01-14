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
	userID.Delete("games/:gameId<int>", handlers.RemoveUserGame)

	// Post/Delete plays for user
	userID.Post("plays", handlers.AddUserPlay)
	userID.Delete("plays/:playId<int>", handlers.RemoveUserPlay)

	// Post/Delete players for user
	userID.Post("persons", handlers.AddUserPlayer)
	userID.Delete("persons/:personId<int>", handlers.RemoveUserPlayer)

	// (Stats)

	// Genres routes
	app.Get("/genres", handlers.GetGenres)
}
