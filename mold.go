package mold

import (
	"bytes"
	"html/template"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//Mold model
type Mold struct {
	HTMLTemplate string //html template string
	HTMLPath     string //html template file path
	buff         []byte
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
	var buff bytes.Buffer
	if m.HTMLTemplate == "" {
		tmpl := template.Must(template.ParseFiles(m.HTMLPath))
		err = tmpl.Execute(&buff, data)
	} else {
		tmpl := template.Must(template.New("").Parse(m.HTMLTemplate))
		err = tmpl.Execute(&buff, data)
	}
	m.buff = buff.Bytes()

	return
}

//String will return string from mold buffer
func (m *Mold) String() string {
	return string(m.buff)
}

//Bytes will return bytes from mold buffer
func (m *Mold) Bytes() []byte {
	return m.buff
}

//PDF will create pdf
func (m *Mold) PDF(path, name string) (err error) {
	m.Dpi.Set(300)
	// m.Grayscale.Set(true)
	m.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(m.String())))
	if err = m.Create(); err == nil {
		err = m.WriteFile(path + "/" + name)
	}
	return
}
