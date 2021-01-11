package main
import "fmt"

const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
)
const (
        StatusOK                   = 200
        StatusCreated              = 201
        StatusAccepted             = 202
        StatusNonAuthoritativeInfo = 203
        StatusNoContent            = 204
        StatusResetContent         = 205
        StatusPartialContent       = 206
)

const Greeting = "ハローワールド" //declaring a constant

func main() {
	
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}
