package main

import (
	"fmt"
	"net/http"

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
	mold.HTMLPath = "layout.html"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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
			w.Write(mold.Bytes())
			mold.PDF(".", "test.pdf")
		} else {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
