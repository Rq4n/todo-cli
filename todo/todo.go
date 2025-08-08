package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title     string
	Completed bool
}

type Todos []Todo

func (todos *Todos) Add(name string) {
	task := Todo{
		Title:     name,
		Completed: false,
	}
	*todos = append(*todos, task)
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed")
	for index, t := range *todos {
		completed := "❌"

		if t.Completed {
			completed = "✅"
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed)
	}
	table.Render()
}

func (todos *Todos) Complete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Completed = !t[index].Completed
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}
