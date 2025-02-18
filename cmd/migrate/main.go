package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/h0rse/ss/config"
)

func main() {
	// Инициализация БД
	if err := config.InitDB(); err != nil {
		log.Fatalf("Не удалось инициализировать базу: %v", err)
	}

	// Читаем файл миграций
	data, err := ioutil.ReadFile("migrations/001_init.sql")
	if err != nil {
		log.Fatalf("Ошибка чтения файла миграций: %v", err)
	}

	// Выполняем SQL
	_, err = config.DB.Exec(string(data))
	if err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	fmt.Println("Миграции успешно применены!")
}
