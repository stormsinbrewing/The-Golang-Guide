package main

import "fmt"

// Base struct
type person struct {
	fName, lName string
	age          int
}

// child struct - composition rather than inheritance
type player struct {
	// every player is a person
	person
	pid string
}

func main() {

	// two ways to declare persons

	// 1. Dot Notation
	nish := player{}
	nish.fName = "Nishkarsh"
	nish.lName = "Raj"
	nish.age = 21
	nish.pid = "R41"
	fmt.Println(nish)

	// 2. Sub set type declaration
	obj := player{
		person{fName: "First", lName: "Last", age: 20},
		"RXX",
	}
	//fmt.Println(obj)
	//%+v shows field + value
	fmt.Printf("%+v", obj)
}
