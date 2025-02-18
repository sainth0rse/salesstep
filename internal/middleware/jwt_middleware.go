package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/h0rse/ss/config"
)

// RequireAuth – middleware для проверки JWT
func RequireAuth(c *fiber.Ctx) error {
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
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Неверный токен (user_id)",
		})
	}

	// userID может быть float64, переведём в int
	userID := int(userIDFloat)

	// Кладём userID в Locals, чтобы обработчики могли его взять
	c.Locals("userID", userID)

	// Идём дальше
	return c.Next()
}
