package main

import (
	"github.com/Saatvik-droid/url-shortner/model"
	"github.com/Saatvik-droid/url-shortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupServer() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetupRoutes(router.Group("/"))

	model.SetupDatabase()

	router.Listen(":3000")
}
