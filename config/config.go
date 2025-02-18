package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB – глобальное соединение с базой
var DB *sql.DB

// JWTSecret – ключ для подписи JWT (в реальном проекте брать из переменной окружения)
var JWTSecret = getEnv("JWT_SECRET", "SUPER_SECRET_KEY")

// InitDB – инициализация соединения с PostgreSQL
func InitDB() error {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "ssuser")
	dbPassword := getEnv("DB_PASSWORD", "sspassword")
	dbName := getEnv("DB_NAME", "ssdb")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("ошибка при открытии соединения: %w", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("ошибка при пинге базы данных: %w", err)
	}

	DB = db
	fmt.Println("Подключение к базе данных успешно!")
	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
