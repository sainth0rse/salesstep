package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/h0rse/ss/config"
	"github.com/h0rse/ss/internal/services"
	minio "github.com/minio/minio-go/v7"
)

// UploadProfilePhoto – загружает файл в MinIO и обновляет поле photo_url в таблице profiles
func UploadProfilePhoto(c *fiber.Ctx) error {
	// Пример: берём userID=1 (в реальном проекте вытаскиваем из JWT)
	userID := 1

	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Файл не передан",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("Ошибка при открытии файла: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось открыть файл",
		})
	}
	defer file.Close()

	fileName := fmt.Sprintf("profile_%d_%d_%s", userID, time.Now().Unix(), fileHeader.Filename)

	// Загрузить в MinIO
	info, err := config.MinioClient.PutObject(
		context.Background(),
		config.MinioBucket,
		fileName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
	)
	if err != nil {
		log.Printf("MinIO PutObject error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при загрузке в MinIO",
		})
	}

	photoURL := fmt.Sprintf("http://localhost:9000/%s/%s", config.MinioBucket, fileName)

	// Обновляем photo_url в таблице profiles
	profileService := services.NewProfileService()
	if err := profileService.UpdatePhotoURL(userID, photoURL); err != nil {
		log.Printf("DB error in UpdatePhotoURL: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить профиль",
		})
	}

	return c.JSON(fiber.Map{
		"message":   "Фото успешно загружено",
		"photo_url": photoURL,
		"size":      info.Size,
	})
}
