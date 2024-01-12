package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

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

// firstThree is a field element of FuncMap and also a function which takes in a string parameter, and returns a string.
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

// fm is a value of type FuncMap. FuncMap is a map, with a key of type string and value of type empty interface (anything).
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("fans.gohtml"))
}

func main() {

	m := movie{
		Name: "Spiderman",
		Line: "Your friendly neighborhood spiderman",
	}

	m2 := movie{
		Name: "Batman",
		Line: "I'm Batman",
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

	films := []movie{m, m2, m3}
	luxury := []brand{b, b2, b3}

	//data is a struct with field elements of type struct-slice or slice of struct.
	data := group{
		Cinema:  films,
		Fashion: luxury,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
