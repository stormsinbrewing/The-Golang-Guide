package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. For loop basic
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// 2. For loop without pre and post statements
	i := 1 // last i is out of scope - redeclaration
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 1
	}

	// 3. For loop as while loop
	i = 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 1
	}

	// 4. Infinite loop - print Hello World for 0.0001 seconds
	start := time.Now()
	for {
		fmt.Println("Hello, World!")
		if time.Since(start).Seconds() > 0.0001 {
			break
		}
	}

}
