package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gotired/POSPharm/internal/handler"
	"github.com/gotired/POSPharm/internal/middleware"
	"github.com/gotired/POSPharm/internal/repository"
	"github.com/gotired/POSPharm/internal/usecase"
	"gorm.io/gorm"
)

func NewFiberApp(db *gorm.DB) *fiber.App {
	app := fiber.New()

	userRepo := repository.NewUser(db)
	userCredentialRepo := repository.NewUserCredential(db)
	roleRepo := repository.NewRole(db)
	branchRepo := repository.NewBranch(db)

	userUsecase := usecase.NewUser(db, userRepo, userCredentialRepo)
	roleUsecase := usecase.NewRole(roleRepo)
	authUsecase := usecase.NewAuth()
	branchUsecase := usecase.NewBranch(branchRepo)

	userHandler := handler.NewUser(userUsecase)
	authHandler := handler.NewAuth(userUsecase, authUsecase)
	roleHandler := handler.NewRole(roleUsecase)
	branchHandler := handler.NewBranch(branchUsecase)

	appMiddleware := middleware.NewMiddleware(authUsecase)
	app.Use(cors.New())
	api := app.Group("/api")

	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)

	restrict := api.Group("")
	restrict.Use(appMiddleware.JWT)

	restrict.Get("/users", userHandler.List)
	restrict.Get("/users/:id", userHandler.GetByID)

	restrict.Get("/roles", roleHandler.List)
	restrict.Get("/branches", branchHandler.List)

	return app
}
