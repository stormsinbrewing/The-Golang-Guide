// Go allows returning multiple values
package main

import "fmt"

// Generic function with no arguments
func Generic() (string, string) {
	return "Nishkarsh", "Raj"
}

func main() {
	// declaring two variables with same function call
	fName, lName := Generic()
	fmt.Printf("My name is %s %s\n", fName, lName)
}
