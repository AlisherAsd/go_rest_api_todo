package main

import (
	todo "alisherAsd"
	"alisherAsd/pkg/handler"
	"log"
)

// Точка входа приложения
// Создаём экземпляр структуры Handler
// Запускаем сервер
func main() {
	handler := new(handler.Handler)

	server := new(todo.Server)
	// Из server.go и handler.go запускаем сервер с обработчиками
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error run http server: %s", err.Error())
	}

}