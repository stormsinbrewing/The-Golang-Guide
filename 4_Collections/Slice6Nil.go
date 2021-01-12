package main

import "fmt"

func main() {
	// Creating a nil slice
	slice := []int{}
	fmt.Println(slice, len(slice), cap(slice)) //lenght and capacity is zero
	if slice == nil {
		fmt.Println("The slice is empty")
	}
}
