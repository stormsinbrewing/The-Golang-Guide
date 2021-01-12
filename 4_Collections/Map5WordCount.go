package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) // converts string into list of string broken by space
	count := map[string]int{}  // Non-Empty Literal declaration
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	fmt.Println(WordCount("As long As you Love me!")) //case sensitive - As not equals as
}
