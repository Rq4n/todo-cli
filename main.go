package main

import (
	"bufio"
	"errors"
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

/*
Receber um índice (ex: posição da tarefa que você quer apagar)
Validar se o índice é válido (não pode ser negativo ou maior que o tamanho da lista)
Remover a tarefa daquele índice do slice
Atualizar o slice para refletir essa remoção
*/
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Indece fora do limite")
	} else {
		return nil
	}
}

func (todos *Todos) delete(index int) {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	// para remover  ele ganha um novo slice no local do indice passado, so que esse valor seria vazio
}
