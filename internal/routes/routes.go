package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h0rse/ss/internal/handlers"
	"github.com/h0rse/ss/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// auth (без middleware)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.RegisterHandler)
	auth.Post("/login", handlers.LoginHandler)

	// Пример защищённого маршрута
	protected := api.Group("/protected")
	// Подключаем middleware RequireAuth
	protected.Use(middleware.RequireAuth)
	protected.Get("/", handlers.ProtectedHandler)

	// Профиль
	profile := api.Group("/profile")
	// Для профиля тоже используем RequireAuth, чтобы брать userID из токена
	profile.Use(middleware.RequireAuth)
	profile.Post("/photo", handlers.UploadProfilePhoto)
	profile.Put("/", handlers.UpdateProfileHandler)
}
