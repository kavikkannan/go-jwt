package main

import (
	/* "log"
	"net/http" */
    "github.com/gofiber/fiber/v2"
	/* "github.com/rs/cors" */ // Import the "rs/cors" package
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kavikkannan/go-jwt/pkg/routes"
    "github.com/kavikkannan/go-jwt/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    config.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:"*",

	}))
	routes.Setup(app)

	
	app.Listen(":9000")
}
