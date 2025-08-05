package main

import (
	"bufio"
	"os"
)

type Todo struct {
	Name      string
	Completed bool
}

type Todos []Todo

func main() {
	var tasks Todos
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tasks.Add(scanner.Text())
	}
}

func (t *Todos) Add(name string) {
	task := Todo{
		Name:      name,
		Completed: false,
	}
	*t = append(*t, task)
}
