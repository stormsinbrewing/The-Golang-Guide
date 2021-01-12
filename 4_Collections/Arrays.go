package main

import "fmt"

func main() {
	// Var declaration
	var hello [2]string
	hello[0] = "Hello,"
	hello[1] = "World!"
	fmt.Println(hello)

	// Shorthand Declaration
	odd := [5]int{1, 3, 5, 7, 9}
	fmt.Println(odd)

	// Implicit Size Declaration
	nish := [...]int{5, 2, 52, 4, 421, 4, 21, 5, 12}
	fmt.Println(nish)
}
