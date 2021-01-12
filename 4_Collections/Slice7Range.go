package main

import "fmt"

func main() {
	sl := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

	// 1. range function returns index and value
	for index, value := range sl {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	// 2. Skip Index with _
	for _, value := range sl {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// 3. Skip Value
	for i := range sl {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 4. break keyword - Print first 4 elements
	for i, v := range sl {
		if i == 4 {
			break
		}
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 5. Continue keyword - Skip 4th element
	for i, v := range sl {
		if i == 3 {
			continue
		}
		fmt.Printf("%d ", v)
	}
	fmt.Println()

}
