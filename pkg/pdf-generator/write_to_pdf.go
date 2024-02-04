package pdfGenerator

import (
	"fmt"
)

func WriteHTMLToPDF(path string, templatePath string, data interface{}) {

	r := NewRequestPdf("")

	//path for download pdf
	outputPath := path

	//html template data
	//templateData := struct {
	//	Title       string
	//	Description string
	//	Company     string
	//	Contact     string
	//	Country     string
	//}{
	//	Title:       "HTML to PDF generator",
	//	Description: "This is the simple HTML to PDF file.",
	//	Company:     "Jhon Lewis",
	//	Contact:     "Maria Anders",
	//	Country:     "Germany",
	//}

	if err := r.ParseTemplate(templatePath, data); err == nil {

		// Generate PDF with custom arguments
		args := []string{"no-pdf-compression"}

		// Generate PDF
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
