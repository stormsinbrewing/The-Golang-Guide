package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Greeting() string { //pointers
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Matt", "Aimonetti"} // Pointer input needs pointer object
	fmt.Println(u.Greeting())
}
