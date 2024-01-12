package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type movie struct {
	Name string
	Line string
}

type brand struct {
	Name string
	Type string
}

type group struct {
	Cinema  []movie
	Fashion []brand
}

func init() {
	tpl = template.Must(template.ParseFiles("task.gohtml"))
}

func main() {

	m := movie{
		Name: "Spiderman",
		Line: "Your friendly neighborhood spiderman",
	}

	m2 := movie{
		Name: "The Flash",
		Line: "My name is Barry Allen and i am the fastest man alive",
	}

	m3 := movie{
		Name: "Game of Thrones",
		Line: "Valla Mogulis",
	}

	b := brand{
		Name: "Nike",
		Type: "Clothing",
	}

	b2 := brand{
		Name: "Tesla",
		Type: "Cars",
	}

	b3 := brand{
		Name: "Richard Mille",
		Type: "Wrist watches",
	}

	film := []movie{m, m2, m3}
	luxury := []brand{b, b2, b3}

	data := group{
		Cinema:  film,
		Fashion: luxury,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
