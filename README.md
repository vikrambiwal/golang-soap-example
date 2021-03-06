# simple example

```go
package main

import (
"os"
"text/template"
)

type Todo struct {
  Name string
  Description string
}

func main() {
  td := Todo{"Test templates", "Let's test a template to see the magic."}

  t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
  if err != nil {
    panic(err)
  }
  err = t.Execute(os.Stdout, td)
  if err != nil {
    panic(err)
  }
}
```

# loop example

```go
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

```
