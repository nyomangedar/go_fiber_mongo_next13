package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func ArticleRoutes(app *fiber.App) {
	app.Post("/article", controllers.CreateArticle)
	app.Get("/article", controllers.GetArticle)
	app.Get("/article/:userId", controllers.FindArticle)
}
