package main

import (
	"log"

	"github.com/Nurka144/todo-app"
	"github.com/Nurka144/todo-app/pkg/handler"
	"github.com/Nurka144/todo-app/pkg/repository"
	"github.com/Nurka144/todo-app/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("5131", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error in server running http server: %s", err.Error())
	}
}
