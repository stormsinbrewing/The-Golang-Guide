package main

import "fmt"

func main() {
	// Slices can be made up of Arrays as well
	array := [5]int{1, 3, 5, 7, 9}

	// Slices of Slice [lo:hi] are from lo to hi-1 inclusive similar to python and by default lo = 0 and hi = len
	fmt.Println(array[2:4])
	fmt.Println(array[2:2]) //Nil slice
}
