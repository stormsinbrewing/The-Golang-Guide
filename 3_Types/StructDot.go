package main

import "fmt"

type Student struct {
	name, roll string
	sap        int
}

func main() {
	nishkarsh := Student{name: "Nishkarsh Raj Khare", roll: "R171217041"}
	nishkarsh.sap = 500060720
	fmt.Println(nishkarsh)
}
