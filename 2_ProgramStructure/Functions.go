package main

import "fmt"

// function to multiply two numbers // (x,y int) is also fine
func Multiply(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(Multiply(5, 2))
}
