// If no explicit return for variable specified, the current (if no current then default) value is returned
package main

import "fmt"

func Generic() (Name, Empty string) {
	//Updating Name
	Name = "Nishkarsh"
	return // returns Nishkash and Empty String
}

func main() {
	name, empty := Generic()
	fmt.Println(empty + " " + name)
}
