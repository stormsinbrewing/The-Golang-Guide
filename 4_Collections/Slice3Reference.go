package main

import "fmt"

func main() {
	array := [5]int{1, 3, 5, 7, 9}

	// Slices are references to arrays - If multiple slice point to same array - Changes to one slice affects all
	sl1 := array[0:3]
	sl2 := array[2:4] //Share element 2 with slice 1 we will update that later

	fmt.Println(sl1, sl2)

	//array[2] == sl2[0]
	sl2[0] = 41
	fmt.Println(sl1, sl2)
}
