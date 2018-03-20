// DN: okay, so a struct is nice to define kind of a new type by combining some old ones.
// DN: it's similar to an interface e.g. in TypeScript?
// DN: Notice the syntax: It's not "struct bla {...}" but insteat "type bla struct {...}"

// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20}) // DN: I don't like the implicit initialization too much...

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// DN: But do I really need a pointer to mutate?
	s.age = 52
	fmt.Println(sp.age)
	// DN: Doesn't look like it... So no pointers needed here...?

}
