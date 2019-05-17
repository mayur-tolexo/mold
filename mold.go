package mold

import (
	"bytes"
	"html/template"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//Mold model
type Mold struct {
	HTMLTemplate string //html template string
	HTMLPath     string //html template file path
	buff         bytes.Buffer
	*wkhtmltopdf.PDFGenerator
}

//NewHTMLTemplate will create new return new html template mold
func NewHTMLTemplate(tmpl ...string) (m *Mold, err error) {
	val := ""
	if len(tmpl) > 0 {
		val = tmpl[0]
	}
	m = &Mold{
		HTMLTemplate: val,
	}
	m.PDFGenerator, err = wkhtmltopdf.NewPDFGenerator()
	return
}

//Execute will execute the template
func (m *Mold) Execute(data interface{}) (err error) {
	tmpl := template.Must(template.ParseFiles(m.HTMLPath))
	err = tmpl.Execute(&m.buff, data)
	return
}

//String will return string from mold buffer
func (m *Mold) String() string {
	return m.buff.String()
}

//Bytes will return bytes from mold buffer
func (m *Mold) Bytes() []byte {
	return m.buff.Bytes()
}

//PDF will create pdf
func (m *Mold) PDF(path, name string) (err error) {
	m.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(m.buff.Bytes())))
	if err = m.Create(); err == nil {
		err = m.WriteFile(path + "/" + name)
	}
	return
}
