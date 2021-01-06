package main

import (
	"fmt"
	"os"
)

func main() {
	// Declaration of multiple variables with single inbuilt function
	name, err := os.Open("/username")
	fmt.Println(name)
	fmt.Println(err) //nil

	// Swapping values like Python
	i, j := 0, 100
	i, j = j, i
	fmt.Println(i) //100
	fmt.Println(j) //0
}
