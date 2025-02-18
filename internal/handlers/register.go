package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h0rse/ss/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`    // обязательно
	Password string `json:"password"` // обязательно
}

// RegisterHandler – обработчик регистрации
func RegisterHandler(c *fiber.Ctx) error {
	var req RegisterRequest
	// Парсим JSON
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Невалидный запрос",
		})
	}

	// Проверяем обязательные поля
	if req.Email == "" || req.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Не заполнены обязательные поля (email, password)",
		})
	}

	userService := services.NewUserService()

	// Проверяем, нет ли пользователя с таким email
	existingUser, err := userService.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil && existingUser.ID != 0 {
		// Значит пользователь найден
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"error": "Пользователь с таким email уже существует",
		})
	}

	// Хэшируем пароль
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при хэшировании пароля",
		})
	}

	// Создаём пользователя
	userID, err := userService.CreateUser(req.Email, string(hashed))
	if err != nil {
		log.Printf("DB error in CreateUser: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при создании пользователя",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Пользователь успешно зарегистрирован",
		"user_id": userID,
	})
}
