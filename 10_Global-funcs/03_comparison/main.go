package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("file.gohtml"))
}

func main() {

	game := struct {
		Score1 int
		Score2 int
	}{
		4,
		9,
	}

	err := tpl.Execute(os.Stdout, game)
	if err != nil {
		log.Fatalln(err)
	}

}
