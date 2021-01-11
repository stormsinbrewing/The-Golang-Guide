// In Go, a name is exported if it begins with a capital letter
package main

import (
    "fmt"
    "math"
    "net/http"
)

func main() {
    fmt.Println(math.Pi) //uppercase name
    fmt.Printf("HTTP Status OK Code is %d\n",http.StatusOK)
}
