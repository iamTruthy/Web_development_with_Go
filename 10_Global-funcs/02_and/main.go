package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("file.gohtml"))
}

func main() {

	user1 := user{
		Name:  "MLK",
		Motto: "I have a dream that one day...",
		Admin: true,
	}

	user2 := user{
		Name:  "Julius Ceaser",
		Motto: "Et tu Brutus?",
		Admin: false,
	}

	user3 := user{
		Name:  "Albert Enstien",
		Motto: "E = mc^2",
		Admin: true,
	}

	user4 := user{
		Name:  "Bobby McFerrin",
		Motto: "Don't worry, be happy",
		Admin: true,
	}

	users := []user{user1, user2, user3, user4}

	err := tpl.ExecuteTemplate(os.Stdout, "file.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}

}
