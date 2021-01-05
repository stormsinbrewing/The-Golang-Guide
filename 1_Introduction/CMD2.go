// Command Line Arguments with variable iterator loop
package main

import (
	"fmt"
	"os"
)

func main() {
	str, seperator := "", ""
	for _, args := range os.Args[1:] {
		str += seperator + args
		seperator = " "
	}
	fmt.Println(str)
}
