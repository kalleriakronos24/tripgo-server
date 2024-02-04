package pdfGenerator

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type RequestPdf struct {
	body string
}

func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

func (r *RequestPdf) GeneratePDF(pdfPath string, args []string) (bool, error) {

	// set the predefined path in the wkhtmltopdf's global state
	wkhtmltopdf.SetPath("/usr/bin/wkhtmltopdf")

	t := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("cloneTemplate/", 0777)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
	err1 := os.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Use arguments to customize PDF generation process
	for _, arg := range args {
		switch arg {
		case "low-quality":
			pdfg.LowQuality.Set(false)
		case "no-pdf-compression":
			pdfg.NoPdfCompression.Set(true)
		case "grayscale":
			pdfg.Grayscale.Set(true)
			// Add other arguments as needed
		}
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(dir + "/cloneTemplate")

	return true, nil
}
