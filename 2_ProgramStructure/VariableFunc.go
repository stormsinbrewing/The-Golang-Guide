package main
import "fmt"

func main() {
	// creating a variable which stores a function
	nish := func() {
	fmt.Println("Hello, World!")
	}
	nish()
	nish()
	nish()
}
