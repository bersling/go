// _range_ iterates over elements in a variety of data
// structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main

import "fmt"

func main() {

	// DN: I don't reaaally get this syntax. "for _, num := range nums"
	// DN: Can this be understood similar to the if where between the if and the curly brace you could define var?
	// DN: and here the range is kinda like a function that extracts all the index and vals from the array?
	// DN: I mean I could just memorize the syntax and I'd know how to use it but I'd like to understand it a bit better.
	// DN: Is it comprised of other GO concepts?
	// DN: So what I can recognize: There's an assignment. Range seems to be a keyword. Num is my array. For is a keyword.
	// DN: without for it wouldn't make any sense
	// DN: couldn't this just have been for _, num in nums {...} ? Why the weird range and define?
	// DN: So I geuss variables are always assigned like this. And range maybe is a function. Dunno.

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
