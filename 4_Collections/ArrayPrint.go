package main

import "fmt"

func main() {
	nish := [...]string{"Nishkarsh", "Raj", "Khare"}

	// Print string as list
	fmt.Println(nish)

	// Print Quoted string
	fmt.Printf("%q\n", nish)

	// Print string as list
	fmt.Printf("%s\n", nish)
}
