package main

import (
	"fmt"
	"log"

	"github.com/Rq4n/todo-cli/storage"
	"github.com/Rq4n/todo-cli/todo"
)

func main() {
	todoStorage := storage.NewStorage[todo.Todos]("todos.json")

	var todos todo.Todos
	err := todoStorage.Load(&todos)
	if err != nil {
		fmt.Println("Nenhum arquivo encontrado, iniciando lista vazia ...")
	}

	todos.Add("Buy milk")
	todos.Add("Buy bread")
	todos.Add("Read 1 page")
	todos.Print()

	err = todoStorage.Save(todos)
	if err != nil {
		log.Fatal(err)
	}
}
