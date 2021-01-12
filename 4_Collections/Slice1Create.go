package main

import "fmt"

func main() {

	// Way 1: Slice Literal - Initialize when declared by passing values
	nish1 := []string{"Nishkarsh", "Raj", "Khare"}
	fmt.Println(nish1)

	// Way 2: Make function - Declare with an initial size
	nish2 := make([]string, 3)
	nish2[0] = "Nishkarsh"
	nish2[1] = "Raj"
	nish2[2] = "Khare"
	fmt.Println(nish2)
}
