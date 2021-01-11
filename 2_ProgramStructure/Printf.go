package main
import "fmt"

func main() {
	name := "Caprica-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s",
		name, aka) //printing variables within the string
}
