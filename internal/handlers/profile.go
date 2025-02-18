package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h0rse/ss/internal/services"
)

// ProfileUpdateRequest – модель входных данных для PUT /api/profile
type ProfileUpdateRequest struct {
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	Requisites string `json:"requisites"`
}

// UpdateProfileHandler – обработчик для PUT /api/profile
func UpdateProfileHandler(c *fiber.Ctx) error {
	// userID достаём из middleware (Locals) или пока хардкодим
	userID := 1

	var req ProfileUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Невалидный запрос",
		})
	}

	profileService := services.NewProfileService()
	err := profileService.UpdateProfile(
		userID,
		req.FullName,
		req.Phone,
		req.Company,
		req.Position,
		req.Requisites,
	)
	if err != nil {
		log.Printf("DB error in UpdateProfile: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при обновлении профиля",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Профиль обновлён",
	})
}
