package main

import "fmt"

type Student struct {
	name   string
	sap    int
	roll   string
	branch string
}

func main() {
	// print subset without declaration - generic
	fmt.Println(Student{name: "Nishkarsh Raj", sap: 500060720}) // named declaration - key value pair
}
