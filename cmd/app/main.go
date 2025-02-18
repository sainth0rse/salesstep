package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/h0rse/ss/config"
	"github.com/h0rse/ss/internal/routes"
)

func main() {
	// Инициализируем базу данных
	if err := config.InitDB(); err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}

	// Инициализация MinIO
	if err := config.InitMinio(); err != nil {
		log.Fatalf("MinIO init error: %v", err)
	}

	// Создаём приложение Fiber
	app := fiber.New()

	// Включаем CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080", // Разрешаем фронтенд
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Authorization",
	}))

	// Подключаем маршруты
	routes.SetupRoutes(app)

	fmt.Println("Сервер запущен на порту 3000")
	app.Listen(":3000")
}
