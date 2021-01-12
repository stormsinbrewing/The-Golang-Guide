package main

import "fmt"

func main() {
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}

	//fmt.Println(matrix)

	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}
}
