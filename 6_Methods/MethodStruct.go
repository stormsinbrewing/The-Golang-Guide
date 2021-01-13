package main

import "fmt"

type student struct {
	name string
	sap  int
}

// method to string
func (s student) Register() string {
	return fmt.Sprintf("%s is registered with %d SAP ID\n", s.name, s.sap)
}

func main() {
	stu := student{"Nishkarsh", 500060720}
	fmt.Println(stu.Register())
}
