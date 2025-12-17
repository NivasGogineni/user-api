package main

import (
	"user-api/config"
	db "user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// DB connection
	conn := config.ConnectDB()
	queries := db.New(conn)

	// Dependency injection
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	// Routes
	routes.Register(app, h)

	// Health check (optional)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK"})
	})

	app.Listen(":8080")
}
