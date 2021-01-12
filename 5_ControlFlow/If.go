package main

import "fmt"

func main() {
	n := 5

	// Single condition
	if n == 5 {
		fmt.Println("First Check!")
	}

	// Multiple conditions
	if i := 4; i%2 == 1 {
		fmt.Println("Second Check!")
	}

	// If else statement

	if false {
		//wont execute
	} else {
		fmt.Println("Third Check!")
	}

}
