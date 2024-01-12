package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	pdf, err := template.ParseFiles("pdf.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = pdf.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
