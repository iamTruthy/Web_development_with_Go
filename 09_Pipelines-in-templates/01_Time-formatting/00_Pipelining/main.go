package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func dayMonthYear(t time.Time) string {
	return t.Format(time.RFC1123)
}

var fm = template.FuncMap{
	"fDMY": dayMonthYear,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("time.gohtml"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "time.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
