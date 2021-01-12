package main

import "fmt"
import "strings"

func main() {
	nish := "Nishkarsh Raj Khare"
	// count number of vowels in name
	count := strings.Count(nish, "a") + strings.Count(nish, "A") + strings.Count(nish, "E") + strings.Count(nish, "e") + strings.Count(nish, "i") + strings.Count(nish, "I") + strings.Count(nish, "O") + strings.Count(nish, "o") + strings.Count(nish, "u") + strings.Count(nish, "U")
	fmt.Println(count)
}
