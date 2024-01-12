package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("test.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "test.gohtml",`Drop by drop the bucket gets filled.` )
if err != nil {
	log.Fatalln(err)
}
}