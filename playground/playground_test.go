package main

import (
	"fmt"
	"testing"
)

func TestImmutability(t *testing.T) {
	bla := []bool{false}
	myFun(&bla)
	fmt.Println(bla[0], len(bla))
	myFunTwo(&bla)
	fmt.Println(bla[0], len(bla))
}

func myFun(bla *[]bool) {
	ptr := *bla
	ptr = append(ptr, true) // Does this create a function scoped variable?
	ptr[0] = true
}
func myFunTwo(bla *[]bool) {
	*bla = append(*bla, true) // Does this create a function scoped variable?
	(*bla)[0] = true
}
