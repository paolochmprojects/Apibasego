package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func newServer() *fiber.App {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	return app
}
