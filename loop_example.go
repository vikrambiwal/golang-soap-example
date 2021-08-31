package main

import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	paths := []string{
		"./loop_todo.tmpl",
	}

	todos := ToDo{
		User: "vkkk",
		List: []entry{
			entry{
				Name: "one",
				Done: true,
			},
			entry{
				Name: "two",
				Done: true,
			},
		},
	}

	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(os.Stdout, todos)
	if err != nil {
		panic(err)
	}
}
