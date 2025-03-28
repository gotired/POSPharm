package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/handler"
	"github.com/gotired/POSPharm/internal/repository"
	"github.com/gotired/POSPharm/internal/usecase"
	"gorm.io/gorm"
)

func NewFiberApp(db *gorm.DB) *fiber.App {
	app := fiber.New()

	// Dependency Injection
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Routes
	api := app.Group("/api")
	api.Get("/users", userHandler.GetUsers)
	api.Get("/users/:id", userHandler.GetUser)
	api.Post("/users", userHandler.CreateUser)

	return app
}
