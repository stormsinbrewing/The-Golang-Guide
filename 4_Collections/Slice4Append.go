package main

import "fmt"

func main() {
	nish := []string{} //create empty string
	// nish[0] = "Nishkarsh" // Error because empty string is nil without any size and index -

	// Append a value
	nish = append(nish, "Nishkarsh")
	fmt.Println(nish)

	// Append multiple values
	nish = append(nish, "Raj", "Khare")
	fmt.Println(nish)

	// Append another slice
	code := []string{"Wrote", "This", "Code"}
	nish = append(nish, code...) //ellipse is used to append only same type slices
	fmt.Println(nish)
}
