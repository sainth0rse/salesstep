package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/h0rse/ss/config"
	"github.com/h0rse/ss/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler – обработчик логина
func LoginHandler(c *fiber.Ctx) error {
	var req LoginRequest
	// Считываем JSON из тела запроса
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Невалидный запрос",
		})
	}

	userService := services.NewUserService()

	// Ищем пользователя по email
	user, err := userService.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный email или пароль",
		})
	}

	// Сравниваем хэш пароля
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный email или пароль",
		})
	}

	// Генерируем JWT (Access Token) на 15 минут
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при генерации токена",
		})
	}

	// Возвращаем токен
	return c.JSON(fiber.Map{
		"message": "Успешный вход",
		"token":   tokenString,
	})
}

// ProtectedHandler – пример защищённого маршрута
func ProtectedHandler(c *fiber.Ctx) error {
	// Считываем заголовок Authorization
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Отсутствует токен",
		})
	}

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неподдерживаемый метод подписи")
		}
		return []byte(config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный токен",
		})
	}

	// Достаём user_id из claims
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	return c.JSON(fiber.Map{
		"message": "Это защищённый маршрут",
		"user_id": userID,
	})
}
