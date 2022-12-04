package routes

import (
	"github.com/Saatvik-droid/url-shortner/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	router.Post("/create", controller.CreateUrl)
	router.Get("/:redirect", controller.Redirect)
}
