package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	Name      string
	Completed bool
}

type Todos []Todo

func main() {
	var Tasks Todos
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "break" {
			Tasks.List()
			break
		} else {
			Tasks.Add(scanner.Text())
		}
	}

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
