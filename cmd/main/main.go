package main

import (
	"log"
	"net/http"
    "github.com/gofiber/fiber/v2"
	/* "github.com/rs/cors" */ // Import the "rs/cors" package
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kavikkannan/go-jwt/pkg/routes"
    "github.com/kavikkannan/go-jwt/pkg/config"
)

func main() {
    config.Connect()
	app := fiber.New()
	routes.Setup(app)

	/* // Create a CORS handler with your desired options.
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Add your frontend origin here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Define allowed HTTP methods
		AllowedHeaders: []string{"Content-Type"}, // Define allowed headers
	})

	// Use the CORS middleware with your router.
	handler := c.Handler(app) */

	// Start your server with the CORS middleware.
	port := ":9000"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
