package services

import (
	"fmt"

	"github.com/h0rse/ss/config"
)

// ProfileService – сервис для работы с таблицей profiles
type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

// CreateProfile – создаёт запись в таблице profiles
func (s *ProfileService) CreateProfile(userID int, fullName, phone, company, position string) error {
	query := `
        INSERT INTO profiles (user_id, full_name, phone, company, position)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := config.DB.Exec(query, userID, fullName, phone, company, position)
	if err != nil {
		return fmt.Errorf("ошибка при создании профиля: %w", err)
	}
	return nil
}

// UpdateProfile – обновляет поля профиля (full_name, phone, company, position, requisites)
func (s *ProfileService) UpdateProfile(userID int, fullName, phone, company, position, requisites string) error {
	query := `
        UPDATE profiles
        SET full_name = $1,
            phone = $2,
            company = $3,
            position = $4,
            requisites = $5
        WHERE user_id = $6
    `
	_, err := config.DB.Exec(query, fullName, phone, company, position, requisites, userID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении профиля: %w", err)
	}
	return nil
}

// UpdatePhotoURL – обновляет поле photo_url у профиля
func (s *ProfileService) UpdatePhotoURL(userID int, photoURL string) error {
	query := `
        UPDATE profiles
        SET photo_url = $1
        WHERE user_id = $2
    `
	_, err := config.DB.Exec(query, photoURL, userID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении photo_url: %w", err)
	}
	return nil
}
