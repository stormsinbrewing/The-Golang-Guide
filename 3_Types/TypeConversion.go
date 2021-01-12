// Go follows explicit only type conversion else error.
// Convert follows - type(value)

package main

import "fmt"

func main() {
	x := 41
	u := uint8(x)
	f := float32(x)
	const format = "%T(%v)\n"

	fmt.Printf(format, x, x)
	fmt.Printf(format, u, u)
	fmt.Printf(format, f, f)
}
