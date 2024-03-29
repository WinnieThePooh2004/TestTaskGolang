package main

import (
	"TestTask/DataFactories"
	"TestTask/RequestHandlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	factory := &DataFactories.DataFactory{}

	migrationsService, err := factory.GetMigrationsService()

	if err != nil {
		return
	}

	err = migrationsService.EnsureMigrated()

	if err != nil {
		return
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		RequestHandlers.HandleGetRequest(factory, ctx)
		return nil
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		RequestHandlers.HandlePostRequest(factory, ctx)
		return nil
	})

	err = app.Listen(":3000")

	if err != nil {
		return
	}
}
