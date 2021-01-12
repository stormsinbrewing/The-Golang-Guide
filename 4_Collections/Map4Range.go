// Range for Map returns key, value
package main

import "fmt"

func main() {
	student := map[string]int{
		"Nishkarsh": 500060720,
		"Karan":     500060911,
		"Manthan":   500060128,
		"Rishabh":   500060226,
	}

	for key, value := range student {
		fmt.Printf("Student %s's SAP is %d\n", key, value)
	}
}
