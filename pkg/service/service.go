package service

import "github.com/Wolkodaf2946/todolist/pkg/repository"

// объеденим заготовки интерфейсов для сущностей

type Authorization interface {
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
	return &Service{}
}
