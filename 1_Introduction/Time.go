package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := time.Since(start).Seconds()
	fmt.Println(end)
}
