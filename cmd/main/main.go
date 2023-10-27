package main

import (
	/* "log"
	"net/http" */
    "github.com/gofiber/fiber/v2"
	/* "github.com/rs/cors" */ // Import the "rs/cors" package
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/BalkanID-University/go-jwt/pkg/routes"
    "github.com/BalkanID-University/go-jwt/pkg/config"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    config.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	 // Add your frontend origin here


	}))
	routes.Setup(app)

	
	app.Listen(":9000")
}
/* // Create a CORS handler with your desired options.
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Add your frontend origin here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Define allowed HTTP methods
		AllowedHeaders: []string{"Content-Type"}, // Define allowed headers
	})

	// Use the CORS middleware with your router.
	handler := c.Handler(app) 

	// Start your server with the CORS middleware.
	port := ":9000"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe("local"))*/