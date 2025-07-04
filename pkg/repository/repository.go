package repository

import (
	"github.com/Wolkodaf2946/todolist"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)

	// для генерации токена нам нужно получить пользователя из базы
	// если такого польователя нет, то возвращаем ошибку
	// иначе генерируем токен, в который записываем id пользователя
	GetUser(username, password string) (todolist.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db), // инициализируем репозиторий в конструкторе
	}
}
