package service

import (
	"alisherAsd/pkg/repository"
)

// Интерфейс для авторизации пользователя
type Authorization interface {

}

// Интерфейс для работы со списками
type TodoList interface {

}

// Интерфейс для работы с задачами
type TodoItem interface {

}

// Структура для работы с сервисами
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Функция для создания нового сервиса
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}