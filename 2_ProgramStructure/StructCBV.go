// Call by value - Default - No mutation to original values
package main
import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newRelease(a Artist) int { //passing an Artist by value
	a.Songs++
	return a.Songs
}

func main() {
	me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
