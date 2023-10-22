package routes

import (
	"github.com/kavikkannan/akg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {
	app.Get("/", controllers.Hello)
	/* router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
 */}