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

	magic := []string{"Houdini", "David Blaine", "Harry Potter", "Dumbledor"}

	err := tpl.Execute(os.Stdout, magic)
	if err != nil {
		log.Fatalln(err)
	}

	/*

		nf, err := os.Create("magic.html")
		if err != nil {
			log.Fatalln(err)
		}
		defer nf.Close()


		err = tpl.Execute(nf, nil)
		if err != nil {
			log.Fatalln(err)
		}

	*/

}
