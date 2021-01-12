package main

import "fmt"

type Points struct {
	x, y int
}

func main() {
	// literal declaration with implicit definition for structures
	m := map[string]Points{
		"A": {2, 4},
		"B": {5, 2},
		"C": {10, 8},
	}
	fmt.Println(m)

	// 1. Insertion - New declaration without literal requires type specification
	// m["Origin"] = {0,0}
	m["Origin"] = Points{0, 0}
	fmt.Println(m)

	// 2. Update
	m["A"] = Points{2, 5}
	fmt.Println(m)

	// 3. Retrieval
	fmt.Println(m["Origin"])

	// 4. delete
	delete(m, "Origin")
	fmt.Println(m)

	// 5. Existence - (value, presence) if present - print value | if absent print default type value
	v, p := m["Origin"]
	fmt.Printf("%t %v", p, v)
	v, p = m["A"]
	fmt.Printf("%t %v", p, v)

}
