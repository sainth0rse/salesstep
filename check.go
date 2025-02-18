package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Генерируем хэш для пароля "12345"
	hash, err := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Хэш для '12345':", string(hash))

	// Проверяем совпадение
	if err := bcrypt.CompareHashAndPassword(hash, []byte("12345")); err != nil {
		fmt.Println("Пароль НЕ совпал:", err)
	} else {
		fmt.Println("Пароль совпал! Всё ок.")
	}
}
