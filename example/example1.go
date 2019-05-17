package main

import (
	"fmt"
	"net/http"

	"github.com/mayur-tolexo/mold"
)

//Company : Registration Mail
type Company struct {
	Company string
}

func main() {
	mold, _ := mold.NewHTMLTemplate()
	mold.HTMLPath = "mail.html"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := Company{"Mayur Das"}

		if err := mold.Execute(data); err == nil {
			w.Write(mold.Bytes())
			mold.PDF(".", "test.pdf")
		} else {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
