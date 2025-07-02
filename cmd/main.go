package main

import (
	"log"

	"github.com/Wolkodaf2946/todolist"
	"github.com/Wolkodaf2946/todolist/pkg/handler"
	"github.com/Wolkodaf2946/todolist/pkg/repository"
	"github.com/Wolkodaf2946/todolist/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todolist.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		// InitRoutes() возвращает объекта типа "указатель на gin.Engine",
		// который реализует интерфейс хэндлера из пакета http,
		// поэтому можем его использовать в качестве аргумента
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
