package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	// Declaration of structures
	obj1 := Point{5, 2}  // Implicit if all parameters specified
	obj2 := &Point{5, 2} // Pointer Reference
	obj3 := Point{x: 5}  // y = 0
	obj4 := Point{}      // All values set to type default
	fmt.Println(obj1, obj2, obj3, obj4)
}
