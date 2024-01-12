package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("magic.gohtml"))
}

func main() {

	lifestyle := map[string]string{
		"Mic":     "Music",
		"Cars":    "Racing",
		"Clothes": "Fashion",
		"Laptops": "Gadgets",
	}

	err := tpl.Execute(os.Stdout, lifestyle)
	if err != nil {
		log.Fatalln(err)
	}
}
