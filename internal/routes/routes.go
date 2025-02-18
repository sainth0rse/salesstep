package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h0rse/ss/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Группа /api
	api := app.Group("/api")

	// Группа авторизации (регистрация и логин)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.RegisterHandler) // из register.go
	auth.Post("/login", handlers.LoginHandler)       // из auth.go

	// Пример защищённого маршрута
	protected := api.Group("/protected")
	protected.Get("/", handlers.ProtectedHandler) // из auth.go

	// Группа /api/profile
	profile := api.Group("/profile")
	profile.Post("/photo", handlers.UploadProfilePhoto)
	profile.Put("/", handlers.UpdateProfileHandler)

}
