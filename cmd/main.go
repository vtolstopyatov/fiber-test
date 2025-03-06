package main

import (
	"fmt"
	"log"

	"fibertest/internal/config"
	"fibertest/internal/db"
	"fibertest/internal/handler"
	"fibertest/internal/seed_test_data"
	"fibertest/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/list", handler.GetNewsList)
	app.Post("/edit/:id", middleware.Auth(), handler.UpdateNews)
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
		ErrorHandler: handler.ErrorHandler,
	})

	setupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", cfg.HTTPPort)))
}
