package main

import "fmt"

// structure
type Points struct {
	x, y int
}

func main() {

	// Simple Map
	var m map[string]int
	//Empty map cannot be assigned - define space using make (not new)
	m = make(map[string]int)
	m["Nishkarsh"] = 500060720
	fmt.Println(m)

	// Complex Map
	var m2 map[string]Points
	m2 = make(map[string]Points)
	// m2["A"] = {2,5} //Implicit only allowed for literal declaration
	m2["B"] = Points{3, 10} //Explicit
	fmt.Println(m2)
}
