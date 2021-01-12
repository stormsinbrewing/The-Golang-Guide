package main

import "fmt"

// Base struct
type person struct {
	fName, lName string
	age          int
}

// Creating method for the structure
// format - func (obj *struct_name) funcname(attributes) return type
func (p *person) Hello() string {
	//sprintf function stores and returns but does not print to console
	return	fmt.Sprintf("Hello, %s %s. How are you?\n", p.fName, p.lName)
}

// child struct
type player struct {
	person
	pid string
}

func main() {
	obj := player{}
	obj.fName = "Nishkarsh"
	obj.lName = "Raj"
	fmt.Println(obj.Hello())
}
