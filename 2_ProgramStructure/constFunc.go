package main
import "fmt"

// const declaration
const freeze, boil = 0,100 // multideclaration must be comma seperated, cannot do like C - declare one then comma

// user defined function - can be written before or after main
// must specify the type of argument and return type explicitly
func CtoF(cel float64) float64{
return ((9*cel + 160)/5)
}

// driver code
func main() {
fmt.Printf("Freezing Point of Water = %g\n",CtoF(freeze))
fmt.Printf("Boiling Point of Water = %g\n",CtoF(boil))
}
