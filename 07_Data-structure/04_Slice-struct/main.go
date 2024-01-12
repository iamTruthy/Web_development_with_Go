package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("sage.gohtml"))

}

func main() {

	termintor := sage{
		Name:  "Terminator",
		Motto: "asta la vista baby",
	}

	harry := sage{
		Name:  "Harry Potter",
		Motto: "Expecto Patronium",
	}

	venom := sage{
		Name:  "Venom",
		Motto: "We are venom",
	}

	uncleBen := sage {
		Name: "Spiderman",
		Motto: "With great power, comes great responsibility",
	}

	sages := []sage{termintor, harry, venom, uncleBen}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}
