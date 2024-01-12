package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("pdf.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// nf is a variable of type pointer to os.File. nf represents a new file (index.html).
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating new file:", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
