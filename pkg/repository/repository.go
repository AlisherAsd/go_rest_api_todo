package repository

// Интерфейс для авторизации пользователя
type Authorization interface {

}

// Интерфейс для работы со списками
type TodoList interface {

}

// Интерфейс для работы с задачами
type TodoItem interface {

}

// Структура для работы с репозиториями
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Функция для создания нового репозитория
func NewRepository() *Repository {
	return &Repository{}
}