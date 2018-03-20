// DN: It's not astonishing that you can define functions (methods...) for structs.
// DN: But what's interesting is that you define the methode OUTSIDE of the struct.
// DN: Doesn't this create a bit of a mess?
// DN: How can I see all the methods on rect?
// DN: I mean okay, the IDE tells me.
// DN: well why not, that way we can keep data and methods a bit separate
// DN: also the syntax with the "receiver type" is a bit weird. But well. Maybe just not used to int
// DN: It's not too complicated actually.

// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
