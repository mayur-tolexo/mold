[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://www.godoc.org/github.com/mayur-tolexo/mold)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayur-tolexo/mold)](https://goreportcard.com/report/github.com/mayur-tolexo/mold)
[![Release](https://img.shields.io/github/release/mayur-tolexo/mold.svg?style=flat-square)](https://github.com/mayur-tolexo/mold/releases)

# mold
mold your templated to HTML/ TEXT/ PDF easily.



### install
```
go get github.com/mayur-tolexo/mold
```


### Example
```
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

	tmpl := `
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
	
	mold, _ := mold.NewHTMLTemplate(tmpl)

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
```

#### Example1 Output
![Screenshot 2019-05-17 at 4 02 15 PM](https://user-images.githubusercontent.com/20511920/57922383-41d2d900-78bd-11e9-8ee3-b3de9b51c8db.png)

