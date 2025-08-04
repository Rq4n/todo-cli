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
	scanner := bufio.NewScanner(os.Stdin)
	var todos Todos

	for scanner.Scan() {
		text := scanner.Text()
		newTask := Todo{
			Name:      text,
			Completed: false,
		}
		todos = append(todos, newTask)
		if text == "break" {
			fmt.Println("Encerrando ...")
			break
		}
	}
	for _, t := range todos {
		fmt.Printf("[%v], [%v]", t.Name, t.Completed)
	}
}
