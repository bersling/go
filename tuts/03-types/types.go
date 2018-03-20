package main

import "fmt"

func main() {

	// DN: The most important types would be:
	var a bool = true
	var b int = -1
	var c float64 = 25.0
	var d string = "hello world"
	var e [5]int = [5]int{1, 2, 3, 4, 5}        // Array, has fixed length
	var f []string = make([]string, 3)          // "Slice", what in other languages would be array. Also new "make keyword".
	var g map[string]int = make(map[string]int) // Map
	var p func(int, int) int = func(a int, b int) int {
		return a + b
	}
	// Note that you don't have to write it like this
	// For one, because of type inference,
	// for another, because functions can be defined directly
	// but it's useful to see how to write the type for building interfaces

	fmt.Println(a, b, c, d, e, f, g, p(1, 2))

}
