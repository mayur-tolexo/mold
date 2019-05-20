package main

import (
	"fmt"

	"github.com/mayur-tolexo/mold"
)

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
