package models

// Сущность списка задач
type TodoList struct {
	Id       	int    `json:"id"`
	Title     	string `json:"title"`
	Description string `json:"description"`
}

// Сущность задачи
type TodoItem struct {
	Id       	int    `json:"id"`
	Title     	string `json:"title"`
	Description string `json:"description"`
	Done 		string `json:"done"`
}

// todo
type UsersList struct {
	Id int
	UserId int
	ListId int
}

// todo
type ListsItem struct {
	Id int
	ItemId int
	ListId int
}



