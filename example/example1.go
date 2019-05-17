package main

import (
	"fmt"

	"github.com/mayur-tolexo/mold"
)

//Company : Registration Mail
type Company struct {
	Company string
}

func main() {
	mold, _ := mold.NewHTMLTemplate()
	mold.HTMLPath = "mail.html"
	data := Company{"Mayur Das"}

	if err := mold.Execute(data); err == nil {
		mold.PDF(".", "test.pdf")
	} else {
		fmt.Println(err)
	}
}
