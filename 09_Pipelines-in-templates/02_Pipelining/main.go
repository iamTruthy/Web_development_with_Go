package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

func double(x int) int {
	return x * 2
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"dbl":  double,
	"sqr":  square,
	"sqrt": sqRoot,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("pipe.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "pipe.gohtml", 2)
	if err != nil {
		log.Fatalln(err)
	}

}
