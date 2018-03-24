package main

import "math"

// That's a bit lame that I have to write stuff like that?
// Also the concept of "just make a capital letter to export it" I don't find the best
// You'll have to modify more places than by just using a private / public modifier...
func Pow(base int, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}
