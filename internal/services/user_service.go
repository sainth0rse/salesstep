package services

import (
	"fmt"

	"github.com/h0rse/ss/config"
	"github.com/h0rse/ss/internal/models"
)

// UserService – сервис для работы с таблицей users
type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser – создаёт пользователя (email, passwordHash)
// Возвращает id нового пользователя и/или ошибку
func (s *UserService) CreateUser(email, passwordHash string) (int, error) {
	var newID int
	query := `
        INSERT INTO users (email, password_hash)
        VALUES ($1, $2)
        RETURNING id
    `
	err := config.DB.QueryRow(query, email, passwordHash).Scan(&newID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return newID, nil
}

// GetUserByEmail – ищет пользователя по email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, email, password_hash
        FROM users
        WHERE email = $1
    `
	row := config.DB.QueryRow(query, email)
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		return nil, err
	}
	return user, nil
}
