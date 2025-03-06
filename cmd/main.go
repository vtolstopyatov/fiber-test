package main

import (
	"errors"
	"fmt"
	"log"

	"fibertest/internal/config"
	"fibertest/internal/db"
	"fibertest/internal/handler"
	"fibertest/internal/seed_test_data"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/list", handler.GetNewsList)
	app.Post("/edit/:id", handler.UpdateNews)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db.ConnectDb()
	defer db.SqlDb.Close()

	cfg := config.GetConfig()
	if cfg.Environment == "development" {
		seed_test_data.SeedData()
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			ctx.Status(code).JSON(fiber.Map{
				"Success": false,
				"Error":   err.Error(),
			})

			return nil
		},
	})
	setupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", cfg.HTTPPort)))
}
