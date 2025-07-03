package repository

import (
	"github.com/Wolkodaf2946/todolist"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
}

// .sofhoi;sajdf
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
