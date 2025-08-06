package main

import (
	"errors"
	"fmt"
)

type Todo struct {
	Name      string
	Completed bool
}

type Todos []Todo

func main() {
	todos := Todos{}
	todos.Add("Buy milk")
	todos.Add("Buy bread")
	todos.Add("Read 1 page")
	todos.List()
	fmt.Printf("\n")
	todos.delete(1)
	todos.List()

}

func (todos *Todos) Add(name string) {
	task := Todo{
		Name:      name,
		Completed: false,
	}
	*todos = append(*todos, task)
}

func (todos *Todos) List() {
	for i, t := range *todos {
		completed := "[ ]"

		if t.Completed {
			completed = "[x]"
		}
		fmt.Printf("%d -  %s %s\n", i+1, t.Name, completed)
	}
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

