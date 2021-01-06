package main

import "fmt"

//package level const declaration
const tempF = 312

func main() {

	var tC, tF float64
	tF = tempF
	tC = (tF - 32) * 5 / 9
	
	//Note: Printf works for C like print statements not println
	fmt.Printf("Temperature in Farenhite is %g and in Celcius, it is %g\n", tF, tC)

}
