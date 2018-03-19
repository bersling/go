// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2 // DN: Why... just why... not orthogonal... JS has shown people don't use this.
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// DN: Important! Different from many other languages!
	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// DN: Important! Different from many other languages!
	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "short"` in this case.
	f := "short"
	fmt.Println(f)
	// DN: Why... I mean it's not orthogonal...
}
