package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type brand struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("brand.gohtml"))
}

func main() {

	nike := brand{
		Name:  "Airforce one",
		Motto: "Just do it",
	}

	err := tpl.Execute(os.Stdout, nike)
	if err != nil {
		log.Fatalln(err)
	}

}
