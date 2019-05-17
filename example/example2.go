package main

import (
	"fmt"

	"github.com/mayur-tolexo/mold"
)

//Todo model
type Todo struct {
	Title string
	Done  bool
}

//TodoPageData model
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	mold, _ := mold.NewHTMLTemplate()
	mold.HTMLTemplate = `
	<h1>{{.PageTitle}}<h1>
	<ul>
	    {{range .Todos}}
	        {{if .Done}}
	            <li class="done">{{.Title}}</li>
	        {{else}}
	            <li>{{.Title}}</li>
	        {{end}}
	    {{end}}
	</ul>
	`

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
			{Title: "Task 4", Done: true},
		},
	}

	if err := mold.Execute(data); err == nil {
		mold.PDF(".", "tmp.pdf")
	} else {
		fmt.Println(err)
	}
}
