package main

import (
	  "github.com/Rq4n/todo-cli/todo"
)

func main() {
	todos := todo.Todos{}
	todos.Add("Buy milk")
	todos.Add("Buy bread")
	todos.Add("Read 1 page")
	todos.Complete(1)
	todos.Print()


}
