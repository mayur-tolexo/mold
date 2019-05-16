package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		var tpl bytes.Buffer
		// tmpl.Execute(w, data)
		tmpl.Execute(&tpl, data)
		w.Write(tpl.Bytes())
		pdfg, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			log.Fatal(err)
		}

		// page := wkhtmltopdf.NewPage(tpl.String())
		// page.FooterRight.Set("[page]")
		// page.FooterFontSize.Set(10)
		// page.Zoom.Set(0.95)

		pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(tpl.Bytes())))

		// Create PDF document in internal buffer
		err = pdfg.Create()
		if err != nil {
			log.Fatal(err)
		}

		// Write buffer contents to file on disk
		err = pdfg.WriteFile("./simplesample.pdf")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Done")
		// Output: Done

	})

	http.ListenAndServe(":80", nil)
}
