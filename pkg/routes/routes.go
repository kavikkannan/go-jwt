package routes

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/BalkanID-University/go-jwt/pkg/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}