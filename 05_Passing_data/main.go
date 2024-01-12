package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("test.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "test.gohtml", "Cookies")
	if err != nil {
		log.Fatalln(err)
	}
	
}