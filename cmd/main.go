package main

import (
	todo "alisherAsd"
	"log"
)

// Точка входа приложения
func main() {
	server := new(todo.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error run http server: %s", err.Error())
	}

}