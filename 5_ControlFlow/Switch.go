// only same type for all cases
// 1. multiple values in same case can be there
// no runs for all successive cases if matched unlike C
// 2. To do above, fallthrough keyword can be used
// 3. To break fallthrough, break can be used

package main

import "fmt"

func main() {
	num := 5

	// 1. Multiple values
	switch num {
	case 0, 1, 2, 3:
		fmt.Println("Fail")
	case 4, 5, 6:
		fmt.Println("Average")
	case 7, 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Perfect")
	default:
		fmt.Println("No match found")
	}

	// 2. Fall through
	switch num {
	case 0, 1, 2, 3:
		fmt.Println("Fail")
		fallthrough
	case 4, 5, 6:
		fmt.Println("Average")
		fallthrough
	case 7, 8, 9:
		fmt.Println("Good")
		fallthrough
	case 10:
		fmt.Println("Perfect")
		fallthrough
	default:
		fmt.Println("No match found")
	}

	// 3. Break
	switch num {
	case 0, 1, 2, 3:
		fmt.Println("Fail")
		fallthrough
	case 4, 5, 6:
		fmt.Println("Average")
		fallthrough
	case 7, 8, 9:
		fmt.Println("Good")
		fallthrough
	case 10:
		fmt.Println("Perfect")
		break // dont print Default
	default:
		fmt.Println("No match found")
	}
}
