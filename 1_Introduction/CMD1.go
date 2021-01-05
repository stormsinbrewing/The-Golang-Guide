// Command Line Arguments with conventional for loop
package main

import (
	"fmt"
	"os"
)

func main() {
	var str, seperator string
	for i := 1; i < len(os.Args); i++ {
		str += seperator + os.Args[i]
		seperator = " "
	}
	fmt.Println(str)
}
