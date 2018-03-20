// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// DN: good implementation of switch
	// DN: no fallthrough like in JS... (thank god)
	// DN: multiple expressions are comma separated (at least one needs to match)

	// DN: Proof for no fall through:
	fmt.Println("============")
	parsed, err := strconv.ParseFloat("36.0", 64)
	if err != nil {
		fmt.Print("Could not parse")
	} else {
		// DN: My first ever null pointer experiment in go. So it's not null pointer safe.
		// DN: But it's somewhat more difficult to make a mistake as there's usually a value there.
		// bla := err.Error()
		// fmt.Print(bla)
		j := 3.0 * 12
		fmt.Print("Write ", j, " as ")
		switch j {
		case 3.0 * 12:
			fmt.Println("one")
		case parsed:
			fmt.Println("two")
		}
		fmt.Print("Write ", j, " as ")
		switch j {
		case parsed:
			fmt.Println("two")
		case 3.0 * 12:
			fmt.Println("one")
		}
	}
	fmt.Println("============")
	// DN: This was a VERY nice experience.
	// DN: I almost couldn't get this example to run, as the compiler detected so many dumb things I tried to do.
	// DN: For example, I couldn't just trick it by using case 3.0 * 12 - 1 + 1.
	// DN: It knew that the cases were redundant.
	// DN: Also, I HAD to check that err was not nil.

	// Here's a basic `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
