package main

import (
	"os"
	"text/template"
)

func main() {

	tpl := template.Must(template.New("Something").Parse("Hello world"))

	//We use Parse() when we are only passing in a string. ParseFile() or ParseGlob() on the other hand is used when
	//we areworking with files.
	//strings are passed into templates.New() when working with **Parse()**

	tpl.ExecuteTemplate(os.Stdout, "Something", nil)

}
