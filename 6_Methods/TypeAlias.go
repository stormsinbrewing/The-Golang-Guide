// Alias is used to define functions on inbuilt types

package main

import (
	"fmt"
	"strings"
)

// String type
type nish string

// User defined alias functions always called as type("argument").function()
func (s nish) UpUp() string {
	return strings.ToUpper(string(s)) //convert user defined type to string
}

func main() {

	// hi := "Hello"  // To use Imported methods - have to export variable - capitalize them
	Hi := "Hello"
	fmt.Println(strings.ToUpper(Hi))
	fmt.Println(nish("World!").UpUp())
}
