// 1. Packages and import
package main

import "fmt"

// 2. Global constants
const (
	fact = "Sun rises in the East"
)

// 3. Global variables
var (
	Exported    = "Hello, World!"
	nonExported = "Bye!"
)

// 4. Type Declarations
type Person struct {
	fName, lName string
}

type Student struct {
	sap  int
	roll string
	p    Person
}

// 5. Functions - New User with function call
func Constructor(f, l string) Student {
	obj := Student{sap: 500000000, roll: "R1XXXXXXXX", p: Person{fName: f, lName: l}}
	return obj
}

// 6. Methods
func (p Person) Greetings() string {
	return fmt.Sprintf("Hello %s\n", p.fName)
}

// 7. Driver Code
func main() {

	// Create a person object
	person1 := Person{fName: "Karan", lName: "Nautiyal"}
	fmt.Println(person1)
	fmt.Println(person1.Greetings())

	// Create a student object
	student1 := Student{sap: 500060720, roll: "R171217041", p: Person{fName: "Nishkarsh", lName: "Raj"}}
	fmt.Println(student1)
	// Composition allows child struct to call method of base struct
	fmt.Println(student1.p.Greetings())

	// Create student with Create function
	student2 := Constructor("Megha", "Rawat")
	fmt.Println(student2)
}
