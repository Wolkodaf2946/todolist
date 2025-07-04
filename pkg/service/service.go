package service

import (
	"github.com/Wolkodaf2946/todolist"
	"github.com/Wolkodaf2946/todolist/pkg/repository"
)

// объеденим заготовки интерфейсов для сущностей

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GenerateToken(username, password string) (string, error) // возвращается токен
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
