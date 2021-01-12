package main

import "fmt"

func main() {
	// literal declaration - map initialized when declared
	// syntax - map[key]value{}
	nish := map[string]int{
		"Nishkarsh": 500060720,
		"Karan":     500060911,
		"Manthan":   500060128,
		"Rishabh":   500060226, // last element must also follow a comma
	}

	fmt.Println(nish) // automatically sorts dictionary because they are lexicographic
	fmt.Printf("%#v", nish)

}
