package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("file.gohtml"))
}

func main() {

	xs := []int{1, 2, 3, 4, 5}

	data := struct {
		Nums []int
		Word string	
	}{
		 xs,

		 "Count Dracular",
	} 

	err := tpl.ExecuteTemplate(os.Stdout, "file.gohtml", data)
	if err != nil {
		log.Println(err)
	}

}
