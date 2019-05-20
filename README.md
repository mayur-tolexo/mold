[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://www.godoc.org/github.com/mayur-tolexo/mold)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayur-tolexo/mold)](https://goreportcard.com/report/github.com/mayur-tolexo/mold)
[![Release](https://img.shields.io/github/release/mayur-tolexo/mold.svg?style=flat-square)](https://github.com/mayur-tolexo/mold/releases)

# mold
mold your templated to HTML/ TEXT/ PDF easily.



### install
```
go get github.com/mayur-tolexo/mold
```


### Example 1
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

### Example 2
```

//Invoice details
type Invoice struct {
	InvoiceNo   string
	InvoiceDate string
	Currency    string
	AmountDue   float64
	Items       []ItemDetail
	Total       float64
}

//ItemDetail : Item details
type ItemDetail struct {
	Name     string
	Desc     string
	Amount   float64
	Qty      int
	Currency string
	Total    float64
}

func main() {
	mold, _ := mold.NewHTMLTemplate()
	mold.HTMLPath = "invoice.html"

	item := []ItemDetail{
		{
			Name:     "Front End Consultation",
			Desc:     "Experience Review",
			Currency: "Rs.",
			Amount:   150,
			Qty:      4,
			Total:    600,
		},
	}

	data := Invoice{
		InvoiceNo:   "Invoice",
		InvoiceDate: "January 1, 2019",
		Currency:    "Rs.",
		AmountDue:   600,
		Items:       item,
		Total:       600,
	}

	if err := mold.Execute(data); err == nil {
		mold.PDF(".", "invoice.pdf")
	} else {
		fmt.Println(err)
	}
}
```

#### Output
![Screenshot 2019-05-20 at 4 44 19 PM](https://user-images.githubusercontent.com/20511920/58017758-a476df80-7b1e-11e9-911b-dcd44c0bfe9a.png)


