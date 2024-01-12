package main

import "fmt"

// Person is a custom type (struct) with field names of type string.
type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says: Drop by drop the bucket gets filled.")
}

// Tutor is a custom type (struct) which holds a field of type bool and also a custom type field of type person struct.
// Type tutor shows polymorphism.
type tutor struct {
	person
	canTeach bool
}

func (t tutor) speak() {
	fmt.Println(t.fname, "- an interface says: hey baby, if you've got this methods, then you're my type.")
}

// developer is  an interface that implements the speak method.
type developer interface {
	speak()
}

// code takes in a parameter of type developer which is an interface.
// the code function also implements the speak method from the developer interface.
func code(d developer) {
	d.speak()
}

func main() {

	p := person{
		"Todd",
		"McLeod",
	}

	code(p)

	t := tutor{
		person{
			"Todd",
			"McLeod",
		},
		true,
	}

	code(t)

}
